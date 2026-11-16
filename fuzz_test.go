package main

import (
	"fmt"
	"math/big"
	"net/netip"
	"strconv"
	"strings"
	"testing"
)

// Property tests: they check what the output must satisfy, not what the C
// implementation printed. `go test` runs the seeds; `go test -fuzz=FuzzName`
// explores further.

func runOK(t *testing.T, args ...string) string {
	t.Helper()
	r := ipcalc(t, opts{}, args...)
	if r.code != 0 {
		t.Fatalf("%v: exit status %d: %s", args, r.code, r.both)
	}
	return r.stdout
}

func addrNum(a netip.Addr) *big.Int { return new(big.Int).SetBytes(a.AsSlice()) }

func u128Addr(hi, lo uint64, v6 bool) netip.Addr {
	if !v6 {
		return netip.AddrFrom4([4]byte{byte(lo >> 24), byte(lo >> 16), byte(lo >> 8), byte(lo)})
	}
	var b [16]byte
	for i := 0; i < 8; i++ {
		b[i] = byte(hi >> (56 - 8*uint(i)))
		b[8+i] = byte(lo >> (56 - 8*uint(i)))
	}
	return netip.AddrFrom16(b)
}

func parsePrefixes(t *testing.T, out string) []netip.Prefix {
	t.Helper()
	var ps []netip.Prefix
	for _, l := range lines(out) {
		p, err := netip.ParsePrefix(l)
		if err != nil {
			t.Fatalf("not a prefix: %q", l)
		}
		ps = append(ps, p)
	}
	return ps
}

// checkTiling verifies that the prefixes are aligned and cover first..last
// exactly, in order and without gaps or overlaps.
func checkTiling(t *testing.T, ps []netip.Prefix, first, last netip.Addr) {
	t.Helper()
	if len(ps) == 0 {
		t.Fatal("no networks")
	}
	next := addrNum(first)
	one := big.NewInt(1)
	for _, p := range ps {
		if p.Masked() != p {
			t.Fatalf("%s has host bits set", p)
		}
		if addrNum(p.Addr()).Cmp(next) != 0 {
			t.Fatalf("%s does not start at %s", p, next)
		}
		size := new(big.Int).Lsh(one, uint(p.Addr().BitLen()-p.Bits()))
		next.Add(next, size)
	}
	end := new(big.Int).Add(addrNum(last), one)
	if next.Cmp(end) != 0 {
		t.Fatalf("ends at %s, want %s", next, end)
	}
}

func FuzzDeaggregate(f *testing.F) {
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0xffffffff), false)
	f.Add(uint64(0), uint64(0xc0a80201), uint64(0), uint64(0xc0a8030a), false)
	f.Add(uint64(0), uint64(5), uint64(0), uint64(5), false)
	f.Add(uint64(0x20010db8), uint64(1), uint64(0x20010db8), uint64(0xffff), true)
	f.Add(uint64(0), uint64(0), ^uint64(0), ^uint64(0), true)
	f.Add(^uint64(0), ^uint64(0), ^uint64(0), ^uint64(0), true)
	f.Fuzz(func(t *testing.T, ahi, alo, bhi, blo uint64, v6 bool) {
		a, b := u128Addr(ahi, alo, v6), u128Addr(bhi, blo, v6)
		if a.Compare(b) > 0 {
			a, b = b, a
		}
		out := runOK(t, "--format=value", "-d", a.String()+"-"+b.String())
		ps := parsePrefixes(t, out)
		checkTiling(t, ps, a, b)
		// minimal: no two neighbors are the two halves of one network
		for i := 1; i < len(ps); i++ {
			x, y := ps[i-1], ps[i]
			if x.Bits() == y.Bits() && x.Bits() > 0 {
				parent, _ := x.Addr().Prefix(x.Bits() - 1)
				if parent.Contains(y.Addr()) && parent.Addr() == x.Addr() {
					t.Fatalf("%s and %s could be merged", x, y)
				}
			}
		}
	})
}

func FuzzSplit(f *testing.F) {
	f.Add(uint64(0), uint64(0xc0a80500), 24, 26, false)
	f.Add(uint64(0), uint64(0), 0, 4, false)
	f.Add(uint64(0x20010db8), uint64(0), 32, 40, true)
	f.Add(^uint64(0), ^uint64(0), 120, 124, true)
	f.Fuzz(func(t *testing.T, hi, lo uint64, p, q int, v6 bool) {
		width := 32
		if v6 {
			width = 128
		}
		if p < 0 || p > width || q < p || q > width || q-p > 8 {
			t.Skip()
		}
		base, _ := u128Addr(hi, lo, v6).Prefix(p)
		out := runOK(t, "--format=value", "-S", strconv.Itoa(q), base.String())
		ps := parsePrefixes(t, out)
		if len(ps) != 1<<uint(q-p) {
			t.Fatalf("%d subnets, want %d", len(ps), 1<<uint(q-p))
		}
		last := lastAddr(base)
		checkTiling(t, ps, base.Addr(), last)
		for _, s := range ps {
			if s.Bits() != q {
				t.Fatalf("%s is not a /%d", s, q)
			}
		}
	})
}

