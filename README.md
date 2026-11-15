# go-ipcalc (prototype)

Go port of [ipcalc](https://gitlab.com/ipcalc/ipcalc), written to measure what a rewrite costs.
Implemented: info mode (summary, `--all-info`, every field-selecting option),
`-S`/`--split`/`--split-hosts` (`split.go`), `-d`/`--deaggregate` (`deaggregate.go`),
`--equals`/`--subnet-of`/`--overlaps` (`compare.go`), `--binary` (`binary.go`), `-g`/`--geoinfo` and `--all-info` geo values (`geo.go`), `-c`, `--format`/`-j`/`--no-decorate`, `-r`, `-h`/`-o`, `-v`, `--help`/`--usage`,
and a GNU `getopt_long` work-alike (`getopt.go`).

Geo lookup uses [maxminddb-golang](https://github.com/oschwald/maxminddb-golang), so
there is no libmaxminddb, no `dlopen` and no geo-less build. The database
paths default to `/usr/share/GeoIP/GeoLite2-*.mmdb` and are overridden at build
time, as in the C build: `-ldflags "-X main.countryDB=... -X main.cityDB=... -X main.asnDB=..."`.

## Tests

The tests are plain Go: `go test ./...`. They run the program in-process and
need nothing else, not even a C compiler or Python.

* `cases_gen_test.go` holds the 446 cases translated from the C project's
  Meson test list (`tests/meson.build`); their expected outputs are the files
  in `testdata/`. The few tests that were shell pipelines (random networks,
  `jq` and `eval` checks, the Debian comparison, the `LD_PRELOAD` fakes) are
  written as Go tests in `ipcalc_test.go`.
* `geodb_test.go` writes the MaxMind test databases from
  `testdata/geo-*.json` (a Go port of `mkmmdb.py`).
* Name lookups are replaced by fakes, so no test depends on DNS.
* `testdata/debian/` has the output of the Debian ipcalc for the
  `--split-hosts` comparison, recorded once, so Perl is not needed.

```sh
go build -o ipcalc .
go test ./...
```

`scripts/run-c-suite.sh` is an optional cross-check that runs the C project's
own Meson suite against the Go binary (needs meson, ninja and libmaxminddb).

## Result

`go test ./...` passes (479 tests). The C suite (475 tests, geo enabled) also
passes against the Go binaries, and differential fuzzing against the C
binaries (info, split, deaggregate, compare, binary and geo, 16000 cases)
found no differences.

## Known, intentional differences

* `getopt` diagnostics use glibc wording (`option '--x' requires an argument`);
  the C build on macOS prints the BSD wording.
* `inet_pton` strictness follows glibc/`net/netip`: leading-zero octets
  (`01.2.3.4`) and IPv6 zone ids (`fe80::1%eth0`) are rejected. macOS libc accepts them.

## License and attribution

A derivative work of ipcalc, licensed under GPL-2.0 (see `COPYING`).
Copyright (c) 1997-2015 Red Hat, Inc.; (c) 2026 Nikos Mavrogiannopoulos and
the other ipcalc authors. `split.go` follows `netsplit.c`, which is
BSD-3-Clause licensed: (c) 2003-2016 Simon Ekstrand, 2010-2016 Joachim
Nilsson, 2016 Nikos Mavrogiannopoulos. This is an unofficial prototype and
not affiliated with the ipcalc project.
