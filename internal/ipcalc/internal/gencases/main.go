// SPDX-License-Identifier: GPL-2.0-only
//
// Copyright (c) 2026 go-ipcalc contributors
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License, version 2,
// as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// Command gencases writes the generated test tables of package ipcalc. It
// runs in internal/ipcalc and reads testdata/upstream/meson-tests.json (the
// upstream test list) and testdata/color-cases.json.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

const buildDir = "@BUILD@/"

type mesonTest struct {
	Cmd  []string          `json:"cmd"`
	Env  map[string]string `json:"env"`
	Name string            `json:"name"`
}

type colorCase struct {
	Name string            `json:"name"`
	Bin  string            `json:"bin"`
	Args []string          `json:"args"`
	Env  map[string]string `json:"env"`
}

type kv struct{ k, v string }

type step struct {
	kind, bin   string
	args, args2 []string
	code        int
	hasCode     bool
	want, file  string
	hasWant     bool
}

type parsed struct {
	bin  string
	args []string
}

var shellMeta = regexp.MustCompile("[|;&$`<>]|\\bgrep\\b|\\bjq\\b|LD_PRELOAD|\\bperl\\b")
var envAssign = regexp.MustCompile(`^[A-Z_]+=.*$`)

// shlex splits like Python's shlex.split in POSIX mode.
func shlex(s string) ([]string, error) {
	var out []string
	var cur strings.Builder
	inWord := false
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		c := rs[i]
		switch c {
		case ' ', '\t', '\n', '\r':
			if inWord {
				out = append(out, cur.String())
				cur.Reset()
				inWord = false
			}
		case '\\':
			if i+1 >= len(rs) {
				return nil, fmt.Errorf("no escaped character in %q", s)
			}
			i++
			cur.WriteRune(rs[i])
			inWord = true
		case '\'':
			inWord = true
			i++
			for ; i < len(rs) && rs[i] != '\''; i++ {
				cur.WriteRune(rs[i])
			}
			if i >= len(rs) {
				return nil, fmt.Errorf("no closing quotation in %q", s)
			}
		case '"':
			inWord = true
			i++
			for ; i < len(rs) && rs[i] != '"'; i++ {
				if rs[i] == '\\' && i+1 < len(rs) && (rs[i+1] == '"' || rs[i+1] == '\\') {
					i++
				}
				cur.WriteRune(rs[i])
			}
			if i >= len(rs) {
				return nil, fmt.Errorf("no closing quotation in %q", s)
			}
		default:
			cur.WriteRune(c)
			inWord = true
		}
	}
	if inWord {
		out = append(out, cur.String())
	}
	return out, nil
}

// parse reduces a test command to the ipcalc binary it runs and its
// arguments; commands that need a shell are not representable.
func parse(c string, env *[]kv) *parsed {
	if shellMeta.MatchString(c) {
		return nil
	}
	a, err := shlex(c)
	if err != nil {
		log.Fatal(err)
	}
	for len(a) > 0 && envAssign.MatchString(a[0]) {
		k, v, _ := strings.Cut(a[0], "=")
		setEnv(env, k, v)
		a = a[1:]
	}
	if len(a) == 0 {
		return nil
	}
	switch a[0] {
	case buildDir + "ipcalc":
		return &parsed{"", a[1:]}
	case buildDir + "tests/ipcalc-geo":
		return &parsed{"geo", a[1:]}
	}
	return nil
}

func setEnv(env *[]kv, k, v string) {
	for i := range *env {
		if (*env)[i].k == k {
			(*env)[i].v = v
			return
		}
	}
	*env = append(*env, kv{k, v})
}

// q quotes like Python's json.dumps(s, ensure_ascii=False).
func q(s string) string {
	var b strings.Builder
	b.WriteByte('"')
	for _, r := range s {
		switch {
		case r == '"':
			b.WriteString(`\"`)
		case r == '\\':
			b.WriteString(`\\`)
		case r == '\n':
			b.WriteString(`\n`)
		case r == '\r':
			b.WriteString(`\r`)
		case r == '\t':
			b.WriteString(`\t`)
		case r == '\b':
			b.WriteString(`\b`)
		case r == '\f':
			b.WriteString(`\f`)
		case r < 0x20:
			fmt.Fprintf(&b, `\u%04x`, r)
		default:
			b.WriteRune(r)
		}
	}
	b.WriteByte('"')
	return b.String()
}

func sl(a []string) string {
	p := make([]string, len(a))
	for i, v := range a {
		p[i] = q(v)
	}
	return "[]string{" + strings.Join(p, ", ") + "}"
}

func envLit(env []kv) string {
	p := make([]string, len(env))
	for i, e := range env {
		p[i] = q(e.k) + ": " + q(e.v)
	}
	return "map[string]string{" + strings.Join(p, ", ") + "}"
}

func write(name, src string) {
	out, err := format.Source([]byte(src))
	if err != nil {
		log.Fatalf("%s: %v", name, err)
	}
	if err := os.WriteFile(name, out, 0o644); err != nil {
		log.Fatal(err)
	}
}

func readJSON(name string, v any) {
	b, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(b, v); err != nil {
		log.Fatalf("%s: %v", name, err)
	}
}

type genCase struct {
	name  string
	env   []kv
	steps []step
}

