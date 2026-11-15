#!/bin/sh
# Run the C ipcalc test suite (meson) against this Go binary.
# usage: scripts/run-c-suite.sh /path/to/ipcalc-c-checkout
# Needs meson+ninja; the C checkout is built with geo support disabled and its
# build/ipcalc is replaced by the Go binary (tests invoke build/ipcalc).
set -e
c_src=$(cd "$1" && pwd)
here=$(cd "$(dirname "$0")/.." && pwd)
(cd "$here" && go build -o ipcalc-go .)
cd "$c_src"
[ -d build ] || meson setup build -Duse_maxminddb=disabled
ninja -C build
cp "$here/ipcalc-go" build/ipcalc
meson test -C build --no-rebuild --num-processes 8 || true
# the same suite, with the C binary rebuilt, to see the baseline: ninja -C build
