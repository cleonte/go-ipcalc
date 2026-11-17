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

## Releasing

1. Add a `* go-ipcalc X.Y.Z (released DATE)` section to `NEWS`; it becomes the
   release notes.
2. Tag and push: `git tag -a vX.Y.Z -m "go-ipcalc X.Y.Z" && git push origin vX.Y.Z`.

The release workflow builds the archives, attests them, publishes the release
and then updates the formula in `cleonte/homebrew-tap` (`brew install
cleonte/tap/go-ipcalc`) from `packaging/homebrew/`. The last step needs the
repository secret `HOMEBREW_TAP_TOKEN`, a token that can push to the tap; without
it the workflow skips the step and the formula must be updated by hand with
`packaging/homebrew/render.sh VERSION SHA256SUMS`.
