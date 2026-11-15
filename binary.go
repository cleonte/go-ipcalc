package main

import (
	"fmt"
	"net/netip"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

const (
	maxCols       = 135
	defaultCols   = 80
	labelCols     = 12
	detailsIndent = labelCols + 17
)

const (
	colDim = "\x1b[2m"

	glyphLeft  = 1
	glyphRight = 2
	glyphFill  = 3
)

const (
	stylePlain = iota
	styleNetwork
	styleHost
)

type layout struct {
	width, bandBits, groupBits int
	sep                        byte
	hex                        bool
}

var (
	ipv4Layout     = layout{32, 32, 8, '.', false}
	ipv6Layout     = layout{128, 64, 16, ':', true}
	ipv6WideLayout = layout{128, 128, 16, ':', true}
)

type segment struct {
	first, end int
	style      int
	labels     []string
}

type row struct {
	name, vr string
	value    uint32
}

type line struct {
	indent int
	n      int
	text   [maxCols]byte
	style  [maxCols]byte
}

func colorFputs(color, s string) {
	if colors && color != "" {
		out.WriteString(color)
	}
	out.WriteString(s)
	if colors && color != "" {
		out.WriteString(colReset)
	}
}

func getBit(addr []byte, k int) int { return int(addr[k/8]>>(7-uint(k%8))) & 1 }

func newLine(indent int) *line {
	l := &line{indent: indent}
	for i := range l.text {
		l.text[i] = ' '
	}
	return l
}

func (l *line) putChar(col int, c byte, style int) {
	if col >= maxCols {
		return
	}
	l.text[col] = c
	l.style[col] = byte(style)
	if col >= l.n {
		l.n = col + 1
	}
}

func (l *line) putText(col int, s string, style int) {
	for i := 0; i < len(s); i++ {
		l.putChar(col+i, s[i], style)
	}
}

func styleColor(s byte) string {
	switch s {
	case styleNetwork:
		return colBlue
	case styleHost:
		return colDim
	}
	return ""
}

func glyphText(c byte) string {
	switch c {
	case glyphLeft:
		return "└"
	case glyphRight:
		return "┘"
	case glyphFill:
		return "─"
	}
	return ""
}

func (l *line) print() {
	for l.n > 0 && l.text[l.n-1] == ' ' {
		l.n--
	}
	if l.n > 0 {
		colorFputs("", strings.Repeat(" ", l.indent))
	}
	col := 0
	for col < l.n {
		var run strings.Builder
		st := l.style[col]
		for ; col < l.n && l.style[col] == st; col++ {
			if g := glyphText(l.text[col]); g != "" {
				run.WriteString(g)
			} else {
				run.WriteByte(l.text[col])
			}
		}
		colorFputs(styleColor(st), run.String())
	}
	colorFputs("", "\n")
}

func (lay *layout) bandColumns() int {
	return lay.bandBits + lay.bandBits/lay.groupBits - 1
}

func outputColumns() int {
	if env, ok := os.LookupEnv("COLUMNS"); ok {
		if v, ok := safeAtoi(env); ok && v > 0 {
			return v
		}
	}
	if isTTY(os.Stdout) {
		var ws [4]uint16
		_, _, e := syscall.Syscall(syscall.SYS_IOCTL, os.Stdout.Fd(), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&ws)))
		if e == 0 && ws[1] > 0 {
			return int(ws[1])
		}
	}
	return defaultCols
}

func (lay *layout) groupColumn(g int) int { return g * (lay.groupBits + 1) }

func (lay *layout) bitColumn(start, k int) int {
	pos := k - start
	return lay.groupColumn(pos/lay.groupBits) + pos%lay.groupBits
}

func groupValue(addr []byte, first, bits int) uint {
	var v uint
	for k := first; k < first+bits; k++ {
		v = v<<1 | uint(getBit(addr, k))
	}
	return v
}

func groupHostBits(prefix, first, bits int) int {
	end := first + bits
	if prefix >= end {
		return 0
	}
	if prefix <= first {
		return bits
	}
	return end - prefix
}

func (lay *layout) formatGroup(value uint, hostBits int) string {
	high := value | (uint(1)<<uint(hostBits) - 1)
	switch {
	case hostBits == 0 && lay.hex:
		return fmt.Sprintf("%04x", value)
	case hostBits == 0:
		return fmt.Sprintf("%d", value)
	case lay.hex:
		return fmt.Sprintf("%04x-%04x", value, high)
	}
	return fmt.Sprintf("%d-%d", value, high)
}

