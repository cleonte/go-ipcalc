#!/bin/sh
# Run the C ipcalc test suite (meson) against this Go binary.
# usage: scripts/run-c-suite.sh /path/to/ipcalc-c-checkout
# Needs meson+ninja and libmaxminddb. The C checkout is built with geo support
# and its build/ipcalc and build/tests/ipcalc-geo are replaced by Go binaries;
# the latter reads the test databases the suite generates, as the C one does.
set -e
c_src=$(cd "$1" && pwd)
here=$(cd "$(dirname "$0")/.." && pwd)
cd "$c_src"
[ -d build ] || meson setup build -Duse_maxminddb=enabled -Duse_runtime_linking=disabled
ninja -C build
dbs=$c_src/build/tests
(cd "$here" && go build -o ipcalc-go . &&
	go build -ldflags "-X main.countryDB=$dbs/geo-country.mmdb -X main.cityDB=$dbs/geo-city.mmdb -X main.asnDB=$dbs/geo-asn.mmdb" -o ipcalc-geo-go .)
cp "$here/ipcalc-go" build/ipcalc
cp "$here/ipcalc-geo-go" build/tests/ipcalc-geo
meson test -C build --no-rebuild --num-processes 8 || true
# the same suite, with the C binaries rebuilt, gives the baseline: ninja -C build
