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

```sh
go build -o ipcalc-go .
scripts/run-c-suite.sh /path/to/ipcalc   # runs the C project's meson suite on this binary
```

## Result against the C suite (geo enabled)

All 475 tests pass (the three `GeoLibraryMissing*` tests need the C build's
`dlopen` and do not exist for a build without runtime linking). Differential
fuzzing against the C binaries (3000 info-mode, 1500 split, 2500 deaggregate,
3000 compare, 3000 binary and 3000 geo cases) found no differences.

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
