package main

import (
	"fmt"
	"math/big"
	"net/netip"
	"strings"
)

type ipInfo struct {
	ip, expandedIP, expandedNetwork string
	reverseDNS                      []string
	network, broadcast, netmask     string
	wildcard, hostname              string
	hosts                           string
	prefix                          int
	hostmin, hostmax                string
	typ, class                      string
	iid, eui64, mac                 string
	macScope, macType               string
	geo                             geoInfo
}

type addrBlock struct {
	net    string
	prefix int
	name   string
}

var ipv4Blocks = []addrBlock{
	{"0.0.0.0", 32, "This host on this network"},
	{"0.0.0.0", 8, "This network"},
	{"10.0.0.0", 8, "Private Use"},
	{"100.64.0.0", 10, "Shared Address Space"},
	{"127.0.0.0", 8, "Loopback"},
	{"169.254.0.0", 16, "Link Local"},
	{"172.16.0.0", 12, "Private Use"},
	{"192.0.0.0", 29, "IPv4 Service Continuity Prefix"},
	{"192.0.0.8", 32, "IPv4 dummy address"},
	{"192.0.0.9", 32, "Port Control Protocol Anycast"},
	{"192.0.0.10", 32, "Traversal Using Relays around NAT Anycast"},
	{"192.0.0.170", 32, "NAT64/DNS64 Discovery"},
	{"192.0.0.171", 32, "NAT64/DNS64 Discovery"},
	{"192.0.0.0", 24, "IETF Protocol Assignments"},
	{"192.0.2.0", 24, "Documentation (TEST-NET-1)"},
	{"192.31.196.0", 24, "AS112-v4"},
	{"192.52.193.0", 24, "AMT"},
	{"192.88.99.2", 32, "6a44-relay anycast address"},
	{"192.88.99.0", 24, "6 to 4 Relay Anycast (Deprecated)"},
	{"192.168.0.0", 16, "Private Use"},
	{"192.175.48.0", 24, "Direct Delegation AS112 Service"},
	{"198.18.0.0", 15, "Benchmarking"},
	{"198.51.100.0", 24, "Documentation (TEST-NET-2)"},
	{"203.0.113.0", 24, "Documentation (TEST-NET-3)"},
	{"224.0.0.0", 4, "Multicast"},
	{"255.255.255.255", 32, "Limited Broadcast"},
	{"240.0.0.0", 4, "Reserved"},
}

var ipv6Blocks = []addrBlock{
	{"::1", 128, "Loopback Address"},
	{"::", 128, "Unspecified Address"},
	{"::ffff:0:0", 96, "IPv4-mapped Address"},
	{"64:ff9b::", 96, "IPv4-IPv6 Translat."},
	{"64:ff9b:1::", 48, "IPv4-IPv6 Translat."},
	{"100::", 64, "Discard-Only Address Block"},
	{"100:0:0:1::", 64, "Dummy IPv6 Prefix"},
	{"2001::", 32, "TEREDO"},
	{"2001:1::1", 128, "Port Control Protocol Anycast"},
	{"2001:1::2", 128, "Traversal Using Relays around NAT Anycast"},
	{"2001:1::3", 128, "DNS-SD Service Registration Protocol Anycast"},
	{"2001:2::", 48, "Benchmarking"},
	{"2001:3::", 32, "AMT"},
	{"2001:4:112::", 48, "AS112-v6"},
	{"2001:10::", 28, "Deprecated (previously ORCHID)"},
	{"2001:20::", 28, "ORCHIDv2"},
	{"2001:30::", 28, "Drone Remote ID Protocol Entity Tags (DETs) Prefix"},
	{"2001::", 23, "IETF Protocol Assignments"},
	{"2001:db8::", 32, "Documentation"},
	{"2002::", 16, "6to4"},
	{"2620:4f:8000::", 48, "Direct Delegation AS112 Service"},
	{"3fff::", 20, "Documentation"},
	{"5f00::", 16, "Segment Routing (SRv6) SIDs"},
	{"2000::", 3, "Global Unicast"},
	{"fc00::", 7, "Unique Local Unicast"},
	{"fe80::", 10, "Link-Scoped Unicast"},
	{"ff00::", 8, "Multicast"},
}

func prefix2mask(prefix int) uint32 {
	if prefix == 0 {
		return 0
	}
	return ^((uint32(1) << (32 - uint(prefix))) - 1)
}

func ip4ToU32(a netip.Addr) uint32 {
	b := a.As4()
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}

func u32ToStr(v uint32) string {
	return netip.AddrFrom4([4]byte{byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)}).String()
}

func ipv4NetToType(net uint32, prefix int) string {
	best := -1
	bestPrefix := 0
	for i, b := range ipv4Blocks {
		if prefix < b.prefix {
			continue
		}
		if best >= 0 && bestPrefix >= b.prefix {
			continue
		}
		blk := ip4ToU32(netip.MustParseAddr(b.net))
		if net&prefix2mask(b.prefix) == blk {
			best, bestPrefix = i, b.prefix
		}
	}
	if best < 0 {
		return "Internet"
	}
	return ipv4Blocks[best].name
}

