package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/netip"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"
)

// The tests run the program in-process; run() resets its global state, so
// they must not run in parallel.

type step struct {
	kind  string // success, failure, status, output, outfile, equal
	bin   string // "" for the plain build, "geo" for the test databases
	args  []string
	args2 []string
	code  int
	want  string
	file  string
}

type testCase struct {
	name  string
	env   map[string]string
	steps []step
}

type result struct {
	code   int
	stdout string
	stderr string
	both   string // stdout and stderr in the order they were written
}

type opts struct {
	env       map[string]string
	geo       bool
	stderrTTY bool
}

func trimNL(s string) string { return strings.TrimRight(s, "\n") }

// fakeResolvers makes name lookups hermetic: localhost resolves to the
// loopback addresses, and the loopback addresses to "localhost".
func fakeResolvers(t *testing.T, hostname string) {
	t.Helper()
	oldHost, oldIP := lookupHostname, lookupIP
	t.Cleanup(func() { lookupHostname, lookupIP = oldHost, oldIP })
	lookupHostname = func(a netip.Addr) (string, bool) {
		if a.IsLoopback() {
			return hostname, true
		}
		return "", false
	}
	lookupIP = func(host string, family int) (string, bool) {
		if host != "localhost" {
			return "", false
		}
		if family == 6 {
			return "::1", true
		}
		return "127.0.0.1", true
	}
}

type interleaved struct {
	both *bytes.Buffer
	only *bytes.Buffer
}

func (w interleaved) Write(p []byte) (int, error) {
	w.both.Write(p)
	return w.only.Write(p)
}

func ipcalc(t *testing.T, o opts, args ...string) result {
	t.Helper()
	savedDBs := [3]string{countryDB, cityDB, asnDB}
	defer func() { countryDB, cityDB, asnDB = savedDBs[0], savedDBs[1], savedDBs[2] }()
	if o.geo {
		dir := geoDatabases(t)
		countryDB = filepath.Join(dir, "geo-country.mmdb")
		cityDB = filepath.Join(dir, "geo-city.mmdb")
		asnDB = filepath.Join(dir, "geo-asn.mmdb")
	} else {
		countryDB, cityDB, asnDB = "/nonexistent/country", "/nonexistent/city", "/nonexistent/asn"
	}

	var both, so, se bytes.Buffer
	code := run(environ{
		Args: args,
		Getenv: func(k string) (string, bool) {
			v, ok := o.env[k]
			return v, ok
		},
		Stdout:    interleaved{&both, &so},
		Stderr:    interleaved{&both, &se},
		StderrTTY: o.stderrTTY,
	})
	return result{code, so.String(), se.String(), both.String()}
}

func (s step) opts(env map[string]string) opts { return opts{env: env, geo: s.bin == "geo"} }

func readGolden(t *testing.T, name string) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("testdata", name))
	if err != nil {
		t.Fatalf("missing file %s", name)
	}
	return trimNL(string(b))
}

func runStep(t *testing.T, s step, env map[string]string) {
	t.Helper()
	fakeResolvers(t, "localhost")
	r := ipcalc(t, s.opts(env), s.args...)
	out := trimNL(r.both)
	switch s.kind {
	case "success":
		if r.code != 0 {
			t.Errorf("%v: exit status %d, want 0: %s", s.args, r.code, out)
		}
	case "failure":
		if r.code == 0 {
			t.Errorf("%v: succeeded, want failure: %s", s.args, out)
		}
	case "status":
		if r.code != s.code {
			t.Errorf("%v: exit status %d, want %d: %s", s.args, r.code, s.code, out)
		}
	case "output":
		if out != s.want {
			t.Errorf("%v: output\n%s\nwant\n%s", s.args, out, s.want)
		}
	case "outfile":
		want := readGolden(t, s.file)
		if want == "" {
			t.Fatalf("empty file %s", s.file)
		}
		if out != want {
			t.Errorf("%v: output differs from %s\ngot:\n%s\nwant:\n%s", s.args, s.file, out, want)
		}
		if r.code != 0 {
			t.Errorf("%v: exit status %d", s.args, r.code)
		}
	case "equal":
		r2 := ipcalc(t, s.opts(env), s.args2...)
		out2 := trimNL(r2.both)
		if r.code != 0 || r2.code != 0 || out == "" || out2 == "" {
			t.Fatalf("%v / %v: failed or empty output: %q %q", s.args, s.args2, out, out2)
		}
		if out != out2 {
			t.Errorf("%q != %q", out, out2)
		}
	default:
		t.Fatalf("unknown step kind %q", s.kind)
	}
}

func TestGenerated(t *testing.T) {
	for _, c := range generatedCases {
		t.Run(c.name, func(t *testing.T) {
			for _, s := range c.steps {
				runStep(t, s, c.env)
			}
		})
	}
}

