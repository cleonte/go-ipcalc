# go-ipcalc

[![pipeline](https://gitlab.com/ipcalc/go-ipcalc/badges/master/pipeline.svg)](https://gitlab.com/ipcalc/go-ipcalc/-/pipelines) [![Release](https://gitlab.com/ipcalc/go-ipcalc/-/badges/release.svg)](https://gitlab.com/ipcalc/go-ipcalc/-/releases) [![Go Reference](https://pkg.go.dev/badge/gitlab.com/ipcalc/go-ipcalc.svg)](https://pkg.go.dev/gitlab.com/ipcalc/go-ipcalc) [![License: GPL-2.0](https://img.shields.io/badge/license-GPL--2.0-blue.svg)](COPYING)

A Go port of [ipcalc](https://gitlab.com/ipcalc/ipcalc): IPv4/IPv6 network
calculations (info, check, split, deaggregate, compare, binary view, random
networks) with optional geo-IP lookup, printed as colored text, `NAME=value`
shell variables or JSON. It is a single static binary (`CGO_ENABLED=0`) that
builds for Linux, macOS, FreeBSD and Windows.

The project lives at <https://gitlab.com/ipcalc/go-ipcalc>; report issues and
send merge requests there. The GitHub repository is a read-only mirror.

The goal is behavior identical to the C ipcalc (currently 1.2.0): same
options, same output bytes, same exit status, same diagnostics. The test
suite is the C project's suite translated to Go, and the port was checked
against the C binary with differential fuzzing.

## Install

Homebrew:

```sh
brew install cleonte/tap/go-ipcalc
```

Linux packages (`.deb`, `.rpm`, `.apk`, named `go-ipcalc`) are on the
[releases page](https://gitlab.com/ipcalc/go-ipcalc/releases):

```sh
sudo apt install ./go-ipcalc_<version>_linux_amd64.deb
sudo dnf install ./go-ipcalc_<version>_linux_amd64.rpm
sudo apk add --allow-untrusted ./go-ipcalc_<version>_linux_amd64.apk
```

They conflict with the distribution's `ipcalc` package, which installs the
same `/usr/bin/ipcalc`. Archives for other systems are on the same page; verify
a download against `SHA256SUMS`. Or build from source:

```sh
go install gitlab.com/ipcalc/go-ipcalc/cmd/ipcalc@latest
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
`-ldflags "-X gitlab.com/ipcalc/go-ipcalc/internal/ipcalc.countryDB=... -X ...cityDB=... -X ...asnDB=..."`.

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
go install gitlab.com/ipcalc/go-ipcalc/cmd/ipcalc@latest   # or: go build -o ipcalc ./cmd/ipcalc
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
`split.go` and `internal/u128` follow `netsplit.c`, which is BSD-3-Clause licensed:
(c) 2003-2016 Simon Ekstrand, 2010-2016 Joachim Nilsson, 2016 Nikos
Mavrogiannopoulos. This is an independent port, not an official part of the
ipcalc project.
