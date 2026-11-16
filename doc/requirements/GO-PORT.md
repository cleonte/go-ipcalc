# Requirements in the Go port

These documents are copied from the C implementation of ipcalc (version
1.2.0) and describe behavior, which the Go port keeps: `go test ./...` runs
the same cases, with the same expected outputs, as the C test suite.

What differs from what the documents say:

* **Source citations** (`Source:` lines) name C files and functions. They map
  to the Go files like this:

  | C | Go (in `internal/ipcalc/`) |
  |---|----|
  | `ipcalc.c` (`main()`, option parsing, output helpers, `show_info_fields()`) | `cli.go`, `getopt.go`, `output.go`, `env.go` |
  | `ipcalc.c` (`get_ipv4_info()`, `get_ipv6_info()`, address tables) | `ipinfo.go` |
  | `ipcalc-reverse.c` | `reverse.go` |
  | `netsplit.c` | `split.go` |
  | `deaggregate.c`, `ipv6.c` | `deaggregate.go` |
  | `netcompare.c` | `compare.go` |
  | `binary.c` | `binary.go` |
  | `ipcalc-maxmind.c` | `geo.go` |

* **Not applicable to the Go port:** `REQ-GEN-BUILD-001`, `REQ-GEN-BUILD-002`
  and `REQ-GEO-BUILD-001` (meson and Makefile configurations, `-Werror`
  builds, `dlopen()` runtime linking, building without a geo backend). The
  Go port has one build, and always supports `-g`/`--geoinfo`.
* **Acceptance tests** name Meson tests; the Go equivalents have the same
  names in `cases_gen_test.go` and `ipcalc_test.go`.
* **Resolver:** `-h` and `-o` use Go's resolver instead of `getnameinfo()`
  and `getaddrinfo()`.