// mustRun runs ipcalc and fails the test unless it succeeds.
func mustRun(t *testing.T, args ...string) string {
	t.Helper()
	fakeResolvers(t, "localhost")
	r := ipcalc(t, opts{}, args...)
	if r.code != 0 {
		t.Fatalf("%v: exit status %d: %s", args, r.code, r.both)
	}
	return r.stdout
}

func TestRandomNetworks(t *testing.T) {
	private := regexp.MustCompile(`(?m)^Network:\s+(10\.|172\.(1[6-9]|2[0-9]|3[01])\.|192\.168\.)`)
	ula := regexp.MustCompile(`(?m)^Network:\s+fd`)
	address := regexp.MustCompile(`(?m)^Address:`)
	for i := 0; i < 100; i++ {
		if out := mustRun(t, "-r", "24"); !private.MatchString(out) || address.MatchString(out) {
			t.Fatalf("-r 24: %s", out)
		}
		for _, args := range [][]string{{"-r", "112"}, {"-6", "-r", "24"}} {
			if out := mustRun(t, args...); !ula.MatchString(out) || address.MatchString(out) {
				t.Fatalf("%v: %s", args, out)
			}
		}
		if out := mustRun(t, "-r", "24", "-a"); !strings.Contains(out, "ADDRESS") {
			t.Fatalf("-r 24 -a: %s", out)
		}
		if out := mustRun(t, "-6", "-r", "48", "-a"); !strings.HasPrefix(out, "ADDRESS=fd") {
			t.Fatalf("-6 -r 48 -a: %s", out)
		}
	}
}

func TestSplitLinesByNetwork(t *testing.T) {
	for _, args := range [][]string{
		{"-S", "29", "192.168.5.0/24"},
		{"-S", "120", "fcfa:b4ca:f1d8:125b:dc00::/112"},
	} {
		out := mustRun(t, args...)
		nets, total := 0, ""
		for _, l := range strings.Split(out, "\n") {
			if strings.HasPrefix(l, "Network") {
				nets++
			}
			if strings.HasPrefix(l, "Total") {
				total = strings.TrimSpace(strings.SplitN(l, ":", 2)[1])
			}
		}
		if fmt.Sprint(nets) != total {
			t.Errorf("%v: %d network lines, total %q", args, nets, total)
		}
	}
}

// shellVars parses the NAME=value lines of the shell output format.
func shellVars(t *testing.T, out string) map[string]string {
	t.Helper()
	vars := map[string]string{}
	for _, l := range strings.Split(trimNL(out), "\n") {
		name, val, ok := strings.Cut(l, "=")
		if !ok {
			t.Fatalf("not a NAME=value line: %q", l)
		}
		if strings.HasPrefix(val, `"`) {
			var sb strings.Builder
			for i := 1; i < len(val)-1; i++ {
				if val[i] == '\\' {
					i++
				}
				sb.WriteByte(val[i])
			}
			val = sb.String()
		}
		vars[name] = val
	}
	return vars
}

func lines(s string) []string { return strings.Split(trimNL(s), "\n") }

// Deaggregating the range of a network, or of the first and last subnets
// of a split, gives back the network.
func TestDeaggregateRoundTrip(t *testing.T) {
	for i := 0; i < 50; i++ {
		v := shellVars(t, mustRun(t, "-r", "96", "--prefix", "-n", "--minaddr", "--maxaddr"))
		want := v["NETWORK"] + "/" + v["PREFIX"]
		got := trimNL(mustRun(t, "-d", v["MINADDR"]+"-"+v["MAXADDR"], "--no-decorate"))
		if got != want {
			t.Fatalf("random IPv6: deaggregated %q, want %q", got, want)
		}
	}
	for _, c := range []struct {
		random, split, first, last string
	}{
		{"24", "26", "-n", "-b"},
		{"56", "64", "--minaddr", "--maxaddr"},
	} {
		for i := 0; i < 50; i++ {
			v := shellVars(t, mustRun(t, "-r", c.random, "--prefix", "-n"))
			net := v["NETWORK"] + "/" + v["PREFIX"]
			split := lines(mustRun(t, "-S", c.split, net, "--no-decorate"))
			first := trimNL(mustRun(t, "--no-decorate", c.first, split[0]))
			last := trimNL(mustRun(t, "--no-decorate", c.last, split[len(split)-1]))
			got := trimNL(mustRun(t, "--no-decorate", "-d", first+"-"+last))
			if got != net {
				t.Fatalf("%s: deaggregated %q, want %q", net, got, net)
			}
		}
	}
}

