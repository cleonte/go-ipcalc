# go-ipcalc (prototype)

Go port of [ipcalc](https://gitlab.com/ipcalc/ipcalc), written to measure what a rewrite costs.
Implemented: info mode (summary, `--all-info`, every field-selecting option),
`-S`/`--split`/`--split-hosts` (`split.go`), `-d`/`--deaggregate` (`deaggregate.go`),
`--equals`/`--subnet-of`/`--overlaps` (`compare.go`), `--binary` (`binary.go`), `-c`, `--format`/`-j`/`--no-decorate`, `-r`, `-h`/`-o`, `-v`, `--help`/`--usage`,
and a GNU `getopt_long` work-alike (`getopt.go`).

Not implemented (exit status 99): nothing except geo lookup, which the C baseline was built without.

```sh
go build -o ipcalc-go .
scripts/run-c-suite.sh /path/to/ipcalc   # runs the C project's meson suite on this binary
```

## Result against the C suite (464 tests, geo disabled)

All 464 pass. Differential fuzzing against the C binary (3000 info-mode,
1500 split, 2500 deaggregate, 3000 compare and 3000 binary cases) found no
differences.

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
