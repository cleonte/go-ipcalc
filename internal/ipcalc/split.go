package ipcalc

import (
	"fmt"
	"math/big"
	"math/bits"
	"net/netip"
	"strconv"
	"strings"
)

type splitReq struct {
	prefix int
	count  uint64 // zero: the fill request
	hosts  uint64
}

type splitRun struct {
	start, last u128
	prefix      int
}

// u128 holds an address of either family; IPv4 uses the low 32 bits.
type u128 struct{ hi, lo uint64 }

func u128From64(v uint64) u128 { return u128{0, v} }

func (a u128) add(b u128) u128 {
	lo, c := bits.Add64(a.lo, b.lo, 0)
	hi, _ := bits.Add64(a.hi, b.hi, c)
	return u128{hi, lo}
}

func (a u128) sub(b u128) u128 {
	lo, bo := bits.Sub64(a.lo, b.lo, 0)
	hi, _ := bits.Sub64(a.hi, b.hi, bo)
	return u128{hi, lo}
}

func (a u128) cmp(b u128) int {
	switch {
	case a.hi != b.hi:
		if a.hi < b.hi {
			return -1
		}
		return 1
	case a.lo != b.lo:
		if a.lo < b.lo {
			return -1
		}
		return 1
	}
	return 0
}

func (a u128) isZero() bool    { return a.hi == 0 && a.lo == 0 }
func (a u128) and(b u128) u128 { return u128{a.hi & b.hi, a.lo & b.lo} }
func (a u128) or(b u128) u128  { return u128{a.hi | b.hi, a.lo | b.lo} }

func (a u128) shl(n uint) u128 {
	switch {
	case n == 0:
		return a
	case n < 64:
		return u128{a.hi<<n | a.lo>>(64-n), a.lo << n}
	case n < 128:
		return u128{a.lo << (n - 64), 0}
	}
	return u128{}
}

func (a u128) shr(n uint) u128 {
	switch {
	case n == 0:
		return a
	case n < 64:
		return u128{a.hi >> n, a.lo>>n | a.hi<<(64-n)}
	case n < 128:
		return u128{0, a.hi >> (n - 64)}
	}
	return u128{}
}

func lowMask(n uint) u128 {
	if n >= 128 {
		return u128{^uint64(0), ^uint64(0)}
	}
	return u128From64(1).shl(n).sub(u128From64(1))
}

func (a u128) plus1Dec() string {
	v := new(big.Int).SetUint64(a.hi)
	v.Lsh(v, 64)
	v.Or(v, new(big.Int).SetUint64(a.lo))
	return v.Add(v, big.NewInt(1)).String()
}

func (ap *app) strToU128(s string) (u128, bool) {
	if ap.flags&flagIPv6 != 0 {
		a, ok := parseIPv6Strict(s)
		if !ok {
			return u128{}, false
		}
		b := a.As16()
		var v u128
		for i := 0; i < 8; i++ {
			v.hi = v.hi<<8 | uint64(b[i])
		}
		for i := 8; i < 16; i++ {
			v.lo = v.lo<<8 | uint64(b[i])
		}
		return v, true
	}
	a, ok := parseIPv4Strict(s)
	if !ok {
		return u128{}, false
	}
	return u128From64(uint64(ip4ToU32(a))), true
}