func ipv4NetToClass(net uint32) string {
	b := net >> 24
	switch {
	case b < 128:
		return "Class A"
	case b < 192:
		return "Class B"
	case b < 224:
		return "Class C"
	case b < 240:
		return "Class D"
	}
	return "Class E"
}

func defaultIPv4Prefix(ip uint32) int {
	b := ip >> 24
	switch {
	case b < 128:
		return 8
	case b < 192:
		return 16
	}
	return 24
}

func ipv4Hosts(prefix int) string {
	if prefix >= 31 {
		return fmt.Sprint(uint64(1) << (32 - uint(prefix)))
	}
	return fmt.Sprint((uint64(1) << (32 - uint(prefix))) - 2)
}

func ipv6Hosts(prefix int) string {
	return new(big.Int).Lsh(big.NewInt(1), uint(128-prefix)).String()
}

func parseIPv4Strict(s string) (netip.Addr, bool) {
	a, err := netip.ParseAddr(s)
	if err != nil || !a.Is4() {
		return netip.Addr{}, false
	}
	return a, true
}

func parseIPv6Strict(s string) (netip.Addr, bool) {
	if !strings.Contains(s, ":") {
		return netip.Addr{}, false
	}
	a, err := netip.ParseAddr(s)
	if err != nil || a.Zone() != "" || a.Is4() {
		return netip.Addr{}, false
	}
	return a, true
}

func getIPv4Info(ipStr string, prefix int, flags uint) (*ipInfo, bool) {
	info := &ipInfo{}
	addr, ok := parseIPv4Strict(ipStr)
	if !ok {
		errorf("ipcalc: bad IPv4 address: %s\n", ipStr)
		return nil, false
	}
	ip := ip4ToU32(addr)

	if prefix < 0 {
		if flags&flagAssumeClassPrefix != 0 {
			prefix = defaultIPv4Prefix(ip)
		} else {
			prefix = 32
		}
	}
	if prefix > 32 {
		errorf("ipcalc: bad IPv4 prefix %d\n", prefix)
		return nil, false
	}

	info.ip = addr.String()
	mask := prefix2mask(prefix)
	info.netmask = u32ToStr(mask)
	info.prefix = prefix
	info.wildcard = u32ToStr(^mask)

	if prefix == 31 {
		info.broadcast = u32ToStr(0xFFFFFFFF)
	} else {
		info.broadcast = u32ToStr((ip & mask) | ^mask)
	}
	bcast := (ip & mask) | ^mask
	if prefix == 31 {
		bcast = 0xFFFFFFFF
	}
	network := ip & mask
	info.reverseDNS = calcReverseDNS4(network, prefix, network, bcast)
	info.network = u32ToStr(network)
	info.typ = ipv4NetToType(network, prefix)
	info.class = ipv4NetToClass(network)

	if prefix < 32 {
		minh := network
		if prefix <= 30 {
			minh |= 1
		}
		info.hostmin = u32ToStr(minh)
		maxh := network | ^mask
		if prefix <= 30 {
			maxh--
		}
		info.hostmax = u32ToStr(maxh)
	} else {
		info.hostmin = info.network
		info.hostmax = info.network
	}
	info.hosts = ipv4Hosts(prefix)

	if flags&flagGetGeoIP != 0 {
		info.geo = geoLookup(ipStr, prefix)
	}

	if flags&flagResolveHost != 0 && prefix == 32 {
		name, ok := lookupHostname(addr)
		if !ok {
			errorf("ipcalc: cannot find hostname for %s: Unknown host\n", ipStr)
			return nil, false
		}
		info.hostname = name
	}
	return info, true
}

// ntop6 follows the BIND/BSD inet_ntop algorithm, as libc does.
func ntop6(b [16]byte) string {
	var words [8]uint16
	for i := 0; i < 8; i++ {
		words[i] = uint16(b[2*i])<<8 | uint16(b[2*i+1])
	}
	bestBase, bestLen := -1, 0
	curBase, curLen := -1, 0
	for i := 0; i < 8; i++ {
		if words[i] == 0 {
			if curBase == -1 {
				curBase, curLen = i, 1
			} else {
				curLen++
			}
		} else if curBase != -1 {
			if bestBase == -1 || curLen > bestLen {
				bestBase, bestLen = curBase, curLen
			}
			curBase = -1
		}
	}
	if curBase != -1 && (bestBase == -1 || curLen > bestLen) {
		bestBase, bestLen = curBase, curLen
	}
	if bestBase != -1 && bestLen < 2 {
		bestBase = -1
	}

	var sb strings.Builder
	for i := 0; i < 8; i++ {
		if bestBase != -1 && i >= bestBase && i < bestBase+bestLen {
			if i == bestBase {
				sb.WriteByte(':')
			}
			continue
		}
		if i != 0 {
			sb.WriteByte(':')
		}
		if i == 6 && bestBase == 0 &&
			(bestLen == 6 || (bestLen == 5 && words[5] == 0xffff)) {
			sb.WriteString(netip.AddrFrom4([4]byte{b[12], b[13], b[14], b[15]}).String())
			break
		}
		fmt.Fprintf(&sb, "%x", words[i])
	}
	if bestBase != -1 && bestBase+bestLen == 8 {
		sb.WriteByte(':')
	}
	return sb.String()
}

