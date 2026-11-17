// SPDX-License-Identifier: BSD-3-Clause
//
// Copyright (c) 2003-2016  Simon Ekstrand
// Copyright (c) 2010-2016  Joachim Nilsson
// Copyright (c) 2016 Nikos Mavrogiannopoulos
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
// notice, this list of conditions and the following disclaimer in the
// documentation and/or other materials provided with the distribution.
// 3. The name of the author may not be used to endorse or promote products
// derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR
// IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
// OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT,
// INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
// THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//

package ipcalc

import (
	"fmt"
	"net/netip"
	"strconv"
	"strings"

	"github.com/cleonte/go-ipcalc/internal/u128"
)

type splitReq struct {
	prefix int
	count  uint64 // zero: the fill request
	hosts  uint64
}

type splitRun struct {
	start, last u128.Uint
	prefix      int
}

func (ap *app) strToU128(s string) (u128.Uint, bool) {
	if ap.flags&flagIPv6 != 0 {
		a, ok := parseIPv6Strict(s)
		if !ok {
			return u128.Uint{}, false
		}
		b := a.As16()
		var v u128.Uint
		for i := 0; i < 8; i++ {
			v.Hi = v.Hi<<8 | uint64(b[i])
		}
		for i := 8; i < 16; i++ {
			v.Lo = v.Lo<<8 | uint64(b[i])
		}
		return v, true
	}
	a, ok := parseIPv4Strict(s)
	if !ok {
		return u128.Uint{}, false
	}
	return u128.From64(uint64(ip4ToU32(a))), true
}

func addrString(v u128.Uint, v6 bool) string {
	if v6 {
		var b [16]byte
		hi, lo := v.Hi, v.Lo
		for i := 15; i >= 8; i-- {
			b[i] = byte(lo)
			lo >>= 8
		}
		for i := 7; i >= 0; i-- {
			b[i] = byte(hi)
			hi >>= 8
		}
		return ntop6(b)
	}
	return netip.AddrFrom4([4]byte{byte(v.Lo >> 24), byte(v.Lo >> 16), byte(v.Lo >> 8), byte(v.Lo)}).String()
}

// defaultPrintf prints "title: value", or only the value with --no-decorate.
func (ap *app) defaultPrintf(state *int, title, value string) {
	switch {
	case ap.flags&flagJSON != 0:
		ap.jsonPrintf(state, "", false, value)
	case ap.flags&flagNoDecorate != 0:
		ap.out.WriteString(value + "\n")
	default:
		ap.colorPrintf(colBlue, title, value)
	}
}

func (ap *app) distPrintf(state *int, title, jsonTitle, value string) {
	switch {
	case ap.flags&flagJSON != 0:
		ap.jsonPrintf(state, jsonTitle, jsonTitle != "", value)
	case ap.flags&flagNoDecorate != 0:
		ap.out.WriteString(value + "\n")
	default:
		ap.colorPrintf(colMag, title, value)
	}
}

func (ap *app) printNetwork(state *int, addr u128.Uint, prefix int) {
	s := fmt.Sprintf("%s/%d", addrString(addr, ap.v6()), prefix)
	if ap.flags&flagNoDecorate == 0 || ap.flags&flagJSON != 0 {
		ap.defaultPrintf(state, "Network:\t", s)
	} else {
		ap.out.WriteString(s + "\n")
	}
}

func largestAlignedBlock(start, last u128.Uint, width uint) uint {
	one := u128.From64(1)
	var hb uint
	for hb < width && start.And(one.Lsh(hb)).IsZero() {
		if start.Or(u128.LowMask(hb+1)).Cmp(last) > 0 {
			break
		}
		hb++
	}
	return hb
}

func (ap *app) printRange(state *int, start, last u128.Uint, width uint) {
	one := u128.From64(1)
	for {
		hb := largestAlignedBlock(start, last, width)
		ap.printNetwork(state, start, int(width-hb))
		start = start.Or(u128.LowMask(hb))
		if start.Cmp(last) >= 0 {
			break
		}
		start = start.Add(one)
	}
}

