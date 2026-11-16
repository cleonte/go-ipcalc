package ipcalc

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const (
	colBlue  = "\x1b[34m"
	colMag   = "\x1b[35m"
	colReset = "\033[0m"
)

type outputFormat int

const (
	formatHuman outputFormat = iota
	formatShell
	formatJSON
	formatValue
)

var formatNames = []string{"human", "shell", "json", "value"}

// stdout is fully buffered and flushed on exit, like C stdio on a pipe, so
// that diagnostics on stderr come first in merged output.
func newOut(w io.Writer) *bufio.Writer { return bufio.NewWriterSize(w, 1<<20) }

const (
	jsonFirst = iota
	jsonNext
	jsonArrayFirst
	jsonArrayNext
)

func (ap *app) errorf(format string, a ...any) {
	if ap.beSilent {
		return
	}
	fmt.Fprintf(ap.errOut, format, a...)
}

func (ap *app) colorPrintf(color, title, value string) {
	ap.out.WriteString(title)
	if ap.colors {
		ap.out.WriteString(color)
	}
	ap.out.WriteString(value)
	ap.out.WriteString("\n")
	if ap.colors {
		ap.out.WriteString(colReset)
	}
}

func (ap *app) jsonString(s string) {
	ap.out.WriteByte('"')
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c == '"' || c == '\\':
			ap.out.WriteByte('\\')
			ap.out.WriteByte(c)
		case c < 0x20:
			fmt.Fprintf(ap.out, "\\u%04x", c)
		default:
			ap.out.WriteByte(c)
		}
	}
	ap.out.WriteByte('"')
}

func (ap *app) jsonPrintf(state *int, title string, hasTitle bool, value string) {
	if *state == jsonArrayNext {
		ap.out.WriteString(",\n  ")
	} else if *state == jsonNext {
		ap.out.WriteString(",\n")
	}
	ap.out.WriteString("  ")
	if hasTitle {
		fmt.Fprintf(ap.out, "\"%s\":", title)
	}
	ap.jsonString(value)
	if *state == jsonFirst {
		*state = jsonNext
	} else if *state == jsonArrayFirst {
		*state = jsonArrayNext
	}
}

func (ap *app) prettyPrintf(title, value string) {
	if ap.flags&flagNoDecorate != 0 {
		ap.out.WriteString(title + value + "\n")
	} else {
		ap.colorPrintf(colBlue, title, value)
	}
}

func (ap *app) prettyDistPrintf(title, value string) {
	if ap.flags&flagNoDecorate != 0 {
		ap.out.WriteString(title + value + "\n")
	} else {
		ap.colorPrintf(colMag, title, value)
	}
}

func (ap *app) outputStart(state *int) {
	if ap.flags&flagJSON != 0 {
		ap.out.WriteString("{\n")
	}
	*state = jsonFirst
}

func (ap *app) outputSeparate() {
	if ap.flags&flagJSON == 0 {
		ap.out.WriteString("\n")
	}
}

func (ap *app) outputStop() {
	if ap.flags&flagJSON != 0 {
		ap.out.WriteString("\n}\n")
	}
}

func (ap *app) arrayStart(state *int, head, jsonHead string) {
	if ap.flags&flagJSON != 0 {
		if *state == jsonNext {
			ap.out.WriteString(",\n")
		}
		fmt.Fprintf(ap.out, "  \"%s\":[\n  ", jsonHead)
	} else if ap.flags&flagNoDecorate == 0 {
		fmt.Fprintf(ap.out, "[%s]\n", head)
	}
	*state = jsonArrayFirst
}

func (ap *app) arrayStop(state *int) {
	if ap.flags&flagJSON != 0 {
		ap.out.WriteString("]")
		*state = jsonNext
	}
}

const (
	attrSummaryAlways = 1 << iota
	attrSummaryWithAllInfo
	attrSummaryJSONOnly
	attrHideForNetworkAddress
	attrHideForSingleAddress
	attrHighlight
	attrVarQuoteAlways
	attrValueQuoteIfSpace
	attrRequiresSingleAddress
)

