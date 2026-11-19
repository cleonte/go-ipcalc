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
(`internal/ipcalc/cases_gen_test.go` / `ipcalc_test.go` with an expected-output file in
`internal/ipcalc/testdata/`), then the code, and keep `man/ipcalc.1.md`, `usage()` in
`internal/ipcalc/cli.go` and `NEWS` in agreement.

## Build and test

Only the Go toolchain is needed.

```sh
go build -o ipcalc ./cmd/ipcalc
go vet ./...
golangci-lint run ./...   # v2, configured in .golangci.yml
go test -race ./...
```

CI is GitLab (`.gitlab-ci.yml`): vet, race tests, minimum-Go tests, Windows
tests, golangci-lint, govulncheck and a `go generate` freshness check. The
GitHub repository is a push mirror that only runs a macOS test job.

The tests run the program in-process through `run()` (`env.go`). The environment, terminal state, name resolvers and geo
database paths are injected through `environ`. Geo tests write MaxMind databases from
`testdata/geo-*.json` (`geodb_test.go`). To add a test, create the
expected-output file in `testdata/` and add a case in `ipcalc_test.go`.

## Architecture

`cmd/ipcalc` is a one-line `main`; almost everything else is in `internal/ipcalc`,
and all state of a run is in the `app` struct (`env.go`), so runs do not share state.
Two self-contained helpers have their own packages: `internal/getopt` (GNU
`getopt_long` work-alike: abbreviations, bundling, `--x=y`) and `internal/u128`
(128-bit unsigned arithmetic for split and deaggregate).

* `cli.go`: options, mode and format selection (`realMain()`), `usage()`.
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
