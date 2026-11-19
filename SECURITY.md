# Security policy

## Supported versions

Only the latest release receives fixes.

## Reporting a vulnerability

Please report privately by opening a
[new issue](https://gitlab.com/ipcalc/go-ipcalc/-/issues/new) on GitLab with
"This issue is confidential" checked, rather than a public one. Include the
command line and input that trigger the problem. Expect a first reply within
a week.

ipcalc parses untrusted address strings and, with `-g`, reads MaxMind
database files; crashes, hangs or memory exhaustion on crafted input are in
scope. The parsing code is fuzzed (`go test -fuzz`) and `govulncheck` runs in
CI.