const (
	sectionNetwork = iota
	sectionRange
	sectionGeo
)

type outputField struct {
	ap          *app
	title, name string
	option      uint
	optionName  string
	attrs       int
	section     int
	value       string
	textValue   string
	isList      bool
	values      []string
}

func (f *outputField) hasValue() bool {
	return f.value != "" || len(f.values) > 0
}

func (f *outputField) inSummary(format outputFormat, single, inputIsNetwork bool) bool {
	if !f.hasValue() {
		return false
	}
	if f.attrs&attrHideForNetworkAddress != 0 && inputIsNetwork {
		return false
	}
	if format == formatHuman &&
		((f.attrs&attrHideForSingleAddress != 0 && single) || f.attrs&attrSummaryJSONOnly != 0) {
		return false
	}
	if f.ap.flags&f.option != 0 {
		return true
	}
	if f.attrs&(attrSummaryAlways|attrSummaryJSONOnly) != 0 {
		return true
	}
	return f.attrs&attrSummaryWithAllInfo != 0 && f.ap.flags&flagShowAllInfo != 0
}

func (f *outputField) showJSON(state *int) {
	if !f.isList {
		f.ap.jsonPrintf(state, f.name, true, f.value)
		return
	}
	f.ap.arrayStart(state, f.title, f.name)
	for _, v := range f.values {
		f.ap.jsonPrintf(state, "", false, v)
	}
	f.ap.arrayStop(state)
}

func (f *outputField) showText() {
	v := f.textValue
	if v == "" {
		v = f.value
	}
	if f.isList {
		for _, s := range f.values {
			f.ap.prettyPrintf(f.title, s)
		}
	} else if f.attrs&attrHighlight != 0 {
		f.ap.prettyDistPrintf(f.title, v)
	} else {
		f.ap.prettyPrintf(f.title, v)
	}
}

func shellSafe(v string) bool {
	for i := 0; i < len(v); i++ {
		c := v[i]
		if c >= 0x80 || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			continue
		}
		if !strings.ContainsRune("._:/-,+@%", rune(c)) {
			return false
		}
	}
	return true
}

func shellEscaped(v string) string {
	var sb strings.Builder
	for i := 0; i < len(v); i++ {
		if strings.IndexByte("\"\\$`", v[i]) >= 0 {
			sb.WriteByte('\\')
		}
		sb.WriteByte(v[i])
	}
	return sb.String()
}

func shellValue(v string, quote bool) string {
	if !quote && shellSafe(v) {
		return v
	}
	return "\"" + shellEscaped(v) + "\""
}

func (f *outputField) showVar(bare bool) {
	if f.isList {
		switch {
		case bare:
			for _, v := range f.values {
				f.ap.out.WriteString(v + "\n")
			}
		case len(f.values) == 1:
			f.ap.out.WriteString(f.name + "=" + shellValue(f.values[0], false) + "\n")
		default:
			esc := make([]string, len(f.values))
			for i, v := range f.values {
				esc[i] = shellEscaped(v)
			}
			f.ap.out.WriteString(f.name + "=\"" + strings.Join(esc, " ") + "\"\n")
		}
		return
	}
	if bare {
		if f.attrs&attrVarQuoteAlways != 0 ||
			(f.attrs&attrValueQuoteIfSpace != 0 && strings.Contains(f.value, " ")) {
			f.ap.out.WriteString("\"" + f.value + "\"\n")
		} else {
			f.ap.out.WriteString(f.value + "\n")
		}
		return
	}
	f.ap.out.WriteString(f.name + "=" + shellValue(f.value, f.attrs&attrVarQuoteAlways != 0) + "\n")
}

func (f *outputField) selected(format outputFormat, single, inputIsNetwork bool) bool {
	if f.ap.flags&flagShowModernInfo != 0 {
		return f.inSummary(format, single, inputIsNetwork)
	}
	return f.hasValue() && f.ap.flags&f.option != 0
}

func (f *outputField) show(state *int, format outputFormat) {
	switch format {
	case formatHuman:
		f.showText()
	case formatShell:
		f.showVar(false)
	case formatJSON:
		f.showJSON(state)
	case formatValue:
		f.showVar(true)
	}
}