// expandTabs is expand(1): tab stops every 8 columns.
func expandTabs(s string) string {
	var sb strings.Builder
	col := 0
	for _, r := range s {
		if r == '\t' {
			n := 8 - col%8
			sb.WriteString(strings.Repeat(" ", n))
			col += n
		} else {
			sb.WriteRune(r)
			col++
		}
	}
	return sb.String()
}

func TestInfoSummarySize(t *testing.T) {
	for _, addr := range []string{"ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/1", "255.255.255.255/1"} {
		out := mustRun(t, "-i", addr)
		ls := lines(out)
		if len(ls) > 15 {
			t.Errorf("%s: %d lines", addr, len(ls))
		}
		for _, l := range ls {
			if n := len([]rune(expandTabs(l))); n > 80 {
				t.Errorf("%s: line of %d columns: %q", addr, n, l)
			}
		}
	}
}

// Every field selected by an option is in the all-info JSON output.
func TestAllInfoJSONKeys(t *testing.T) {
	key := regexp.MustCompile(`"[A-Z0-9]*":`)
	for _, addr := range []string{"192.168.2.7/24", "fe80::21b:21ff:fe3a:5c7d/64"} {
		all := map[string]bool{}
		for _, k := range key.FindAllString(mustRun(t, "-j", "--all-info", addr), -1) {
			all[k] = true
		}
		sel := key.FindAllString(mustRun(t, "-j", "-a", "-b", "-m", "-n", "-p", "--cidr",
			"--wildcard", "--reverse-dns", "--addrspace", "--minaddr", "--maxaddr",
			"--addresses", addr), -1)
		if len(sel) == 0 {
			t.Fatalf("%s: no keys", addr)
		}
		for _, k := range sel {
			if !all[k] {
				t.Errorf("%s: missing %s", addr, k)
			}
		}
	}
}

func TestAllInfoNoBits(t *testing.T) {
	for _, args := range [][]string{
		{"--all-info", "192.168.92.6/22"},
		{"--all-info", "--format=json", "192.168.92.6/22"},
		{"--all-info", "--format=shell", "2001:db8::1/56"},
	} {
		if out := mustRun(t, args...); strings.Contains(strings.ToLower(out), "bits") {
			t.Errorf("%v: mentions bits: %s", args, out)
		}
	}
}

// A name that comes from outside ipcalc is escaped in every format.
func TestHostnameEscaping(t *testing.T) {
	const name = "a\"b\\c\td"
	fakeResolvers(t, name)

	r := ipcalc(t, opts{}, "-j", "-h", "127.0.0.1")
	var j map[string]any
	if err := json.Unmarshal([]byte(r.stdout), &j); err != nil || j["HOSTNAME"] != name {
		t.Errorf("JSON: %v %v", err, r.stdout)
	}
	for _, args := range [][]string{{"--format=shell", "-h", "127.0.0.1"}, {"-h", "127.0.0.1"}} {
		r := ipcalc(t, opts{}, args...)
		if v := shellVars(t, r.stdout); v["HOSTNAME"] != name {
			t.Errorf("%v: HOSTNAME=%q", args, v["HOSTNAME"])
		}
	}
	for _, args := range [][]string{{"--no-decorate", "-h", "127.0.0.1"}, {"--format=value", "-h", "127.0.0.1"}} {
		r := ipcalc(t, opts{}, args...)
		if r.code != 0 || trimNL(r.both) != name {
			t.Errorf("%v: %d %q", args, r.code, r.both)
		}
	}
	if r := ipcalc(t, opts{}, "-j", "--all-info", "127.0.0.1"); strings.Contains(r.stdout, "HOSTNAME") {
		t.Errorf("--all-info looked up the hostname: %s", r.stdout)
	}
}

func TestHostnameNetworkNoLookup(t *testing.T) {
	oldHost := lookupHostname
	defer func() { lookupHostname = oldHost }()
	lookupHostname = func(netip.Addr) (string, bool) {
		t.Error("looked up the name of a network")
		return "", false
	}
	r := ipcalc(t, opts{}, "-h", "127.0.0.1/8")
	if r.code != 1 || trimNL(r.both) != "ipcalc: -h requires a single address, not a network: 127.0.0.1/8" {
		t.Errorf("%d %q", r.code, r.both)
	}
}