func (v u128) str(v6 bool) string {
	if v6 {
		var b [16]byte
		hi, lo := v.hi, v.lo
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
	return netip.AddrFrom4([4]byte{byte(v.lo >> 24), byte(v.lo >> 16), byte(v.lo >> 8), byte(v.lo)}).String()
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

func (ap *app) printNetwork(state *int, addr u128, prefix int) {
	s := fmt.Sprintf("%s/%d", addr.str(ap.v6()), prefix)
	if ap.flags&flagNoDecorate == 0 || ap.flags&flagJSON != 0 {
		ap.defaultPrintf(state, "Network:\t", s)
	} else {
		ap.out.WriteString(s + "\n")
	}
}

func largestAlignedBlock(start, last u128, width uint) uint {
	one := u128From64(1)
	var hb uint
	for hb < width && start.and(one.shl(hb)).isZero() {
		if start.or(lowMask(hb+1)).cmp(last) > 0 {
			break
		}
		hb++
	}
	return hb
}

func (ap *app) printRange(state *int, start, last u128, width uint) {
	one := u128From64(1)
	for {
		hb := largestAlignedBlock(start, last, width)
		ap.printNetwork(state, start, int(width-hb))
		start = start.or(lowMask(hb))
		if start.cmp(last) >= 0 {
			break
		}
		start = start.add(one)
	}
}

func (ap *app) printHostSizedSubnet(state *int, idx int, req *splitReq, run *splitRun, width uint) {
	hb := width - uint(run.prefix)
	one := u128From64(1)
	mask := lowMask(width).sub(lowMask(hb))

	if idx > 0 {
		ap.outputSeparate()
	}
	ap.defaultPrintf(state, fmt.Sprintf("%d. Requested size: ", idx+1), fmt.Sprintf("%d hosts", req.hosts))
	ap.printNetwork(state, run.start, run.prefix)
	ap.defaultPrintf(state, "Netmask:\t", fmt.Sprintf("%s = %d", mask.str(ap.v6()), run.prefix))

	if ap.flags&flagIPv6 != 0 {
		ap.defaultPrintf(state, "HostMin:\t", run.start.str(ap.v6()))
		ap.defaultPrintf(state, "HostMax:\t", run.last.str(ap.v6()))
		ap.defaultPrintf(state, "Hosts/Net:\t", ipv6Hosts(run.prefix))
	} else {
		ap.defaultPrintf(state, "Broadcast:\t", run.last.str(ap.v6()))
		ap.defaultPrintf(state, "HostMin:\t", run.start.add(one).str(ap.v6()))
		ap.defaultPrintf(state, "HostMax:\t", run.last.sub(one).str(ap.v6()))
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
	one := u128From64(1)
	var gapLast, cursor u128

	net, ok := ap.strToU128(info.network)
	if !ok {
		fam := "IPv4"
		if ap.flags&flagIPv6 != 0 {
			fam = "IPv6"
		}
		ap.errorf("ipcalc: bad %s address: %s\n", fam, info.network)
		ap.exit(1)
	}
	last := net.or(lowMask(width - uint(info.prefix)))

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

		if !hasSpace || u128From64(req.count-1).cmp(last.sub(cursor).shr(hb)) > 0 {
			ap.tooSmall(info)
		}
		run.start = cursor
		run.last = cursor.add(u128From64(req.count).shl(hb).sub(one))
		run.prefix = req.prefix

		if run.last.cmp(last) == 0 {
			hasSpace = false
		} else {
			cursor = run.last.add(one)
		}
	}

	if fill < n {
		start := cursor
		hb := width - uint(reqs[fill].prefix)
		lm := lowMask(hb)
		rem := cursor.and(lm)

		if hasSpace && !rem.isZero() {
			gap := lm.sub(rem).add(one)
			if last.sub(cursor).cmp(gap) < 0 {
				hasSpace = false
			} else {
				start = cursor.add(gap)
				gapLast = start.sub(one)
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

	var netsM1 u128
	for i := 0; i < n; i++ {
		hb := width - uint(runs[i].prefix)
		lm := lowMask(hb)
		step := one.shl(hb)

		if i > 0 {
			netsM1 = netsM1.add(one)
		}
		netsM1 = netsM1.add(runs[i].last.sub(runs[i].start).shr(hb))

		if details {
			ap.printHostSizedSubnet(&state, i, &reqs[i], &runs[i], width)
			continue
		}
		for addr := runs[i].start; ; addr = addr.add(step) {
			ap.printNetwork(&state, addr, runs[i].prefix)
			if addr.add(lm).cmp(runs[i].last) == 0 {
				break
			}
		}
	}
	ap.arrayStop(&state)

	if ap.flags&flagNoDecorate == 0 || ap.flags&flagJSON != 0 {
		ap.distPrintf(&state, "\nTotal:  \t", "NETS", netsM1.plus1Dec())
		if samePrefix {
			if ap.flags&flagIPv6 != 0 {
				ap.distPrintf(&state, "Hosts/Net:\t", "ADDRESSES", ipv6Hosts(reqs[0].prefix))
			} else {
				ap.distPrintf(&state, "Hosts/Net:\t", "ADDRESSES", ipv4Hosts(reqs[0].prefix))
			}
		}
		if hostSized {
			var neededM1 u128
			if hasSpace {
				neededM1 = cursor.sub(net).sub(one)
			} else {
				neededM1 = last.sub(net)
			}
			hb := uint(0)
			for hb < width && lowMask(hb).cmp(neededM1) < 0 {
				hb++
			}
			ap.distPrintf(&state, "Needed size:\t", "NEEDED", neededM1.plus1Dec())
			ap.distPrintf(&state, "Used network:\t", "USEDNETWORK",
				fmt.Sprintf("%s/%d", net.str(ap.v6()), width-hb))
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
