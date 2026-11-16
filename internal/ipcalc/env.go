package ipcalc

import (
	"bufio"
	"io"
	"net/netip"
	"os"
)

// environ is everything the program takes from the process, so that tests
// can run it in-process.
type environ struct {
	Prog       string // argv[0], which getopt diagnostics start with
	Args       []string
	Getenv     func(string) (string, bool)
	Stdout     io.Writer
	Stderr     io.Writer
	StdoutTTY  bool
	StderrTTY  bool
	StdoutFile *os.File // for the terminal size, when Stdout is a terminal

	// Optional replacements, for tests
	CountryDB, CityDB, ASNDB string
	LookupHostname           func(netip.Addr) (string, bool)
	LookupIP                 func(host string, family int) (string, bool)
}

// app is the state of one run.
type app struct {
	flags       uint
	beSilent    bool
	exitFailure int
	colors      bool
	out         *bufio.Writer
	errOut      io.Writer
	env         environ

	countryDB, cityDB, asnDB string
	lookupHostname           func(netip.Addr) (string, bool)
	lookupIP                 func(host string, family int) (string, bool)
}

type exitPanic int

func isTTY(f *os.File) bool {
	st, err := f.Stat()
	return err == nil && st.Mode()&os.ModeCharDevice != 0
}

func osEnviron() environ {
	return environ{
		Prog:       os.Args[0],
		Args:       os.Args[1:],
		Getenv:     os.LookupEnv,
		Stdout:     os.Stdout,
		Stderr:     os.Stderr,
		StdoutTTY:  isTTY(os.Stdout),
		StderrTTY:  isTTY(os.Stderr),
		StdoutFile: os.Stdout,
	}
}

func orString(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

func newApp(e environ) *app {
	if e.Prog == "" {
		e.Prog = "ipcalc"
	}
	ap := &app{
		exitFailure:    1,
		out:            newOut(e.Stdout),
		errOut:         e.Stderr,
		env:            e,
		countryDB:      orString(e.CountryDB, countryDB),
		cityDB:         orString(e.CityDB, cityDB),
		asnDB:          orString(e.ASNDB, asnDB),
		lookupHostname: e.LookupHostname,
		lookupIP:       e.LookupIP,
	}
	if ap.lookupHostname == nil {
		ap.lookupHostname = resolveHostname
	}
	if ap.lookupIP == nil {
		ap.lookupIP = resolveIP
	}
	return ap
}

// run executes the program and returns its exit status.
func run(e environ) (code int) {
	ap := newApp(e)
	defer func() {
		if r := recover(); r != nil {
			ep, ok := r.(exitPanic)
			if !ok {
				panic(r)
			}
			code = int(ep)
		}
	}()
	ap.realMain()
	ap.exit(0)
	return 0
}

// exit flushes the output and ends the run with the given status.
func (ap *app) exit(code int) {
	ap.out.Flush()
	panic(exitPanic(code))
}

// Main runs the program with the process environment and returns the exit status.
func Main() int { return run(osEnviron()) }

func (ap *app) v6() bool { return ap.flags&flagIPv6 != 0 }
