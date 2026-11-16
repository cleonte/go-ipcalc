package main

import (
	"io"
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
}

var (
	env    environ
	errOut io.Writer
)

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

// run executes the program and returns its exit status. The global state
// is reset first, so that it can be called repeatedly.
func run(e environ) (code int) {
	env = e
	if env.Prog == "" {
		env.Prog = "ipcalc"
	}
	errOut = e.Stderr
	out = newOut(e.Stdout)
	beSilent = false
	exitFailure = 1
	colors = false
	flags = 0

	defer func() {
		if r := recover(); r != nil {
			ep, ok := r.(exitPanic)
			if !ok {
				panic(r)
			}
			code = int(ep)
		}
	}()
	realMain()
	exit(0)
	return 0
}

// exit flushes the output and ends the run with the given status.
func exit(code int) {
	out.Flush()
	panic(exitPanic(code))
}

func main() {
	os.Exit(run(osEnviron()))
}