func printBand(addr []byte, prefix, start int, lay *layout, indent int) {
	gbits := lay.groupBits
	ngroups := lay.bandBits / gbits
	sep := string(lay.sep)

	l := newLine(indent)
	for g := 0; g < ngroups; g++ {
		first := start + g*gbits
		col := lay.groupColumn(g)
		hb := groupHostBits(prefix, first, gbits)
		st := styleNetwork
		if hb > 0 {
			st = styleHost
		}
		s := lay.formatGroup(groupValue(addr, first, gbits), hb)
		l.putText(col+gbits-len(s), s, st)
		if g+1 < ngroups {
			l.putText(col+gbits, sep, stylePlain)
		}
	}
	l.print()

	l = newLine(indent)
	for g := 0; g < ngroups; g++ {
		first := start + g*gbits
		col := lay.groupColumn(g)
		for i := 0; i < gbits; i++ {
			k := first + i
			if k < prefix {
				l.putChar(col+i, byte('0'+getBit(addr, k)), styleNetwork)
			} else {
				l.putChar(col+i, 'x', styleHost)
			}
		}
		if g+1 < ngroups {
			l.putText(col+gbits, sep, stylePlain)
		}
	}
	l.print()
}

func (l *line) putSegmentEnds(first, last int, left, right bool, style int) {
	if left {
		l.putChar(first, glyphLeft, style)
	}
	if right {
		l.putChar(last, glyphRight, style)
	}
}

