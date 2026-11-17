# go-ipcalc

A Go port of [ipcalc](https://gitlab.com/ipcalc/ipcalc): IPv4/IPv6 network
calculations (info, check, split, deaggregate, compare, binary view, random
networks) with optional geo-IP lookup, printed as colored text, `NAME=value`
shell variables or JSON. It is a single static binary (`CGO_ENABLED=0`) that
builds for Linux, macOS, FreeBSD and Windows.

The goal is behavior identical to the C ipcalc (currently 1.2.0): same
options, same output bytes, same exit status, same diagnostics. The test
suite is the C project's suite translated to Go, and the port was checked
against the C binary with differential fuzzing.

## Install

Download an archive from the
[releases page](https://github.com/cleonte/go-ipcalc/releases) (verify it
against `SHA256SUMS`), or build from source:

```sh
go install github.com/cleonte/go-ipcalc/cmd/ipcalc@latest
```

```
$ ipcalc 192.168.1.10/24
Address:	192.168.1.10
Network:	192.168.1.0/24
Netmask:	255.255.255.0 = 24
Broadcast:	192.168.1.255
...
```

See `ipcalc --help` and `man/ipcalc.1`.

## Geo lookup

`-g`/`--geoinfo` and `--all-info` read MaxMind databases through
[maxminddb-golang](https://github.com/oschwald/maxminddb-golang); there is no
libmaxminddb, no `dlopen` and no build without geo support. The database paths
default to `/usr/share/GeoIP/GeoLite2-*.mmdb` and are overridden at build time:
`-ldflags "-X github.com/cleonte/go-ipcalc/internal/ipcalc.countryDB=... -X ...cityDB=... -X ...asnDB=..."`.

## Tests

The tests are plain Go: `go test ./...`. They run the program in-process and
need nothing else, not even a C compiler or Python.

* `internal/ipcalc/cases_gen_test.go` holds the 446 cases translated from the C project's
  Meson test list (`tests/meson.build`); their expected outputs are the files
  in `internal/ipcalc/testdata/`. The few tests that were shell pipelines (random networks,
  `jq` and `eval` checks, the Debian comparison, the `LD_PRELOAD` fakes) are
  written as Go tests in `internal/ipcalc/ipcalc_test.go`.
* `internal/ipcalc/geodb_test.go` writes the MaxMind test databases from
  `internal/ipcalc/testdata/geo-*.json` (a Go port of `mkmmdb.py`).
* Name lookups are replaced by fakes, so no test depends on DNS.
* `internal/ipcalc/testdata/debian/` has the output of the Debian ipcalc for the
  `--split-hosts` comparison, recorded once, so Perl is not needed.

```sh
go install github.com/cleonte/go-ipcalc/cmd/ipcalc@latest   # or: go build -o ipcalc ./cmd/ipcalc
go test ./...
```

## Verification

`go test ./...` passes. The C suite (475 tests, geo enabled) also
passes against the Go binaries, and differential fuzzing against the C
binaries (info, split, deaggregate, compare, binary and geo, 16000 cases)
found no differences.

## Known, intentional differences

* `getopt` diagnostics use glibc wording (`option '--x' requires an argument`);
  the C build on macOS prints the BSD wording.
* `inet_pton` strictness follows glibc/`net/netip`: leading-zero octets
  (`01.2.3.4`) and IPv6 zone ids (`fe80::1%eth0`) are rejected. macOS libc accepts them.

## Contributing and security

See `CONTRIBUTING.md` and `SECURITY.md`. Requirements live in
`doc/requirements/`; `doc/requirements/GO-PORT.md` maps their C source
citations to the Go files.

## License and attribution

A derivative work of ipcalc, licensed under GPL-2.0-only (see `COPYING`;
each source file carries an SPDX header). Copyright (c) 1997-2015 Red Hat,
Inc.; (c) 2026 Nikos Mavrogiannopoulos and the other ipcalc authors.
`split.go` follows `netsplit.c`, which is BSD-3-Clause licensed:
(c) 2003-2016 Simon Ekstrand, 2010-2016 Joachim Nilsson, 2016 Nikos
Mavrogiannopoulos. This is an independent port, not an official part of the
ipcalc project.