// Deprecated options suggest --format, on a terminal only.
func TestFormatHints(t *testing.T) {
	for _, c := range []struct{ name, args, hint string }{
		{"HintJson", "-j 10.0.0.1/24", "ipcalc: hint: -j is deprecated; use --format=json"},
		{"HintJsonFields", "-j -m 10.0.0.1/24", "ipcalc: hint: -j is deprecated; use --format=json, which prints only the selected fields"},
		{"HintJsonEquals", "-j 10.0.0.0/8 --equals=10.0.0.0/8", "ipcalc: hint: -j is deprecated; use -s to print no result"},
		{"HintNoDecorateFields", "--no-decorate -m 10.0.0.1/24", "ipcalc: hint: --no-decorate is deprecated; use --format=value"},
		{"HintNoDecorateSplit", "--no-decorate -S 26 10.0.0.0/24", "ipcalc: hint: --no-decorate is deprecated; use --format=value"},
		{"HintNoDecorateSummary", "--no-decorate 10.0.0.1/24", "ipcalc: hint: --no-decorate is deprecated; use --format=human"},
		{"NoHintFormat", "--format=json -m 10.0.0.1/24", ""},
		{"NoHintFields", "-m -n 10.0.0.1/24", ""},
		{"NoHintFormatValueSplit", "--format=value -S 26 10.0.0.0/24", ""},
		{"NoHintSilent", "-s -j 10.0.0.1/24", ""},
	} {
		t.Run(c.name, func(t *testing.T) {
			r := ipcalc(t, opts{stderrTTY: true}, strings.Fields(c.args)...)
			if r.code != 0 || trimNL(r.stderr) != c.hint {
				t.Errorf("%d %q, want hint %q", r.code, r.stderr, c.hint)
			}
		})
	}
}

// The values of the split-hosts layout match the Debian ipcalc (-s COUNT...),
// whose recorded output is in testdata/debian. Both outputs are reduced to
// "<request> <field> <value>" lines, since the tools lay out the fields
// differently.
func normalizeSplit(out string) []string {
	var res []string
	req, unused := 0, 0
	inUnused := false
	for _, l := range strings.Split(out, "\n") {
		f := strings.Fields(l)
		if len(f) == 0 {
			continue
		}
		field := func(i int) string {
			if i < len(f) {
				return f[i]
			}
			return ""
		}
		switch {
		case regexp.MustCompile(`^[0-9]+\. Requested size:`).MatchString(l):
			req++
			res = append(res, fmt.Sprintf("%d Requested %s", req, field(3)))
			continue
		case strings.HasPrefix(l, "Total") || strings.HasPrefix(l, "Needed size:"):
			if strings.HasPrefix(l, "Total") || strings.HasPrefix(l, "Needed size:") {
				if req > 0 {
					req = 0
				}
			}
		}
		if req > 0 {
			switch {
			case regexp.MustCompile(`^(Network|Broadcast|HostMin|HostMax|Hosts/Net):`).MatchString(l):
				res = append(res, fmt.Sprintf("%d %s %s", req, field(0), field(1)))
				continue
			case strings.HasPrefix(l, "Netmask:"):
				res = append(res, fmt.Sprintf("%d %s %s %s", req, field(0), field(1), field(3)))
				continue
			}
		}
		switch {
		case strings.HasPrefix(l, "Needed size:"):
			res = append(res, "Needed "+field(2))
		case strings.HasPrefix(l, "Used network:"):
			res = append(res, "Used "+field(2))
		case strings.HasPrefix(l, "Unused:") || strings.HasPrefix(l, "[Unused networks]"):
			inUnused = true
		case inUnused && strings.HasPrefix(l, "Network:"):
			unused++
			res = append(res, fmt.Sprintf("Unused %d %s", unused, field(1)))
		case inUnused && l[0] >= '0' && l[0] <= '9':
			unused++
			res = append(res, fmt.Sprintf("Unused %d %s", unused, field(0)))
		}
	}
	sort.Strings(res)
	return res
}

func TestDebianSplitHosts(t *testing.T) {
	for _, c := range debianSplitHostsCases {
		t.Run(c.network+"-"+c.counts, func(t *testing.T) {
			got := normalizeSplit(mustRun(t, "--split-hosts="+c.counts, c.network))
			g, _ := os.ReadFile(filepath.Join("testdata", "debian", c.golden))
			want := strings.Split(trimNL(string(g)), "\n")
			sort.Strings(want)
			if len(got) == 0 || strings.Join(got, "\n") != strings.Join(want, "\n") {
				t.Errorf("differs from the Debian ipcalc\ngot:\n%s\nwant:\n%s",
					strings.Join(got, "\n"), strings.Join(want, "\n"))
			}
		})
	}
}

func TestVersion(t *testing.T) {
	old := version
	defer func() { version = old }()
	version = "1.2.3"
	if r := ipcalc(t, opts{}, "-v"); r.code != 0 || r.stdout != "ipcalc 1.2.3\n" {
		t.Errorf("%d %q", r.code, r.stdout)
	}
	version = ""
	if r := ipcalc(t, opts{}, "--version"); r.code != 0 || !strings.HasPrefix(r.stdout, "ipcalc ") {
		t.Errorf("%d %q", r.code, r.stdout)
	}
}
