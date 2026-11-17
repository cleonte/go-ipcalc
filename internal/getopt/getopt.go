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

// Package getopt is a GNU getopt_long work-alike: option permutation, bundled
// short options, unique-prefix abbreviations of long options and
// "--name=value".
package getopt

import (
	"fmt"
	"io"
	"strings"
)

// LongOption describes a long option; Val is returned for it in Result.Opt.
type LongOption struct {
	Name   string
	HasArg bool
	Val    int
}

// Result is one parsed option. Opt is the short option character or the Val
// of the long option, or '?' for an error; in that case Optopt is the
// offending short option or the Val of the long option, if known.
type Result struct {
	Opt    int
	Arg    string
	Optopt int
}

// Parser iterates over the options of an argument list.
type Parser struct {
	// Quiet suppresses the diagnostics of later errors.
	Quiet bool

	prog     string
	errOut   io.Writer
	args     []string
	shortOpt string
	long     []LongOption
	pos      int // index into args
	nextChar int // position inside a bundled short option cluster
	rest     []string
}

// New returns a Parser for args (without the program name). short is a
// getopt(3) option string; diagnostics are written to errOut prefixed with
// prog.
func New(prog string, errOut io.Writer, args []string, short string, long []LongOption) *Parser {
	return &Parser{prog: prog, errOut: errOut, args: args, shortOpt: short, long: long}
}

func (g *Parser) shortSpec(c byte) (found, hasArg bool) {
	if c == ':' || c == '?' {
		return false, false
	}
	i := strings.IndexByte(g.shortOpt, c)
	if i < 0 {
		return false, false
	}
	return true, i+1 < len(g.shortOpt) && g.shortOpt[i+1] == ':'
}

func (g *Parser) warn(format string, a ...any) {
	if !g.Quiet {
		fmt.Fprintf(g.errOut, g.prog+": "+format+"\n", a...)
	}
}

// Next returns the next option; ok is false when the options are exhausted.
func (g *Parser) Next() (Result, bool) {
	if g.nextChar > 0 {
		return g.nextShort(), true
	}
	for g.pos < len(g.args) {
		a := g.args[g.pos]
		switch {
		case a == "--":
			g.rest = append(g.rest, g.args[g.pos+1:]...)
			g.pos = len(g.args)
			return Result{}, false
		case strings.HasPrefix(a, "--"):
			g.pos++
			return g.nextLong(a[2:]), true
		case len(a) > 1 && a[0] == '-':
			g.nextChar = 1
			return g.nextShort(), true
		default:
			g.rest = append(g.rest, a)
			g.pos++
		}
	}
	return Result{}, false
}

func (g *Parser) nextShort() Result {
	a := g.args[g.pos]
	c := a[g.nextChar]
	g.nextChar++
	atEnd := g.nextChar >= len(a)
	if atEnd {
		g.nextChar = 0
		g.pos++
	}
	found, hasArg := g.shortSpec(c)
	if !found {
		g.warn("invalid option -- '%c'", c)
		return Result{Opt: '?', Optopt: int(c)}
	}
	if !hasArg {
		return Result{Opt: int(c)}
	}
	if !atEnd {
		arg := a[g.nextChar:]
		g.nextChar = 0
		g.pos++
		return Result{Opt: int(c), Arg: arg}
	}
	if g.pos < len(g.args) {
		arg := g.args[g.pos]
		g.pos++
		return Result{Opt: int(c), Arg: arg}
	}
	g.warn("option requires an argument -- '%c'", c)
	return Result{Opt: '?', Optopt: int(c)}
}

func (g *Parser) nextLong(body string) Result {
	name, val, hasVal := strings.Cut(body, "=")
	var match *LongOption
	var candidates []string
	ambiguous := false
	for i := range g.long {
		l := &g.long[i]
		if l.Name == name {
			match, ambiguous = l, false
			break
		}
		if strings.HasPrefix(l.Name, name) {
			candidates = append(candidates, "'--"+l.Name+"'")
			if match == nil {
				match = l
			} else if match.Val != l.Val || match.HasArg != l.HasArg {
				ambiguous = true
			}
		}
	}
	if ambiguous {
		g.warn("option '--%s' is ambiguous; possibilities: %s", name, strings.Join(candidates, " "))
		return Result{Opt: '?'}
	}
	if match == nil {
		g.warn("unrecognized option '--%s'", name)
		return Result{Opt: '?'}
	}
	if !match.HasArg {
		if hasVal {
			g.warn("option '--%s' doesn't allow an argument", match.Name)
			return Result{Opt: '?', Optopt: match.Val}
		}
		return Result{Opt: match.Val}
	}
	if hasVal {
		return Result{Opt: match.Val, Arg: val}
	}
	if g.pos < len(g.args) {
		arg := g.args[g.pos]
		g.pos++
		return Result{Opt: match.Val, Arg: arg}
	}
	g.warn("option '--%s' requires an argument", match.Name)
	return Result{Opt: '?', Optopt: match.Val}
}

// Remaining returns the non-option arguments, including any that follow
// the last option.
func (g *Parser) Remaining() []string {
	return append(g.rest, g.args[g.pos:]...)
}