// convert turns a --test-* runner command line into steps.
func convert(cmd []string, env *[]kv) ([]step, bool) {
	var steps []step
	for i := 1; i < len(cmd); {
		m := cmd[i]
		arg := func(n int) string {
			if i+n >= len(cmd) {
				log.Fatalf("truncated test command %q", cmd)
			}
			return cmd[i+n]
		}
		switch m {
		case "--test-success", "--test-failure":
			p := parse(arg(1), env)
			if p == nil {
				return nil, false
			}
			steps = append(steps, step{kind: m[7:], bin: p.bin, args: p.args})
			i += 2
		case "--test-status":
			p := parse(arg(2), env)
			if p == nil {
				return nil, false
			}
			code, err := strconv.Atoi(arg(1))
			if err != nil {
				log.Fatal(err)
			}
			steps = append(steps, step{kind: "status", bin: p.bin, args: p.args, code: code, hasCode: true})
			i += 3
		case "--test-output":
			p := parse(arg(1), env)
			if p == nil {
				return nil, false
			}
			steps = append(steps, step{kind: "output", bin: p.bin, args: p.args, want: arg(2), hasWant: true})
			i += 3
		case "--test-outfile":
			p := parse(arg(1), env)
			if p == nil {
				return nil, false
			}
			steps = append(steps, step{kind: "outfile", bin: p.bin, args: p.args, file: path.Base(arg(2))})
			i += 3
		case "--test-equal":
			p, p2 := parse(arg(1), env), parse(arg(2), env)
			if p == nil || p2 == nil {
				return nil, false
			}
			steps = append(steps, step{kind: "equal", bin: p.bin, args: p.args, args2: p2.args})
			i += 3
		default:
			return nil, false
		}
	}
	return steps, true
}

func genMain(tests []mesonTest) {
	var cases []genCase
	for _, t := range tests {
		if len(t.Cmd) < 2 || !strings.HasPrefix(t.Cmd[1], "--test") {
			continue
		}
		var env []kv
		steps, ok := convert(t.Cmd, &env)
		if !ok {
			continue
		}
		if v, ok := t.Env["COLUMNS"]; ok {
			setEnv(&env, "COLUMNS", v)
		}
		cases = append(cases, genCase{t.Name, env, steps})
	}
	var b strings.Builder
	b.WriteString("// Code generated from the ipcalc C project test list; DO NOT EDIT.\n\npackage ipcalc\n\nvar generatedCases = []testCase{\n")
	for _, c := range cases {
		fmt.Fprintf(&b, "\t{name: %s,\n", q(c.name))
		if len(c.env) > 0 {
			fmt.Fprintf(&b, "\t\tenv: %s,\n", envLit(c.env))
		}
		b.WriteString("\t\tsteps: []step{\n")
		for _, s := range c.steps {
			f := []string{"kind: " + q(s.kind)}
			if s.bin != "" {
				f = append(f, "bin: "+q(s.bin))
			}
			f = append(f, "args: "+sl(s.args))
			if s.args2 != nil {
				f = append(f, "args2: "+sl(s.args2))
			}
			if s.hasCode {
				f = append(f, "code: "+strconv.Itoa(s.code))
			}
			if s.hasWant {
				f = append(f, "want: "+q(s.want))
			}
			if s.file != "" {
				f = append(f, "file: "+q(s.file))
			}
			fmt.Fprintf(&b, "\t\t\t{%s},\n", strings.Join(f, ", "))
		}
		b.WriteString("\t\t}},\n")
	}
	b.WriteString("}\n")
	write("cases_gen_test.go", b.String())
}

// genDebian derives the golden names from the arguments of the
// ipcalc-debian-split-hosts.sh tests.
func genDebian(tests []mesonTest) {
	var b bytes.Buffer
	b.WriteString("// Code generated from the Debian ipcalc outputs; DO NOT EDIT.\n\npackage ipcalc\n\nvar debianSplitHostsCases = []struct{ network, counts, golden string }{\n")
	for _, t := range tests {
		if len(t.Cmd) != 3 || !strings.HasSuffix(t.Cmd[0], "/ipcalc-debian-split-hosts.sh") {
			continue
		}
		network, counts := t.Cmd[1], t.Cmd[2]
		golden := strings.ReplaceAll(network, "/", "-") + "-" + counts
		fmt.Fprintf(&b, "\t{%s, %s, %s},\n", q(network), q(counts), q(golden))
	}
	b.WriteString("}\n")
	write("debian_cases_gen_test.go", b.String())
}

func genColor() {
	var cases []colorCase
	readJSON("testdata/color-cases.json", &cases)
	var b strings.Builder
	b.WriteString("// Code generated from the output of the C ipcalc on a terminal; DO NOT EDIT.\n\npackage ipcalc\n\nvar colorCases = []struct {\n\tname, bin string\n\targs []string\n\tenv map[string]string\n}{\n")
	for _, c := range cases {
		env := "nil"
		if len(c.Env) > 0 {
			var kvs []kv
			for k, v := range c.Env {
				kvs = append(kvs, kv{k, v})
			}
			sortKV(kvs)
			env = envLit(kvs)
		}
		fmt.Fprintf(&b, "\t{%s, %s, %s, %s},\n", q(c.Name), q(c.Bin), sl(c.Args), env)
	}
	b.WriteString("}\n")
	write("color_cases_gen_test.go", b.String())
}

func sortKV(a []kv) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && a[j].k < a[j-1].k; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}

func main() {
	log.SetFlags(0)
	var tests []mesonTest
	readJSON("testdata/upstream/meson-tests.json", &tests)
	genMain(tests)
	genDebian(tests)
	genColor()
}