func (ap *app) printHostSizedSubnet(state *int, idx int, req *splitReq, run *splitRun, width uint) {
	hb := width - uint(run.prefix)
	one := u128.From64(1)
	mask := u128.LowMask(width).Sub(u128.LowMask(hb))

	if idx > 0 {
		ap.outputSeparate()
	}
	ap.defaultPrintf(state, fmt.Sprintf("%d. Requested size: ", idx+1), fmt.Sprintf("%d hosts", req.hosts))
	ap.printNetwork(state, run.start, run.prefix)
	ap.defaultPrintf(state, "Netmask:\t", fmt.Sprintf("%s = %d", addrString(mask, ap.v6()), run.prefix))

	if ap.flags&flagIPv6 != 0 {
		ap.defaultPrintf(state, "HostMin:\t", addrString(run.start, ap.v6()))
		ap.defaultPrintf(state, "HostMax:\t", addrString(run.last, ap.v6()))
		ap.defaultPrintf(state, "Hosts/Net:\t", ipv6Hosts(run.prefix))
	} else {
		ap.defaultPrintf(state, "Broadcast:\t", addrString(run.last, ap.v6()))
		ap.defaultPrintf(state, "HostMin:\t", addrString(run.start.Add(one), ap.v6()))
		ap.defaultPrintf(state, "HostMax:\t", addrString(run.last.Sub(one), ap.v6()))
		ap.defaultPrintf(state, "Hosts/Net:\t", ipv4Hosts(run.prefix))
	}
}

func (ap *app) tooSmall(info *ipInfo) {
	ap.errorf("ipcalc: %s/%d is too small for the requested subnets\n", info.network, info.prefix)
	ap.exit(1)
}

func (ap *app) showSplitNetworks(reqs []splitReq, info *ipInfo) {
	width := uint(32)
	if ap.flags&flagIPv6 != 0 {
		width = 128
	}
	n := len(reqs)
	fill := n
	hasSpace, hasGap, samePrefix := true, false, true
	hostSized := reqs[0].hosts != 0
	details := hostSized && ap.flags&(flagJSON|flagNoDecorate) == 0
	var state int
	one := u128.From64(1)
	var gapLast, cursor u128.Uint

	net, ok := ap.strToU128(info.network)
	if !ok {
		fam := "IPv4"
		if ap.flags&flagIPv6 != 0 {
			fam = "IPv6"
		}
		ap.errorf("ipcalc: bad %s address: %s\n", fam, info.network)
		ap.exit(1)
	}
	last := net.Or(u128.LowMask(width - uint(info.prefix)))

	runs := make([]splitRun, n)
	order := make([]int, 0, n)

	for i := 0; i < n; i++ {
		if hostSized && reqs[i].prefix < info.prefix {
			ap.tooSmall(info)
		}
		if reqs[i].prefix < info.prefix || uint(reqs[i].prefix) > width {
			ap.errorf("Cannot subnet to /%d with this base network, use a prefix > /%d\n",
				reqs[i].prefix, info.prefix)
			ap.exit(1)
		}
		if reqs[i].prefix != reqs[0].prefix {
			samePrefix = false
		}
		if reqs[i].count == 0 {
			fill = i
			continue
		}
		// stable: equal prefixes keep the request order
		j := len(order)
		order = append(order, i)
		for ; j > 0 && reqs[order[j-1]].prefix > reqs[i].prefix; j-- {
			order[j] = order[j-1]
		}
		order[j] = i
	}

	cursor = net
	for _, oi := range order {
		req := &reqs[oi]
		run := &runs[oi]
		hb := width - uint(req.prefix)

		if !hasSpace || u128.From64(req.count-1).Cmp(last.Sub(cursor).Rsh(hb)) > 0 {
			ap.tooSmall(info)
		}
		run.start = cursor
		run.last = cursor.Add(u128.From64(req.count).Lsh(hb).Sub(one))
		run.prefix = req.prefix

		if run.last.Cmp(last) == 0 {
			hasSpace = false
		} else {
			cursor = run.last.Add(one)
		}
	}

	if fill < n {
		start := cursor
		hb := width - uint(reqs[fill].prefix)
		lm := u128.LowMask(hb)
		rem := cursor.And(lm)

		if hasSpace && !rem.IsZero() {
			gap := lm.Sub(rem).Add(one)
			if last.Sub(cursor).Cmp(gap) < 0 {
				hasSpace = false
			} else {
				start = cursor.Add(gap)
				gapLast = start.Sub(one)
				hasGap = true
			}
		}
		if !hasSpace {
			ap.tooSmall(info)
		}
		runs[fill] = splitRun{start: start, last: last, prefix: reqs[fill].prefix}
	} else if hasSpace {
		gapLast = last
		hasGap = true
	}

	ap.outputStart(&state)
	ap.arrayStart(&state, "Split networks", "SPLITNETWORK")

	var netsM1 u128.Uint
	for i := 0; i < n; i++ {
		hb := width - uint(runs[i].prefix)
		lm := u128.LowMask(hb)
		step := one.Lsh(hb)

		if i > 0 {
			netsM1 = netsM1.Add(one)
		}
		netsM1 = netsM1.Add(runs[i].last.Sub(runs[i].start).Rsh(hb))

		if details {
			ap.printHostSizedSubnet(&state, i, &reqs[i], &runs[i], width)
			continue
		}
		for addr := runs[i].start; ; addr = addr.Add(step) {
			ap.printNetwork(&state, addr, runs[i].prefix)
			if addr.Add(lm).Cmp(runs[i].last) == 0 {
				break
			}
		}
	}
	ap.arrayStop(&state)

	if ap.flags&flagNoDecorate == 0 || ap.flags&flagJSON != 0 {
		ap.distPrintf(&state, "\nTotal:  \t", "NETS", netsM1.PlusOneString())
		if samePrefix {
			if ap.flags&flagIPv6 != 0 {
				ap.distPrintf(&state, "Hosts/Net:\t", "ADDRESSES", ipv6Hosts(reqs[0].prefix))
			} else {
				ap.distPrintf(&state, "Hosts/Net:\t", "ADDRESSES", ipv4Hosts(reqs[0].prefix))
			}
		}
		if hostSized {
			var neededM1 u128.Uint
			if hasSpace {
				neededM1 = cursor.Sub(net).Sub(one)
			} else {
				neededM1 = last.Sub(net)
			}
			hb := uint(0)
			for hb < width && u128.LowMask(hb).Cmp(neededM1) < 0 {
				hb++
			}
			ap.distPrintf(&state, "Needed size:\t", "NEEDED", neededM1.PlusOneString())
			ap.distPrintf(&state, "Used network:\t", "USEDNETWORK",
				fmt.Sprintf("%s/%d", addrString(net, ap.v6()), width-hb))
		}
		if hasGap {
			ap.outputSeparate()
			ap.arrayStart(&state, "Unused networks", "UNUSED")
			ap.printRange(&state, cursor, gapLast, width)
			ap.arrayStop(&state)
		}
	}
	ap.outputStop()
}

