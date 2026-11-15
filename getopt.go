package main

import (
	"fmt"
	"os"
	"strings"
)

// A GNU getopt_long work-alike: option permutation, bundled short options,
// unique-prefix abbreviations of long options and "--name=value".

type longOpt struct {
	name   string
	hasArg bool
	val    int
}

type optResult struct {
	c      int
	arg    string
	optopt int
}

type getopt struct {
	args     []string
	shortOpt string
	long     []longOpt
	pos      int // index into args
	nextChar int // position inside a bundled short option cluster
	Rest     []string
	done     bool
	opterr   bool
}

func newGetopt(args []string, short string, long []longOpt) *getopt {
	return &getopt{args: args, shortOpt: short, long: long, opterr: true}
}

func (g *getopt) shortSpec(c byte) (found, hasArg bool) {
	if c == ':' || c == '?' {
		return false, false
	}
	i := strings.IndexByte(g.shortOpt, c)
	if i < 0 {
		return false, false
	}
	return true, i+1 < len(g.shortOpt) && g.shortOpt[i+1] == ':'
}

func (g *getopt) warn(format string, a ...any) {
	if g.opterr {
		fmt.Fprintf(os.Stderr, "ipcalc: "+format+"\n", a...)
	}
}

// next returns ok=false when the options are exhausted.
func (g *getopt) next() (optResult, bool) {
	if g.nextChar > 0 {
		return g.nextShort(), true
	}
	for g.pos < len(g.args) {
		a := g.args[g.pos]
		switch {
		case a == "--":
			g.Rest = append(g.Rest, g.args[g.pos+1:]...)
			g.pos = len(g.args)
			return optResult{}, false
		case strings.HasPrefix(a, "--"):
			g.pos++
			return g.nextLong(a[2:]), true
		case len(a) > 1 && a[0] == '-':
			g.nextChar = 1
			return g.nextShort(), true
		default:
			g.Rest = append(g.Rest, a)
			g.pos++
		}
	}
	return optResult{}, false
}

func (g *getopt) nextShort() optResult {
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
		return optResult{c: '?', optopt: int(c)}
	}
	if !hasArg {
		return optResult{c: int(c)}
	}
	if !atEnd {
		arg := a[g.nextChar:]
		g.nextChar = 0
		g.pos++
		return optResult{c: int(c), arg: arg}
	}
	if g.pos < len(g.args) {
		arg := g.args[g.pos]
		g.pos++
		return optResult{c: int(c), arg: arg}
	}
	g.warn("option requires an argument -- '%c'", c)
	return optResult{c: '?', optopt: int(c)}
}

func (g *getopt) nextLong(body string) optResult {
	name, val, hasVal := strings.Cut(body, "=")
	var match *longOpt
	ambiguous := false
	for i := range g.long {
		l := &g.long[i]
		if l.name == name {
			match, ambiguous = l, false
			break
		}
		if strings.HasPrefix(l.name, name) {
			if match == nil {
				match = l
			} else if match.val != l.val || match.hasArg != l.hasArg {
				ambiguous = true
			}
		}
	}
	if ambiguous {
		g.warn("option '--%s' is ambiguous", name)
		return optResult{c: '?'}
	}
	if match == nil {
		g.warn("unrecognized option '--%s'", name)
		return optResult{c: '?'}
	}
	if !match.hasArg {
		if hasVal {
			g.warn("option '--%s' doesn't allow an argument", match.name)
			return optResult{c: '?', optopt: match.val}
		}
		return optResult{c: match.val}
	}
	if hasVal {
		return optResult{c: match.val, arg: val}
	}
	if g.pos < len(g.args) {
		arg := g.args[g.pos]
		g.pos++
		return optResult{c: match.val, arg: arg}
	}
	g.warn("option '--%s' requires an argument", match.name)
	return optResult{c: '?', optopt: match.val}
}

// remaining returns the non-option arguments, including any that follow
// the last option.
func (g *getopt) remaining() []string {
	return append(g.Rest, g.args[g.pos:]...)
}
