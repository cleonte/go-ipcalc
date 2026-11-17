// SPDX-License-Identifier: GPL-2.0-only
//
// Copyright (c) 2019 Nikos Mavrogiannopoulos
// Copyright (c) 1991-1997, 1999-2014 Free Software Foundation, Inc.
// Copyright (c) 2026 go-ipcalc contributors
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License, version 2,
// as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

package ipcalc

import (
	"fmt"
	"strings"
)

func isSpace(c byte) bool {
	return c == ' ' || (c >= '\t' && c <= '\r')
}

// strtok(str, "-") yields the non-empty pieces between dashes
func (ap *app) deaggregate(str string) {
	k := 0
	for k < len(str) && str[k] == '-' {
		k++
	}
	if k == len(str) {
		ap.errorf("ipcalc: bad deaggregation string: %s\n", str)
		ap.exit(1)
	}
	end1 := k + strings.IndexByte(str[k:]+"-", '-')
	d1 := str[k:end1]
	// strtok wrote a NUL after the first token, which truncates str
	trunc := str[:end1]

	rest := ""
	if end1 < len(str) {
		rest = str[end1+1:]
	}
	rest = strings.TrimLeft(rest, "-")
	if rest == "" {
		ap.errorf("ipcalc: bad deaggregation string: %s\n", trunc)
		ap.exit(1)
	}
	d2 := rest
	if i := strings.IndexByte(rest, '-'); i >= 0 {
		d2 = rest[:i]
	}

	d1 = strings.TrimFunc(d1, func(r rune) bool { return r < 128 && isSpace(byte(r)) })
	d2 = strings.TrimFunc(d2, func(r rune) bool { return r < 128 && isSpace(byte(r)) })

	if ap.flags&flagIPv6 != 0 {
		ap.deaggregateV6(d1, d2)
	} else {
		ap.deaggregateV4(d1, d2)
	}
}

func (ap *app) deaggregateV4(s1, s2 string) {
	a1, ok := parseIPv4Strict(s1)
	if !ok {
		ap.errorf("ipcalc: bad IPv4 address: %s\n", s1)
		ap.exit(1)
	}
	a2, ok := parseIPv4Strict(s2)
	if !ok {
		ap.errorf("ipcalc: bad IPv4 address: %s\n", s2)
		ap.exit(1)
	}
	base, end := uint64(ip4ToU32(a1)), uint64(ip4ToU32(a2))
	if base > end {
		ap.errorf("ipcalc: bad range\n")
		ap.exit(1)
	}

	var state int
	ap.outputStart(&state)
	ap.arrayStart(&state, "Deaggregated networks", "DEAGGREGATEDNETWORK")
	for base <= end {
		step := uint(0)
		for step < 32 && base|(uint64(1)<<step) != base {
			b1 := base | (uint64(0xFFFFFFFF) >> (31 - step))
			if b1 > end {
				break
			}
			step++
		}
		ap.defaultPrintf(&state, "Network:\t", fmt.Sprintf("%s/%d", u32ToStr(uint32(base)), 32-step))
		base += uint64(1) << step
	}
	ap.arrayStop(&state)
	ap.outputStop()
}

func (ap *app) deaggregateV6(s1, s2 string) {
	if _, ok := parseIPv6Strict(s1); !ok {
		ap.errorf("ipcalc: bad IPv6 address: %s\n", s1)
		ap.exit(1)
	}
	if _, ok := parseIPv6Strict(s2); !ok {
		ap.errorf("ipcalc: bad IPv6 address: %s\n", s2)
		ap.exit(1)
	}
	base, _ := ap.strToU128(s1)
	end, _ := ap.strToU128(s2)
	if base.cmp(end) > 0 {
		ap.errorf("ipcalc: bad IPv6 range\n")
		ap.exit(1)
	}

	one := u128From64(1)
	var state int
	ap.outputStart(&state)
	ap.arrayStart(&state, "Deaggregated networks", "DEAGGREGATEDNETWORK")
	for {
		step := uint(0)
		for step < 128 && base.and(one.shl(step)).isZero() {
			if base.or(lowMask(step+1)).cmp(end) > 0 {
				break
			}
			step++
		}
		ap.defaultPrintf(&state, "Network:\t", fmt.Sprintf("%s/%d", base.str(ap.v6()), 128-step))

		// stop at the network holding the end; advancing past the last
		// address of the IPv6 space would wrap around
		tmp := base.or(lowMask(step))
		if tmp.cmp(end) >= 0 {
			break
		}
		base = tmp.add(one)
	}
	ap.arrayStop(&state)
	ap.outputStop()
}