func (ap *app) showInfoFields(info *ipInfo, format outputFormat, randomized bool) {
	humanSummary := ap.flags&flagShowModernInfo != 0 && format == formatHuman
	width := 32
	if ap.flags&flagIPv6 != 0 {
		width = 128
	}
	single := info.prefix == width
	inputIsNetwork := !single && (randomized || info.network == info.ip)

	prefix := fmt.Sprint(info.prefix)
	cidr := fmt.Sprintf("%s/%d", info.network, info.prefix)
	netmask := fmt.Sprintf("%s = %d", info.netmask, info.prefix)
	fullNetwork := ""
	if info.expandedNetwork != "" {
		fullNetwork = fmt.Sprintf("%s/%d", info.expandedNetwork, info.prefix)
	}
	radius := ""
	if info.geo.accuracyRadius != "" {
		radius = info.geo.accuracyRadius + " km"
	}
	hosts := ""
	if ap.flags&flagIPv6 != 0 && info.prefix < 112 {
		hosts = fmt.Sprintf("2^(%d) = %s", 128-info.prefix, info.hosts)
	}

	const (
		sumAll = attrSummaryAlways
		hideNA = attrHideForNetworkAddress
		hideSA = attrHideForSingleAddress
	)
	fields := []outputField{
		{title: "Full Address:\t", name: "FULLADDRESS", attrs: sumAll | hideNA, value: info.expandedIP},
		{title: "Address:\t", name: "ADDRESS", option: flagShowAddress | flagResolveIP, attrs: sumAll | hideNA, value: info.ip},
		{title: "Hostname:\t", name: "HOSTNAME", option: flagResolveHost, optionName: "-h", attrs: sumAll | attrRequiresSingleAddress, value: info.hostname},
		{title: "Full Network:\t", name: "FULLNETWORK", attrs: sumAll | hideSA, value: info.expandedNetwork, textValue: fullNetwork},
		{title: "Network:\t", name: "NETWORK", option: flagShowNetwork, attrs: sumAll | hideSA, value: info.network, textValue: cidr},
		{title: "Netmask:\t", name: "NETMASK", option: flagShowNetmask, attrs: sumAll | hideSA, value: info.netmask, textValue: netmask},
		{title: "Wildcard:\t", name: "WILDCARD", option: flagShowWildcard, attrs: attrSummaryWithAllInfo | hideSA, value: info.wildcard},
		{title: "Prefix:\t", name: "PREFIX", option: flagShowPrefix, attrs: attrSummaryJSONOnly, value: prefix},
		{title: "CIDR:\t", name: "CIDR", option: flagShowCIDR, attrs: attrSummaryJSONOnly, value: cidr},
		{title: "Broadcast:\t", name: "BROADCAST", option: flagShowBroadcast, attrs: sumAll | hideSA, value: info.broadcast},
		{title: "Reverse DNS:\t", name: "REVERSEDNS", option: flagShowReverse, attrs: attrSummaryWithAllInfo, isList: true, values: info.reverseDNS},
		{title: "Address space:\t", name: "ADDRSPACE", option: flagShowAddrspace, attrs: sumAll | attrHighlight | attrValueQuoteIfSpace, section: sectionRange, value: info.typ},
		{title: "Address class:\t", name: "ADDRCLASS", attrs: attrSummaryWithAllInfo | attrHighlight, section: sectionRange, value: info.class},
		{title: "Interface ID:\t", name: "INTERFACEID", attrs: attrSummaryWithAllInfo | attrHighlight, section: sectionRange, value: info.iid},
		{title: "EUI-64:\t\t", name: "EUI64", attrs: attrSummaryWithAllInfo | attrHighlight, section: sectionRange, value: info.eui64},
		{title: "MAC address:\t", name: "MACADDR", attrs: attrSummaryWithAllInfo | attrHighlight, section: sectionRange, value: info.mac},
		{title: "MAC scope:\t", name: "MACSCOPE", attrs: attrSummaryWithAllInfo | attrHighlight, section: sectionRange, value: info.macScope},
		{title: "MAC type:\t", name: "MACTYPE", attrs: attrSummaryWithAllInfo | attrHighlight, section: sectionRange, value: info.macType},
		{title: "HostMin:\t", name: "MINADDR", option: flagShowMinaddr, attrs: sumAll | hideSA, section: sectionRange, value: info.hostmin},
		{title: "HostMax:\t", name: "MAXADDR", option: flagShowMaxaddr, attrs: sumAll | hideSA, section: sectionRange, value: info.hostmax},
		{title: "Hosts/Net:\t", name: "ADDRESSES", option: flagShowAddresses, attrs: sumAll | hideSA | attrValueQuoteIfSpace, section: sectionRange, value: info.hosts, textValue: hosts},
		{title: "Continent:\t", name: "CONTINENT", option: flagShowGeoIP, attrs: sumAll | attrHighlight | attrValueQuoteIfSpace, section: sectionGeo, value: info.geo.continent},
		{title: "Country code:\t", name: "COUNTRYCODE", option: flagShowGeoIP, attrs: sumAll | attrHighlight, section: sectionGeo, value: info.geo.ccode},
		{title: "Country:\t", name: "COUNTRY", option: flagShowGeoIP, attrs: sumAll | attrHighlight | attrValueQuoteIfSpace, section: sectionGeo, value: info.geo.country},
		{title: "City:\t\t", name: "CITY", option: flagShowGeoIP, attrs: sumAll | attrHighlight | attrValueQuoteIfSpace, section: sectionGeo, value: info.geo.city},
		{title: "Postal code:\t", name: "POSTALCODE", option: flagShowGeoIP, attrs: sumAll | attrHighlight | attrValueQuoteIfSpace, section: sectionGeo, value: info.geo.postalCode},
		{title: "Coordinates:\t", name: "COORDINATES", option: flagShowGeoIP, attrs: sumAll | attrHighlight | attrVarQuoteAlways, section: sectionGeo, value: info.geo.coord},
		{title: "Accuracy:\t", name: "ACCURACYRADIUS", option: flagShowGeoIP, attrs: sumAll | attrHighlight, section: sectionGeo, value: info.geo.accuracyRadius, textValue: radius},
		{title: "Time zone:\t", name: "TIMEZONE", option: flagShowGeoIP, attrs: sumAll | attrHighlight, section: sectionGeo, value: info.geo.timeZone},
		{title: "Routed by AS:\t", name: "ASN", option: flagShowGeoIP, attrs: sumAll | attrHighlight, section: sectionGeo, value: info.geo.asn},
		{title: "Operated by:\t", name: "ASORG", option: flagShowGeoIP, attrs: sumAll | attrHighlight | attrValueQuoteIfSpace, section: sectionGeo, value: info.geo.asOrg},
		{title: "AS network:\t", name: "ASNETWORK", option: flagShowGeoIP, attrs: sumAll | attrHighlight, section: sectionGeo, value: info.geo.asNetwork},
		{title: "Geo network:\t", name: "GEONETWORK", option: flagShowGeoIP, attrs: sumAll | attrHighlight, section: sectionGeo, value: info.geo.network},
	}

	for i := range fields {
		fields[i].ap = ap
	}
	for i := range fields {
		f := &fields[i]
		if f.attrs&attrRequiresSingleAddress != 0 && ap.flags&f.option != 0 && !single {
			ap.errorf("ipcalc: %s requires a single address, not a network: %s/%d\n",
				f.optionName, info.ip, info.prefix)
			ap.exit(ap.exitFailure)
		}
	}

	var state int
	section := sectionNetwork
	ap.outputStart(&state)
	for i := range fields {
		f := &fields[i]
		if !f.selected(format, single, inputIsNetwork) {
			continue
		}
		if humanSummary && f.section != section {
			if f.section == sectionGeo || !single {
				ap.outputSeparate()
			}
			section = f.section
		}
		f.show(&state, format)
	}
	ap.outputStop()
}