func isDigits(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

// [COUNT:]PREFIX, where PREFIX may also be a netmask
func (ap *app) parseSplitReq(str string) (splitReq, bool) {
	var req splitReq
	prefixStr := str
	if i := strings.IndexByte(str, ':'); i >= 0 {
		cnt := str[:i]
		v, err := strconv.ParseUint(cnt, 10, 64)
		if !isDigits(cnt) || err != nil || v == 0 {
			ap.errorf("ipcalc: bad split count: %s\n", str)
			return req, false
		}
		req.count = v
		prefixStr = str[i+1:]
	}
	p := ap.strToPrefix(prefixStr, true)
	if p < 0 {
		fam := "IPv4"
		if ap.flags&flagIPv6 != 0 {
			fam = "IPv6"
		}
		ap.errorf("ipcalc: bad %s prefix: %s\n", fam, str)
		return req, false
	}
	req.prefix = p
	return req, true
}

func (ap *app) parseSplitHosts(str string, reqs []splitReq) ([]splitReq, bool) {
	width := 32
	if ap.flags&flagIPv6 != 0 {
		width = 128
	}
	for _, item := range strings.Split(str, ",") {
		hosts, err := strconv.ParseUint(item, 10, 64)
		if !isDigits(item) || err != nil || hosts == 0 {
			ap.errorf("ipcalc: bad host count: %s\n", str)
			return nil, false
		}
		size := hosts
		if ap.flags&flagIPv6 == 0 {
			if hosts > 0xFFFFFFFF-1 {
				ap.errorf("ipcalc: too many hosts requested: %d\n", hosts)
				return nil, false
			}
			size += 2
		}
		b := 0
		for b < 64 && uint64(1)<<uint(b) < size {
			b++
		}
		reqs = append(reqs, splitReq{prefix: width - b, count: 1, hosts: hosts})
	}
	return reqs, true
}
