# AGENTS.md

Guidance for AI agents (and humans) working in this repository.

## Overview

`go-ipcalc` is a Go port of [ipcalc](https://gitlab.com/ipcalc/ipcalc), a
command-line tool for IPv4/IPv6 network calculations (info, check, split,
deaggregate, compare, binary form, random networks, geo-IP). Output is
human-readable (colored), shell style (`NAME=value`), or JSON. The behavior
must stay identical to the C implementation.

## Requirements first

`doc/requirements/` is the normative description of the behavior (see
`doc/requirements/GO-PORT.md` for how it maps to the Go code). For a change
that alters observable behavior: update or add the requirement, then the tests
(`cases_gen_test.go` / `ipcalc_test.go` with an expected-output file in
`testdata/`), then the code, and keep `man/ipcalc.1.md`, `usage()` in
`main.go` and `NEWS` in agreement.

## Build and test

Only the Go toolchain is needed.

```sh
go build -o ipcalc .
go vet ./...
go test -race ./...
```

The tests run the program in-process through `run()` (`env.go`), which resets
the global state, so they must not use `t.Parallel()`. The environment,
terminal state and name resolvers are injected through `environ` and the
`lookupHostname`/`lookupIP` variables. Geo tests write MaxMind databases from
`testdata/geo-*.json` (`geodb_test.go`). To add a test, create the
expected-output file in `testdata/` and add a case in `ipcalc_test.go`.

## Architecture

* `main.go`: options, mode and format selection (`realMain()`), `usage()`.
* `getopt.go`: GNU `getopt_long` work-alike (abbreviations, bundling, `--x=y`).
* `ipinfo.go`: parsing, address-space tables, IPv4/IPv6 info (`ntop6()`
  follows libc's `inet_ntop`).
* `output.go`: output helpers, the field table and `showInfoFields()`; all
  output goes through them, and JSON separators are tracked by the state
  they thread.
* `split.go`, `deaggregate.go`, `compare.go`, `binary.go`, `reverse.go`,
  `geo.go`: one file per mode.

Exit goes through `exit()`, which flushes the buffered stdout and unwinds to
`run()`; do not call `os.Exit` elsewhere.

## Style

`gofmt`, `go vet` clean. Comments explain *why* only. Keep the buffered
stdout / immediate stderr ordering, since tests compare the merged stream.

## Commits

Describe the change in a short paragraph and add a `Signed-off-by:` line.
