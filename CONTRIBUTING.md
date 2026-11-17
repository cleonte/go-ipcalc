# Contributing

go-ipcalc is a Go port of [ipcalc](https://gitlab.com/ipcalc/ipcalc) and aims
to behave byte-for-byte like it. Behavior questions are settled by the
upstream C implementation and by `doc/requirements/`.

## Workflow

For any change that alters observable behavior, work in this order:

1. Update (or add) the requirement in `doc/requirements/`.
2. Write the test the requirement's acceptance criteria imply.
3. Change the code.

Keep `internal/ipcalc/cli.go` `usage()`, `man/ipcalc.1.md` and `NEWS` in step
with the code. See `AGENTS.md` for the details.

## Build and test

```sh
go build ./...
go test -race ./...
go vet ./...
go tool staticcheck ./...
go generate ./... && git diff --exit-code man/   # man page is generated
```

CI runs the same checks on Linux and macOS. Everything is plain Go; no C
compiler, Python or Perl is needed. New cases go in the table in
`internal/ipcalc/` with an expected-output file in `testdata/`.

## Pull requests

* One logical change per commit, with a message describing what changed.
* Sign off every commit (`git commit -s`): by doing so you certify the
  [Developer Certificate of Origin](https://developercertificate.org/).
* New source files carry the SPDX/GPL-2.0 header used by the existing ones.
* Contributions are licensed GPL-2.0-only, like the rest of the project.

Behavior that is wrong in upstream ipcalc as well should be reported there
first; this port tracks upstream.
