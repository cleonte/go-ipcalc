package main

import (
	"fmt"
	"net/netip"
	"strconv"
	"strings"

	"github.com/oschwald/maxminddb-golang/v2"
)

// The database locations can be overridden at build time with
// -ldflags "-X main.countryDB=..." (the tests do), as the C build does with
// MAXMINDDB_LOCATION_*.
var (
	countryDB = "/usr/share/GeoIP/GeoLite2-Country.mmdb"
	cityDB    = "/usr/share/GeoIP/GeoLite2-City.mmdb"
	asnDB     = "/usr/share/GeoIP/GeoLite2-ASN.mmdb"
)

type geoInfo struct {
	continent, country, ccode, city, postalCode, coord string
	accuracyRadius, timeZone, asn, asOrg               string
	asNetwork, network                                 string
}

// path walks nested maps of a decoded record
func path(rec any, keys ...string) any {
	for _, k := range keys {
		m, ok := rec.(map[string]any)
		if !ok {
			return nil
		}
		rec = m[k]
	}
	return rec
}

func getString(rec any, keys ...string) string {
	s, _ := path(rec, keys...).(string)
	return s
}

func getUnsigned(rec any, keys ...string) string {
	if v, ok := path(rec, keys...).(uint64); ok {
		return strconv.FormatUint(v, 10)
	}
	return ""
}

// lookupRecord finds the record of ip when it covers the whole network
// ip/prefix, and the prefix of its network
func lookupRecord(file string, ip netip.Addr, v6text bool, prefix int) (any, int, bool) {
	db, err := maxminddb.Open(file)
	if err != nil {
		return nil, 0, false
	}
	defer db.Close()

	if ip.Is6() && db.Metadata.IPVersion == 4 {
		return nil, 0, false
	}
	res := db.Lookup(ip)
	if res.Err() != nil || !res.Found() {
		return nil, 0, false
	}
	recPrefix := res.Prefix().Bits()
	if ip.Is4() && !res.Prefix().Addr().Is4() {
		// network wider than the IPv4 subtree
		recPrefix = 0
	}
	if prefix < recPrefix {
		return nil, 0, false
	}
	var rec any
	if res.Decode(&rec) != nil {
		return nil, 0, false
	}
	return rec, recPrefix, true
}

// recordNetwork is the network ip/prefix in CIDR notation
func recordNetwork(ipStr string, prefix int) string {
	var addr netip.Addr
	var ok bool
	if strings.Contains(ipStr, ":") {
		addr, ok = parseIPv6Strict(ipStr)
	} else {
		addr, ok = parseIPv4Strict(ipStr)
	}
	if !ok {
		return ""
	}
	p, err := addr.Prefix(prefix)
	if err != nil {
		return ""
	}
	a := p.Addr()
	if a.Is6() {
		return fmt.Sprintf("%s/%d", ntop6(a.As16()), prefix)
	}
	return fmt.Sprintf("%s/%d", a, prefix)
}

// geoLookup fills the geo values; a missing database or record is not an
// error: the values it would provide stay empty.
func geoLookup(ipStr string, prefix int) geoInfo {
	var geo geoInfo
	var ip netip.Addr
	var ok bool
	v6 := strings.Contains(ipStr, ":")
	if v6 {
		ip, ok = parseIPv6Strict(ipStr)
	} else {
		ip, ok = parseIPv4Strict(ipStr)
	}
	if !ok {
		return geo
	}

	found := false
	networkPrefix := 0

	if rec, rp, ok := lookupRecord(countryDB, ip, v6, prefix); ok {
		found = true
		networkPrefix = rp
		geo.continent = getString(rec, "continent", "names", "en")
		for _, key := range []string{"country", "represented_country", "registered_country"} {
			if geo.country != "" || geo.ccode != "" {
				break
			}
			geo.country = getString(rec, key, "names", "en")
			geo.ccode = getString(rec, key, "iso_code")
		}
	}
	if rec, rp, ok := lookupRecord(cityDB, ip, v6, prefix); ok {
		found = true
		if rp > networkPrefix {
			networkPrefix = rp
		}
		geo.city = getString(rec, "city", "names", "en")
		geo.postalCode = getString(rec, "postal", "code")
		lat, ok1 := path(rec, "location", "latitude").(float64)
		lon, ok2 := path(rec, "location", "longitude").(float64)
		if ok1 && ok2 {
			geo.coord = fmt.Sprintf("%f,%f", lat, lon)
		}
		geo.accuracyRadius = getUnsigned(rec, "location", "accuracy_radius")
		geo.timeZone = getString(rec, "location", "time_zone")
	}
	if rec, rp, ok := lookupRecord(asnDB, ip, v6, prefix); ok {
		found = true
		if rp > networkPrefix {
			networkPrefix = rp
		}
		geo.asn = getUnsigned(rec, "autonomous_system_number")
		geo.asOrg = getString(rec, "autonomous_system_organization")
		geo.asNetwork = recordNetwork(ipStr, rp)
	}
	if found {
		geo.network = recordNetwork(ipStr, networkPrefix)
	}
	return geo
}
