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

package getopt

import (
	"bytes"
	"reflect"
	"testing"
)

var testLong = []LongOption{
	{Name: "netmask", HasArg: false, Val: 'm'},
	{Name: "network", HasArg: false, Val: 'n'},
	{Name: "split", HasArg: true, Val: 'S'},
	{Name: "split-hosts", HasArg: true, Val: 1000},
	{Name: "check", HasArg: false, Val: 'c'},
	{Name: "binary", HasArg: false, Val: 1001},
}

const testShort = "S:cr:i46mn"

type parsed struct {
	opts []Result
	rest []string
	err  string
}

func parse(args ...string) parsed {
	var errOut bytes.Buffer
	p := New("prog", &errOut, args, testShort, testLong)
	var out parsed
	for {
		r, ok := p.Next()
		if !ok {
			break
		}
		out.opts = append(out.opts, r)
	}
	out.rest = p.Remaining()
	out.err = errOut.String()
	return out
}

func TestParse(t *testing.T) {
	tests := []struct {
		name string
		args []string
		opts []Result
		rest []string
	}{
		{"empty", nil, nil, nil},
		{"operands only", []string{"a", "b"}, nil, []string{"a", "b"}},
		{"lone dash is an operand", []string{"-"}, nil, []string{"-"}},
		{"short options", []string{"-4", "-c"}, []Result{{Opt: '4'}, {Opt: 'c'}}, nil},
		{"bundled", []string{"-4cm"}, []Result{{Opt: '4'}, {Opt: 'c'}, {Opt: 'm'}}, nil},
		{"bundled ending in an argument", []string{"-c4S26"}, []Result{{Opt: 'c'}, {Opt: '4'}, {Opt: 'S', Arg: "26"}}, nil},
		{"short argument attached", []string{"-S26"}, []Result{{Opt: 'S', Arg: "26"}}, nil},
		{"short argument separate", []string{"-S", "26"}, []Result{{Opt: 'S', Arg: "26"}}, nil},
		{"argument may look like an option", []string{"-S", "-c"}, []Result{{Opt: 'S', Arg: "-c"}}, nil},
		{"long", []string{"--check"}, []Result{{Opt: 'c'}}, nil},
		{"long with =", []string{"--split=26"}, []Result{{Opt: 'S', Arg: "26"}}, nil},
		{"long with empty =", []string{"--split="}, []Result{{Opt: 'S'}}, nil},
		{"long with separate argument", []string{"--split", "26"}, []Result{{Opt: 'S', Arg: "26"}}, nil},
		{"long argument containing =", []string{"--split=a=b"}, []Result{{Opt: 'S', Arg: "a=b"}}, nil},
		{"unique abbreviation", []string{"--chec", "--bin"}, []Result{{Opt: 'c'}, {Opt: 1001}}, nil},
		{"exact name beats longer ones", []string{"--split", "1"}, []Result{{Opt: 'S', Arg: "1"}}, nil},
		{"abbreviation with argument", []string{"--split-h=5"}, []Result{{Opt: 1000, Arg: "5"}}, nil},
		{"permutation", []string{"10.0.0.1", "-4", "x", "-c"}, []Result{{Opt: '4'}, {Opt: 'c'}}, []string{"10.0.0.1", "x"}},
		{"double dash ends options", []string{"-4", "--", "-c", "--check"}, []Result{{Opt: '4'}}, []string{"-c", "--check"}},
		{"operands before double dash", []string{"a", "--", "-c"}, nil, []string{"a", "-c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parse(tt.args...)
			if got.err != "" {
				t.Errorf("unexpected diagnostic %q", got.err)
			}
			if !reflect.DeepEqual(got.opts, tt.opts) {
				t.Errorf("options %+v, want %+v", got.opts, tt.opts)
			}
			if len(got.rest) != len(tt.rest) || (len(tt.rest) > 0 && !reflect.DeepEqual(got.rest, tt.rest)) {
				t.Errorf("operands %q, want %q", got.rest, tt.rest)
			}
		})
	}
}

func TestDiagnostics(t *testing.T) {
	tests := []struct {
		name string
		args []string
		opts []Result
		err  string
	}{
		{"invalid short", []string{"-x"}, []Result{{Opt: '?', Optopt: 'x'}},
			"prog: invalid option -- 'x'\n"},
		{"invalid short in a bundle continues", []string{"-4x6"},
			[]Result{{Opt: '4'}, {Opt: '?', Optopt: 'x'}, {Opt: '6'}},
			"prog: invalid option -- 'x'\n"},
		{"missing short argument", []string{"-S"}, []Result{{Opt: '?', Optopt: 'S'}},
			"prog: option requires an argument -- 'S'\n"},
		{"missing short argument after bundle", []string{"-4S"},
			[]Result{{Opt: '4'}, {Opt: '?', Optopt: 'S'}},
			"prog: option requires an argument -- 'S'\n"},
		{"colon is not an option", []string{"-:"}, []Result{{Opt: '?', Optopt: ':'}},
			"prog: invalid option -- ':'\n"},
		{"question mark is not an option", []string{"-?"}, []Result{{Opt: '?', Optopt: '?'}},
			"prog: invalid option -- '?'\n"},
		{"unrecognized long", []string{"--bogus"}, []Result{{Opt: '?'}},
			"prog: unrecognized option '--bogus'\n"},
		{"ambiguous long", []string{"--net"}, []Result{{Opt: '?'}},
			"prog: option '--net' is ambiguous; possibilities: '--netmask' '--network'\n"},
		{"ambiguous with a shared prefix", []string{"--s"}, []Result{{Opt: '?'}},
			"prog: option '--s' is ambiguous; possibilities: '--split' '--split-hosts'\n"},
		{"missing long argument", []string{"--split"}, []Result{{Opt: '?', Optopt: 'S'}},
			"prog: option '--split' requires an argument\n"},
		{"missing long argument through an abbreviation", []string{"--split-"},
			[]Result{{Opt: '?', Optopt: 1000}},
			"prog: option '--split-hosts' requires an argument\n"},
		{"unexpected long argument", []string{"--check=1"}, []Result{{Opt: '?', Optopt: 'c'}},
			"prog: option '--check' doesn't allow an argument\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parse(tt.args...)
			if got.err != tt.err {
				t.Errorf("diagnostic %q, want %q", got.err, tt.err)
			}
			if !reflect.DeepEqual(got.opts, tt.opts) {
				t.Errorf("options %+v, want %+v", got.opts, tt.opts)
			}
		})
	}
}

func TestSameTargetIsNotAmbiguous(t *testing.T) {
	long := []LongOption{
		{Name: "network", Val: 'n'},
		{Name: "netaddr", Val: 'n'},
	}
	var errOut bytes.Buffer
	p := New("prog", &errOut, []string{"--net"}, "n", long)
	r, ok := p.Next()
	if !ok || r.Opt != 'n' || errOut.Len() != 0 {
		t.Errorf("got %+v %v %q", r, ok, errOut.String())
	}
}

func TestQuiet(t *testing.T) {
	var errOut bytes.Buffer
	p := New("prog", &errOut, []string{"-x", "-y"}, testShort, testLong)
	p.Next()
	p.Quiet = true
	r, _ := p.Next()
	if r.Opt != '?' || r.Optopt != 'y' {
		t.Errorf("got %+v", r)
	}
	if want := "prog: invalid option -- 'x'\n"; errOut.String() != want {
		t.Errorf("diagnostics %q, want %q", errOut.String(), want)
	}
}