func lastAddr(p netip.Prefix) netip.Addr {
	b := p.Addr().AsSlice()
	for i := range b {
		bits := 8
		if rem := p.Bits() - 8*i; rem < 8 {
			bits = max(rem, 0)
		}
		b[i] |= byte(0xff >> uint(bits))
	}
	a, _ := netip.AddrFromSlice(b)
	return a
}

// The summary of a network agrees with netip's arithmetic.
func FuzzInfo(f *testing.F) {
	f.Add(uint64(0), uint64(0xc0a80207), 24, false)
	f.Add(uint64(0), uint64(0x0a000001), 31, false)
	f.Add(uint64(0), uint64(0x0a000001), 32, false)
	f.Add(uint64(0), uint64(0), 0, false)
	f.Add(uint64(0x20010db8), uint64(1), 64, true)
	f.Add(^uint64(0), ^uint64(0), 128, true)
	f.Add(uint64(0), uint64(0), 0, true)
	f.Fuzz(func(t *testing.T, hi, lo uint64, p int, v6 bool) {
		width := 32
		if v6 {
			width = 128
		}
		if p < 0 || p > width {
			t.Skip()
		}
		addr := u128Addr(hi, lo, v6)
		if addr.Is4In6() {
			t.Skip()
		}
		in := fmt.Sprintf("%s/%d", addr, p)
		vars := shellVars(t, runOK(t, "--format=shell", "-a", "-n", "-p", "--minaddr", "--maxaddr", in))
		net, _ := addr.Prefix(p)
		if vars["NETWORK"] != net.Addr().String() && !v6 {
			t.Fatalf("%s: NETWORK=%s, want %s", in, vars["NETWORK"], net.Addr())
		}
		if v6 {
			got, err := netip.ParseAddr(vars["NETWORK"])
			if err != nil || got != net.Addr() {
				t.Fatalf("%s: NETWORK=%s, want %s", in, vars["NETWORK"], net.Addr())
			}
		}
		if vars["PREFIX"] != strconv.Itoa(p) {
			t.Fatalf("%s: PREFIX=%s", in, vars["PREFIX"])
		}
		min, err1 := netip.ParseAddr(vars["MINADDR"])
		max, err2 := netip.ParseAddr(vars["MAXADDR"])
		if err1 != nil || err2 != nil || min.Compare(max) > 0 || !net.Contains(min) || !net.Contains(max) {
			t.Fatalf("%s: host range %s-%s", in, vars["MINADDR"], vars["MAXADDR"])
		}
		// a network is its own network, and equals itself
		if r := ipcalc(t, opts{}, "--equals", net.String(), in); r.code != 0 {
			t.Fatalf("%s: network %s not equal to its input: %s", in, net, r.both)
		}
		if p > 0 {
			super, _ := addr.Prefix(p - 1)
			if r := ipcalc(t, opts{}, "-s", "--subnet-of", super.String(), in); r.code != 0 {
				t.Fatalf("%s not a subnet of %s", in, super)
			}
			if r := ipcalc(t, opts{}, "-s", "--overlaps", in, super.String()); r.code != 0 {
				t.Fatalf("%s and %s do not overlap", in, super)
			}
		}
		// the netmask form of the prefix gives the same network (IPv4)
		if !v6 {
			mask := vars2(t, "-m", in)
			if r := runOK(t, "--format=value", "-p", addr.String(), mask); strings.TrimSpace(r) != strconv.Itoa(p) {
				t.Fatalf("%s %s: prefix %q", addr, mask, r)
			}
		}
	})
}

func vars2(t *testing.T, opt, in string) string {
	t.Helper()
	return trimNL(runOK(t, "--format=value", opt, in))
}

// Whatever the arguments, the program must exit with a status and not panic.
func FuzzArgs(f *testing.F) {
	for _, s := range []string{
		"192.168.2.7/24", "-S 26 10.0.0.0/24", "-d 1.2.3.4-5.6.7.8", "--binary ::1/128",
		"--equals 10.0.0.0/8 10.1.0.0/16", "--split-hosts=10,20 10.0.0.0/24", "-g 192.0.2.5",
		"-r 24", "-6 -r 64 -a", "--all-info fe80::1%eth0", "-m 1.2.3.4 255.0.255.0",
		"--format=json -n -- -1.2.3.4", "-S 1:25 -S 25 10.0.0.0/24 extra", "-c 300.1.1.1",
	} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		args := strings.Fields(s)
		if len(args) > 12 {
			t.Skip()
		}
		for _, a := range args {
			if a == "-r" || strings.HasPrefix(a, "--random") {
				t.Skip() // random output, and slow to explore
			}
		}
		fakeResolvers(t, "localhost")
		ipcalc(t, opts{geo: true}, args...)
	})
}