func expandIPv6(b [16]byte) string {
	var sb strings.Builder
	for i := 0; i < 16; i++ {
		fmt.Fprintf(&sb, "%02x", b[i])
		if i%2 != 0 && i != 15 {
			sb.WriteByte(':')
		}
	}
	return sb.String()
}

func ipv6Mask(prefix int) [16]byte {
	var m [16]byte
	for i, j := prefix, 0; i > 0; i, j = i-8, j+1 {
		if i >= 8 {
			m[j] = 0xff
		} else {
			m[j] = byte(0xff << uint(8-i))
		}
	}
	return m
}

func ipv6NetToType(net [16]byte, prefix int) string {
	best := -1
	bestPrefix := 0
	for i, blk := range ipv6Blocks {
		if prefix < blk.prefix {
			continue
		}
		if best >= 0 && bestPrefix >= blk.prefix {
			continue
		}
		bb := netip.MustParseAddr(blk.net).As16()
		mask := ipv6Mask(blk.prefix)
		match := true
		for j := 0; j < 16; j++ {
			if net[j]&mask[j] != bb[j] {
				match = false
				break
			}
		}
		if match {
			best, bestPrefix = i, blk.prefix
		}
	}
	if best < 0 {
		return "Reserved"
	}
	return ipv6Blocks[best].name
}

func (info *ipInfo) ipv6IIDInfo(ip [16]byte) {
	iid := ip[8:]
	if ip[0]&0xe0 == 0 || ip[0] == 0xff {
		return
	}
	zero := true
	for _, v := range iid {
		if v != 0 {
			zero = false
		}
	}
	if zero {
		return
	}
	info.iid = fmt.Sprintf("%02x%02x:%02x%02x:%02x%02x:%02x%02x",
		iid[0], iid[1], iid[2], iid[3], iid[4], iid[5], iid[6], iid[7])
	if !(iid[3] == 0xff && iid[4] == 0xfe) {
		return
	}
	mac0 := iid[0] ^ 0x02
	info.eui64 = fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x:%02x:%02x",
		mac0, iid[1], iid[2], iid[3], iid[4], iid[5], iid[6], iid[7])
	info.mac = fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
		mac0, iid[1], iid[2], iid[5], iid[6], iid[7])
	info.macScope = "universal"
	if mac0&0x02 != 0 {
		info.macScope = "local"
	}
	info.macType = "individual"
	if mac0&0x01 != 0 {
		info.macType = "group"
	}
}

func getIPv6Info(ipStr string, prefix int, flags uint) (*ipInfo, bool) {
	info := &ipInfo{}
	addr, ok := parseIPv6Strict(ipStr)
	if !ok {
		errorf("ipcalc: bad IPv6 address: %s\n", ipStr)
		return nil, false
	}
	ip := addr.As16()
	info.expandedIP = expandIPv6(ip)
	info.ip = ntop6(ip)

	if prefix > 128 {
		errorf("ipcalc: bad IPv6 prefix: %d\n", prefix)
		return nil, false
	} else if prefix < 0 {
		prefix = 128
	}
	info.prefix = prefix

	mask := ipv6Mask(prefix)
	info.netmask = ntop6(mask)

	var network [16]byte
	for i := range network {
		network[i] = ip[i] & mask[i]
	}
	info.network = ntop6(network)
	info.expandedNetwork = expandIPv6(network)
	info.typ = ipv6NetToType(network, prefix)
	info.reverseDNS = calcReverseDNS6(network, prefix)
	info.ipv6IIDInfo(ip)

	if prefix < 128 {
		info.hostmin = info.network
		var max [16]byte
		for i := range max {
			max[i] = network[i] | ^mask[i]
		}
		info.hostmax = ntop6(max)
	} else {
		info.hostmin = info.network
		info.hostmax = info.network
	}
	info.hosts = ipv6Hosts(prefix)

	if flags&flagGetGeoIP != 0 {
		info.geo = geoLookup(ipStr, prefix)
	}

	if flags&flagResolveHost != 0 && prefix == 128 {
		name, ok := lookupHostname(addr)
		if !ok {
			errorf("ipcalc: cannot find hostname for %s: Unknown host\n", ipStr)
			return nil, false
		}
		info.hostname = name
	}
	return info, true
}
