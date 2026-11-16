package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

// The goldens in testdata/color are the output of the C ipcalc on a
// terminal (recorded under a pty); the tests run without one, so this is
// the only coverage of the escape sequences, which the C suite lacks too.
func TestColorOutput(t *testing.T) {
	for _, c := range colorCases {
		t.Run(c.name, func(t *testing.T) {
			want, err := os.ReadFile(filepath.Join("testdata", "color", c.name))
			if err != nil {
				t.Fatal(err)
			}
			env := map[string]string{"COLUMNS": "80"}
			for k, v := range c.env {
				env[k] = v
			}
			fakeResolvers(t, "localhost")
			r := ipcalcTTY(t, opts{env: env, geo: c.bin == "geo"}, c.args...)
			if r.code != 0 {
				t.Fatalf("exit status %d: %s", r.code, r.stderr)
			}
			if !bytes.Equal([]byte(r.stdout), want) {
				t.Errorf("output differs\ngot:\n%q\nwant:\n%q", r.stdout, want)
			}
		})
	}
}

// Without a terminal, or with NO_COLOR, there are no escape sequences.
func TestNoColorWithoutTerminal(t *testing.T) {
	for _, args := range [][]string{
		{"--all-info", "192.168.2.7/24"},
		{"--binary", "2001:db8::/48"},
		{"--all-info", "2001:db8::1/64"},
	} {
		if r := ipcalc(t, opts{env: map[string]string{"COLUMNS": "80"}}, args...); bytes.Contains([]byte(r.stdout), []byte("\x1b[")) {
			t.Errorf("%v: escape sequence without a terminal: %q", args, r.stdout)
		}
	}
}