func (l *line) putSegment(first, last int, left, right bool, seg *segment) {
	width := last - first + 1
	ends := 0
	if left {
		ends++
	}
	if right {
		ends++
	}
	inner := 0
	if width > ends {
		inner = width - ends
	}
	shortest := seg.labels[len(seg.labels)-1]

	i := 0
	for ; i < len(seg.labels); i++ {
		if len(seg.labels[i]) <= inner {
			break
		}
	}
	if i == len(seg.labels) {
		if len(shortest) <= width {
			l.putText(first+(width-len(shortest))/2, shortest, seg.style)
		} else if ends <= width {
			l.putSegmentEnds(first, last, left, right, seg.style)
		}
		return
	}

	text := seg.labels[i]
	if len(text)+2 <= inner {
		text = " " + text + " "
	}
	for col := first; col <= last; col++ {
		l.putChar(col, glyphFill, seg.style)
	}
	l.putSegmentEnds(first, last, left, right, seg.style)
	l.putText(first+btoi(left)+(inner-len(text))/2, text, seg.style)
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func addSegment(segs []segment, first, end, style int, name string, isPrefix bool) []segment {
	s := segment{first: first, end: end, style: style}
	if isPrefix {
		s.labels = []string{fmt.Sprintf("%s /%d", name, end), fmt.Sprintf("/%d", end), fmt.Sprintf("%d", end)}
	} else {
		s.labels = []string{fmt.Sprintf("%s %d", name, end-first), fmt.Sprintf("%d", end-first)}
	}
	return append(segs, s)
}

func isStructured(net []byte, lay *layout) bool {
	if lay.width != 128 {
		return false
	}
	return net[0]&0xe0 == 0x20 || net[0]&0xfe == 0xfc
}

func makeSegments(prefix int, lay *layout, structured bool) []segment {
	var segs []segment
	switch {
	case structured && prefix < 64:
		segs = addSegment(segs, 0, prefix, styleNetwork, "routing prefix", true)
		segs = addSegment(segs, prefix, 64, styleHost, "subnet ID", false)
		segs = addSegment(segs, 64, 128, styleHost, "interface ID", false)
	case structured && prefix == 64:
		segs = addSegment(segs, 0, 64, styleNetwork, "subnet prefix", true)
		segs = addSegment(segs, 64, 128, styleHost, "interface ID", false)
	default:
		if prefix > 0 {
			segs = addSegment(segs, 0, prefix, styleNetwork, "network", true)
		}
		segs = addSegment(segs, prefix, lay.width, styleHost, "host", false)
	}
	return segs
}

func printBrackets(segs []segment, start int, lay *layout, indent int) {
	end := start + lay.bandBits
	l := newLine(indent)
	for i := range segs {
		first, last := segs[i].first, segs[i].end-1
		if segs[i].end <= start || segs[i].first >= end {
			continue
		}
		if first < start {
			first = start
		}
		if last >= end {
			last = end - 1
		}
		l.putSegment(lay.bitColumn(start, first), lay.bitColumn(start, last),
			segs[i].first >= start, segs[i].end <= end, &segs[i])
	}
	l.print()
}

func printDiagram(info *ipInfo, net []byte, lay *layout) {
	prefix := info.prefix
	indent := 0
	structured := isStructured(net, lay)

	if prefix == lay.width {
		colorFputs("", fmt.Sprintf("%-*s%s\n\n", labelCols, "Address:", info.ip))
	} else {
		colorFputs("", fmt.Sprintf("%-*s%s/%d\n\n", labelCols, "Network:", info.network, prefix))
	}
	if lay.width == 32 && prefix < 32 {
		indent = detailsIndent
	}
	for start := 0; start < lay.width; start += lay.bandBits {
		end := start + lay.bandBits
		if start > 0 && prefix <= start {
			what := "host"
			if structured {
				what = "interface ID"
			}
			colorFputs(colDim, fmt.Sprintf("bits %d-%d: %s, all x", start, lay.width-1, what))
			colorFputs("", "\n")
			return
		}
		printBand(net, prefix, start, lay, indent)
		if prefix < lay.width && prefix <= end {
			printBrackets(makeSegments(prefix, lay, structured), start, lay, indent)
		}
	}
}

func printRow(r *row, prefix int) {
	ip4 := netip.AddrFrom4([4]byte{byte(r.value >> 24), byte(r.value >> 16), byte(r.value >> 8), byte(r.value)})
	l := newLine(0)
	l.putText(0, r.name+":", stylePlain)
	l.putText(labelCols, ip4.String(), stylePlain)
	for k := 0; k < 32; k++ {
		col := detailsIndent + ipv4Layout.bitColumn(0, k)
		st := stylePlain
		if k < prefix {
			st = styleNetwork
		}
		l.putChar(col, byte('0'+((r.value>>(31-uint(k)))&1)), st)
		if k%8 == 7 && k < 31 {
			l.putChar(col+1, '.', stylePlain)
		}
	}
	l.print()
}

func makeRows(info *ipInfo, net []byte, alwaysAddress bool) []row {
	network := uint32(net[0])<<24 | uint32(net[1])<<16 | uint32(net[2])<<8 | uint32(net[3])
	var mask uint32
	if info.prefix != 0 {
		mask = ^uint32(0) << uint(32-info.prefix)
	}
	broadcast := network | ^mask
	var rows []row
	if a, ok := parseIPv4Strict(info.ip); ok && (alwaysAddress || ip4ToU32(a) != network) {
		rows = append(rows, row{"Address", "ADDRESS", ip4ToU32(a)})
	}
	rows = append(rows, row{"Netmask", "NETMASK", mask})
	if info.prefix == 31 {
		rows = append(rows, row{"HostMin", "HOSTMIN", network}, row{"HostMax", "HOSTMAX", broadcast})
	} else {
		rows = append(rows, row{"Network", "NETWORK", network},
			row{"HostMin", "HOSTMIN", network + 1},
			row{"HostMax", "HOSTMAX", broadcast - 1},
			row{"Broadcast", "BROADCAST", broadcast})
	}
	return rows
}

func printDetails(info *ipInfo, net []byte) {
	rows := makeRows(info, net, false)
	colorFputs("", "\n")
	for i := range rows {
		printRow(&rows[i], info.prefix)
	}
}

func formatBits(addr []byte, nbits int) string {
	b := make([]byte, nbits)
	for k := range b {
		b[k] = byte('0' + getBit(addr, k))
	}
	return string(b)
}

func printVar(format outputFormat, state *int, name, value string) {
	if format == formatShell {
		out.WriteString(name + "=" + value + "\n")
	} else {
		jsonPrintf(state, name, true, value)
	}
}

func printVars(info *ipInfo, net []byte, lay *layout, format outputFormat, state *int) {
	if info.prefix == lay.width {
		printVar(format, state, "ADDRESS", formatBits(net, lay.width))
		return
	}
	printVar(format, state, "PREFIX", fmt.Sprint(info.prefix))
	if lay.width == 128 {
		printVar(format, state, "NETWORK", formatBits(net, 128))
		return
	}
	for _, r := range makeRows(info, net, true) {
		b := []byte{byte(r.value >> 24), byte(r.value >> 16), byte(r.value >> 8), byte(r.value)}
		printVar(format, state, r.vr, formatBits(b, 32))
	}
}

func showBinary(info *ipInfo, format outputFormat) {
	lay := &ipv4Layout
	var net [16]byte
	if flags&flagIPv6 != 0 {
		lay = &ipv6Layout
		a, ok := parseIPv6Strict(info.network)
		if !ok {
			errorf("ipcalc: cannot convert the address to binary\n")
			exit(exitFailure)
		}
		net = a.As16()
	} else {
		a, ok := parseIPv4Strict(info.network)
		if !ok {
			errorf("ipcalc: cannot convert the address to binary\n")
			exit(exitFailure)
		}
		b := a.As4()
		copy(net[:], b[:])
	}

	var state int
	switch format {
	case formatValue:
		out.WriteString(formatBits(net[:], lay.width) + "\n")
		return
	case formatShell:
		printVars(info, net[:], lay, format, &state)
		return
	case formatJSON:
		outputStart(&state)
		printVars(info, net[:], lay, format, &state)
		outputStop()
		return
	}

	if flags&flagIPv6 != 0 && outputColumns() >= ipv6WideLayout.bandColumns() {
		lay = &ipv6WideLayout
	}
	printDiagram(info, net[:], lay)
	if flags&flagIPv6 == 0 && info.prefix < lay.width {
		printDetails(info, net[:])
	}
}
