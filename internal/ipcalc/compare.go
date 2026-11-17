// SPDX-License-Identifier: GPL-2.0-only
//
// Copyright (c) 2026 Nikos Mavrogiannopoulos
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

const (
	cmpEquals = iota
	cmpSubnetOf
	cmpOverlaps
)

type binaryNetwork struct {
	addr   [16]byte
	prefix int
	ipv6   bool
}

func sameLeadingBits(a, b *[16]byte, nbits int) bool {
	n, rem := nbits/8, uint(nbits%8)
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	if rem == 0 {
		return true
	}
	return (a[n]^b[n])&(0xff<<(8-rem))&0xff == 0
}

func (o *binaryNetwork) contains(in *binaryNetwork) bool {
	return o.ipv6 == in.ipv6 && o.prefix <= in.prefix && sameLeadingBits(&o.addr, &in.addr, o.prefix)
}

func toBinaryNetwork(info *ipInfo, fl uint) (binaryNetwork, bool) {
	n := binaryNetwork{prefix: info.prefix, ipv6: fl&flagIPv6 != 0}
	if n.ipv6 {
		a, ok := parseIPv6Strict(info.network)
		if !ok {
			return n, false
		}
		n.addr = a.As16()
	} else {
		a, ok := parseIPv4Strict(info.network)
		if !ok {
			return n, false
		}
		b := a.As4()
		copy(n.addr[:], b[:])
	}
	return n, true
}

func comparisonHolds(op int, a, b *binaryNetwork) bool {
	switch op {
	case cmpEquals:
		return a.ipv6 == b.ipv6 && a.prefix == b.prefix && a.addr == b.addr
	case cmpSubnetOf:
		return b.contains(a)
	}
	return a.contains(b) || b.contains(a)
}

func (ap *app) parseNetwork(str string, fl *uint) (*ipInfo, bool) {
	prefix := -1
	if *fl&flagIPv4 == 0 && strings.Contains(str, ":") {
		*fl |= flagIPv6
	}
	if i := strings.IndexByte(str, '/'); i >= 0 {
		prefixStr := str[i+1:]
		str = str[:i]
		// str_to_prefix() reads the global flags in C; here the caller's
		// flags are the user's, so they are swapped in for the call
		saved := ap.flags
		ap.flags = *fl
		prefix = ap.strToPrefix(prefixStr, false)
		*fl = ap.flags
		ap.flags = saved
		if prefix < 0 {
			fam := "IPv4"
			if *fl&flagIPv6 != 0 {
				fam = "IPv6"
			}
			ap.errorf("ipcalc: bad %s prefix: %s\n", fam, prefixStr)
			return nil, false
		}
	}
	if *fl&flagIPv6 != 0 {
		return ap.getIPv6Info(str, prefix, *fl)
	}
	return ap.getIPv4Info(str, prefix, *fl)
}

// compareNetworks returns the exit status, as cmp(1).
func (ap *app) compareNetworks(op int, network *ipInfo, fl uint, otherStr string, userFlags uint) int {
	other, ok := ap.parseNetwork(otherStr, &userFlags)
	if !ok {
		return ap.exitFailure
	}
	a, ok1 := toBinaryNetwork(network, fl)
	b, ok2 := toBinaryNetwork(other, userFlags)
	if !ok1 || !ok2 {
		return ap.exitFailure
	}
	holds := comparisonHolds(op, &a, &b)
	if !ap.beSilent && fl&(flagJSON|flagNoDecorate) == 0 {
		not := ""
		if !holds {
			not = "not "
		}
		var msg string
		switch op {
		case cmpSubnetOf:
			msg = fmt.Sprintf("%s/%d is %sa subnet of %s/%d", network.network, network.prefix, not, other.network, other.prefix)
		case cmpEquals:
			msg = fmt.Sprintf("The networks are %sequal", not)
		default:
			if holds {
				msg = "The networks overlap"
			} else {
				msg = "The networks do not overlap"
			}
		}
		ap.colorPrintf(colBlue, "", msg)
	}
	if holds {
		return 0
	}
	return 1
}
