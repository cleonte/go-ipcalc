package ipcalc

import (
	"context"
	"fmt"
	"net"
	"net/netip"
	"strings"
)

func calcReverseDNS4(ip uint32, prefix int, network, broadcast uint32) []string {
	b1, b2, b3, b4 := ip>>24&0xff, ip>>16&0xff, ip>>8&0xff, ip&0xff
	var s string
	switch {
	case prefix == 32:
		s = fmt.Sprintf("%d.%d.%d.%d.in-addr.arpa.", b4, b3, b2, b1)
	case prefix == 24:
		s = fmt.Sprintf("%d.%d.%d.in-addr.arpa.", b3, b2, b1)
	case prefix == 16:
		s = fmt.Sprintf("%d.%d.in-addr.arpa.", b2, b1)
	case prefix == 8:
		s = fmt.Sprintf("%d.in-addr.arpa.", b1)
	case prefix > 24:
		s = fmt.Sprintf("%d-%d.%d.%d.%d.in-addr.arpa.", network&0xff, broadcast&0xff, b3, b2, b1)
	case prefix > 16:
		s = fmt.Sprintf("%d-%d.%d.%d.in-addr.arpa.", network>>8&0xff, broadcast>>8&0xff, b2, b1)
	case prefix > 8:
		s = fmt.Sprintf("%d-%d.%d.in-addr.arpa.", network>>16&0xff, broadcast>>16&0xff, b1)
	case prefix > 0:
		s = fmt.Sprintf("%d-%d.in-addr.arpa.", network>>24&0xff, broadcast>>24&0xff)
	default:
		s = "in-addr.arpa."
	}
	return []string{s}
}

func nibble(ip [16]byte, idx int) uint {
	b := uint(ip[idx/2])
	if idx%2 == 0 {
		return b >> 4
	}
	return b & 0xf
}

func calcReverseDNS6(ip [16]byte, prefix int) []string {
	nibbles := (prefix + 3) / 4
	hostBits := nibbles*4 - prefix
	count := 1 << uint(hostBits)
	names := make([]string, 0, count)
	for k := 0; k < count; k++ {
		var sb strings.Builder
		for i := nibbles; i > 0; i-- {
			val := nibble(ip, i-1)
			if i == nibbles {
				val = (val &^ uint(count-1)) + uint(k)
			}
			fmt.Fprintf(&sb, "%x.", val)
		}
		sb.WriteString("ip6.arpa.")
		names = append(names, sb.String())
	}
	return names
}

func resolveHostname(a netip.Addr) (string, bool) {
	names, err := net.DefaultResolver.LookupAddr(context.Background(), a.String())
	if err != nil || len(names) == 0 {
		return "", false
	}
	return strings.TrimSuffix(names[0], "."), true
}

// lookupIP returns the first address of the host in the given family
// (0 for any), formatted as inet_ntop would.
func resolveIP(host string, family int) (string, bool) {
	addrs, err := net.DefaultResolver.LookupIPAddr(context.Background(), host)
	if err != nil {
		return "", false
	}
	for _, a := range addrs {
		ip, ok := netip.AddrFromSlice(a.IP)
		if !ok {
			continue
		}
		ip = ip.Unmap()
		switch {
		case ip.Is4() && family != 6:
			return ip.String(), true
		case ip.Is6() && family != 4:
			return ntop6(ip.As16()), true
		}
	}
	return "", false
}
