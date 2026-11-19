// SPDX-License-Identifier: GPL-2.0-only
//
// Copyright (c) 1997-2015 Red Hat, Inc.
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
	"crypto/rand"
	"fmt"
	"math/bits"
	"net/netip"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"gitlab.com/ipcalc/go-ipcalc/internal/getopt"
)

// version is set at build time:
// -ldflags "-X gitlab.com/ipcalc/go-ipcalc/internal/ipcalc.version=1.2.0".
// Without it, the module version of the build (go install) or "dev" is used.
var version string

func versionString() string {
	if version != "" {
		return version
	}
	if bi, ok := debug.ReadBuildInfo(); ok && bi.Main.Version != "" && bi.Main.Version != "(devel)" {
		return strings.TrimPrefix(bi.Main.Version, "v")
	}
	return "dev"
}

const (
	appVersion = 1 << iota
	appCheck
	appShowInfo
	appSplit
	appDeaggregate
	appCompare
	appShowBits
)

const (
	flagIPv6 uint = 1 << (iota + 1)
	flagIPv4
	flagShowModernInfo
	flagResolveIP
	flagResolveHost
	flagShowBroadcast
	flagShowNetmask
	flagShowNetwork
	flagShowPrefix
	flagShowMinaddr
	flagShowMaxaddr
	flagShowAddresses
	flagShowAddrspace
	flagGetGeoIP
	flagShowGeoIPBit
	flagShowAllInfo
	flagShowReverse
	flagAssumeClassPrefix
	flagShowCIDR
	flagNoDecorate
	flagShowAddress
	flagJSON
	flagRandom
	flagShowWildcard
)

const flagShowGeoIP = flagShowGeoIPBit | flagGetGeoIP

const (
	optAllInfo = 1 + iota
	optMinaddr
	optMaxaddr
	optAddresses
	optAddrspace
	optUsage
	optReverse
	optClassPrefix
	optNoDecorate
	optHelp
	optSplitHosts
	optEquals
	optSubnetOf
	optOverlaps
	optCIDR
	optWildcard
	optFormat
	optBinary
)

var longOptions = []getopt.LongOption{
	{Name: "check", HasArg: false, Val: 'c'},
	{Name: "equals", HasArg: true, Val: optEquals},
	{Name: "subnet-of", HasArg: true, Val: optSubnetOf},
	{Name: "overlaps", HasArg: true, Val: optOverlaps},
	{Name: "random-private", HasArg: true, Val: 'r'},
	{Name: "split", HasArg: true, Val: 'S'},
	{Name: "split-hosts", HasArg: true, Val: optSplitHosts},
	{Name: "deaggregate", HasArg: true, Val: 'd'},
	{Name: "binary", HasArg: false, Val: optBinary},
	{Name: "info", HasArg: false, Val: 'i'},
	{Name: "all-info", HasArg: false, Val: optAllInfo},
	{Name: "ipv4", HasArg: false, Val: '4'},
	{Name: "ipv6", HasArg: false, Val: '6'},
	{Name: "address", HasArg: false, Val: 'a'},
	{Name: "broadcast", HasArg: false, Val: 'b'},
	{Name: "hostname", HasArg: false, Val: 'h'},
	{Name: "lookup-host", HasArg: true, Val: 'o'},
	{Name: "geoinfo", HasArg: false, Val: 'g'},
	{Name: "reverse-dns", HasArg: false, Val: optReverse},
	{Name: "netmask", HasArg: false, Val: 'm'},
	{Name: "network", HasArg: false, Val: 'n'},
	{Name: "cidr", HasArg: false, Val: optCIDR},
	{Name: "wildcard", HasArg: false, Val: optWildcard},
	{Name: "prefix", HasArg: false, Val: 'p'},
	{Name: "class-prefix", HasArg: false, Val: optClassPrefix},
	{Name: "minaddr", HasArg: false, Val: optMinaddr},
	{Name: "maxaddr", HasArg: false, Val: optMaxaddr},
	{Name: "addresses", HasArg: false, Val: optAddresses},
	{Name: "addrspace", HasArg: false, Val: optAddrspace},
	{Name: "silent", HasArg: false, Val: 's'},
	{Name: "no-decorate", HasArg: false, Val: optNoDecorate},
	{Name: "json", HasArg: false, Val: 'j'},
	{Name: "format", HasArg: true, Val: optFormat},
	{Name: "version", HasArg: false, Val: 'v'},
	{Name: "help", HasArg: false, Val: optHelp},
	{Name: "usage", HasArg: false, Val: optUsage},
}

var fieldSelectingOptions = map[int]uint{
	'a':          flagShowAddress,
	'o':          flagResolveIP,
	'b':          flagShowBroadcast,
	'h':          flagResolveHost,
	'm':          flagShowNetmask,
	'n':          flagShowNetwork,
	'p':          flagShowPrefix,
	'g':          flagShowGeoIP,
	optReverse:   flagShowReverse,
	optCIDR:      flagShowCIDR,
	optWildcard:  flagShowWildcard,
	optMinaddr:   flagShowMinaddr,
	optMaxaddr:   flagShowMaxaddr,
	optAddresses: flagShowAddresses,
	optAddrspace: flagShowAddrspace,
}

func (ap *app) usage(verbose bool) {
	e := func(s string) { fmt.Fprint(ap.errOut, s) }
	if verbose {
		e(`Usage: ipcalc [OPTION...]
  -c, --check                     Validate IP address
      --equals=NET                Succeed if the network equals NET
      --subnet-of=NET             Succeed if the network is within NET
      --overlaps=NET              Succeed if the network shares an address
                                  with NET
  -r, --random-private=PREFIX     Generate a random private IP network using
                                  the supplied prefix or mask.
  -S, --split=[COUNT:]PREFIX      Split the provided network using the
                                  provided prefix/netmask; repeat with
                                  COUNT for subnets of different sizes
      --split-hosts=N1,N2,...     Split the provided network into subnets
                                  that hold N1, N2, ... hosts
  -d, --deaggregate=IP1-IP2       Deaggregate the provided address range
      --binary                    Print the bits of the provided address
                                  or of the network part of a network
  -i, --info                      Print information on the provided IP address
                                  (default)
      --all-info                  Print verbose information on the provided IP
                                  address

Specific info options:
      --reverse-dns               Print network in a the reverse DNS format
  -a, --address                   Display IP address
  -b, --broadcast                 Display calculated broadcast address
  -m, --netmask                   Display netmask for IP
      --wildcard                  Display wildcard (inverse) mask for IP
  -n, --network                   Display network address
      --cidr                      Display network address and prefix in
                                  CIDR notation
  -p, --prefix                    Display network prefix
      --minaddr                   Display the minimum address in the network
      --maxaddr                   Display the maximum address in the network
      --addresses                 Display the maximum number of addresses in
                                  the network
      --addrspace                 Display the address space the network
                                  resides on
  -h, --hostname                  Show hostname determined via DNS
  -o, --lookup-host=STRING        Show IP as determined via DNS
  -g, --geoinfo                   Show Geographic information about the
                                  provided IP

Output format:
      --format=FORMAT             Print the output as human (default for the
                                  summary), shell (NAME=value, default for
                                  specific info options), json, or value

Other options:
  -4, --ipv4                      Explicitly specify the IPv4 address family
  -6, --ipv6                      Explicitly specify the IPv6 address family
      --class-prefix              When specified the default prefix will be determined
                                  by the IPv4 address class
  -s, --silent                    Don't ever display error messages, nor
                                  the result of a comparison
  -v, --version                   Display program version
  -?, --help                      Show this help message
      --usage                     Display brief usage message
`)
		return
	}
	e(`Usage: ipcalc [-46sv?] [-c|--check] [--equals=NET] [--subnet-of=NET]
        [--overlaps=NET] [-r|--random-private=STRING] [-i|--info]
        [--all-info] [-4|--ipv4] [-6|--ipv6] [-a|--address] [-b|--broadcast]
        [-h|--hostname] [-o|--lookup-host=STRING] [-g|--geoinfo]
        [-m|--netmask] [--wildcard] [-n|--network] [--cidr] [-p|--prefix]
        [--minaddr] [--maxaddr] [--addresses] [--addrspace]
        [--format=FORMAT]
        [-s|--silent] [-v|--version] [--reverse-dns] [--class-prefix]
        [--binary] [-?|--help] [--usage]
`)
}

// strtol(s, &end, 0) semantics as used by safe_atoi()
func safeAtoi(s string) (int, bool) {
	t := strings.TrimLeft(s, " \t\n\v\f\r")
	if t == "" || strings.ContainsAny(t, "_") {
		return 0, false
	}
	body := strings.TrimLeft(t, "+-")
	if len(t)-len(body) > 1 {
		return 0, false
	}
	if len(body) > 1 && body[0] == '0' && (body[1] == 'b' || body[1] == 'B' || body[1] == 'o' || body[1] == 'O') {
		return 0, false
	}
	v, err := strconv.ParseInt(t, 0, 64)
	if err != nil || v != int64(int32(v)) {
		return 0, false
	}
	return int(v), true
}

func mask2prefix(a netip.Addr) int {
	i := ip4ToU32(a)
	c := 0
	seenOne := false
	for ; i > 0; i >>= 1 {
		if i&1 != 0 {
			seenOne = true
			c++
		} else if seenOne {
			return -1
		}
	}
	return c
}

func (ap *app) strToPrefix(prefixStr string, fix bool) int {
	var prefix int
	if ap.flags&flagIPv6 == 0 && strings.Contains(prefixStr, ".") {
		a, ok := parseIPv4Strict(prefixStr)
		if !ok {
			return -1
		}
		prefix = mask2prefix(a)
	} else {
		v, ok := safeAtoi(prefixStr)
		if !ok {
			return -1
		}
		prefix = v
	}
	if fix && prefix > 32 && ap.flags&flagIPv6 == 0 {
		ap.flags |= flagIPv6
	}
	if prefix < 0 || (ap.flags&flagIPv6 != 0 && prefix > 128) || (ap.flags&flagIPv6 == 0 && prefix > 32) {
		return -1
	}
	return prefix
}

func (ap *app) generateIPNetwork(prefix int) (string, bool) {
	var ipbuf string
	if ap.flags&flagIPv6 != 0 {
		var b [16]byte
		if _, err := rand.Read(b[1:]); err != nil {
			return "", false
		}
		b[0] = 0xfd
		ipbuf = ntop6(b)
	} else {
		var b [4]byte
		if _, err := rand.Read(b[:]); err != nil {
			return "", false
		}
		ns := time.Now().UnixNano()
		c := ns % 4
		if prefix >= 16 && c < 2 {
			if c == 1 {
				b[0], b[1] = 192, 168
			} else {
				b[0], b[1] = 172, byte(16|((ns>>4)&0x0f))
			}
		} else {
			b[0] = 10
		}
		ipbuf = netip.AddrFrom4(b).String()
	}
	return fmt.Sprintf("%s/%d", ipbuf, prefix), true
}

func (ap *app) parseFormat(name string) int {
	for i, n := range formatNames {
		if n == name {
			return i
		}
	}
	ap.errorf("ipcalc: unknown output format: %s; use human, shell, json or value\n", name)
	return -1
}

func (ap *app) legacyFormat() outputFormat {
	switch {
	case ap.flags&flagJSON != 0:
		return formatJSON
	case ap.flags&flagShowModernInfo != 0:
		return formatHuman
	case ap.flags&flagNoDecorate != 0:
		return formatValue
	}
	return formatShell
}

func (ap *app) formatConflict(option string, f int) bool {
	ap.errorf("ipcalc: conflicting output formats: %s and --format=%s\n", option, formatNames[f])
	return false
}

func (ap *app) selectFormat(name, other string, selected *uint, format *outputFormat) bool {
	f := ap.parseFormat(name)
	if f < 0 {
		return false
	}
	if other != "" {
		if ap.parseFormat(other) < 0 {
			return false
		}
		ap.errorf("ipcalc: conflicting output formats: --format=%s and --format=%s\n", name, other)
		return false
	}
	if ap.flags&flagJSON != 0 && f != int(formatJSON) {
		return ap.formatConflict("-j", f)
	}
	if ap.flags&flagNoDecorate != 0 && f != int(formatValue) {
		return ap.formatConflict("--no-decorate", f)
	}
	if ap.flags&flagShowModernInfo != 0 {
		*selected &^= flagResolveIP
		if *selected != 0 {
			ap.errorf("ipcalc: -i and --all-info cannot be combined with specific field options\n")
			return false
		}
	}
	if f == int(formatJSON) {
		ap.flags |= flagJSON
	} else if f == int(formatValue) {
		ap.flags |= flagNoDecorate
	}
	*format = outputFormat(f)
	return true
}

func (ap *app) checkFormatForApp(format outputFormat, app int) bool {
	unsupported := ""
	if app&(appSplit|appDeaggregate) != 0 && format == formatShell {
		unsupported = "--split, --split-hosts or --deaggregate; use --format=value"
	} else if app&appCompare != 0 && format != formatHuman {
		unsupported = "--equals, --subnet-of or --overlaps"
	}
	if unsupported != "" {
		ap.errorf("ipcalc: --format=%s is not supported with %s\n", formatNames[format], unsupported)
		return false
	}
	if app&^appShowInfo == 0 && format == formatValue && ap.flags&flagShowModernInfo != 0 {
		ap.errorf("ipcalc: --format=value needs specific field options; use --format=shell or --format=json\n")
		return false
	}
	return true
}

func (ap *app) suggestFormat(app int, selected uint) {
	if ap.beSilent || !ap.env.StderrTTY {
		return
	}
	var option, replacement string
	switch {
	case ap.flags&flagJSON != 0:
		option = "-j"
		switch {
		case app&appCompare != 0:
			replacement = "-s to print no result"
		case selected != 0:
			replacement = "--format=json, which prints only the selected fields"
		default:
			replacement = "--format=json"
		}
	case ap.flags&flagNoDecorate != 0:
		option = "--no-decorate"
		switch {
		case app&appCompare != 0:
			replacement = "-s to print no result"
		case selected != 0 || app&(appSplit|appDeaggregate) != 0:
			replacement = "--format=value"
		default:
			replacement = "--format=human"
		}
	default:
		return
	}
	fmt.Fprintf(ap.errOut, "ipcalc: hint: %s is deprecated; use %s\n", option, replacement)
}

func (ap *app) realMain() {
	var (
		randomStr, hostname, ipStr, chptr, prefixStr string
		haveRandom, haveIP, haveHost, haveChptr      bool
		app                                          int
		ncmps                                        int
		selected                                     uint
		format                                       outputFormat
		formatStr, otherFormatStr                    string
		badOption, noDecorate                        bool
		splitStrs, splitHostsStrs                    []string
		cmpStr                                       string
		cmpOp                                        int
		userFlags                                    uint
	)

	g := getopt.New(ap.env.Prog, ap.errOut, ap.env.Args, "S:cr:i46abho:gmnpjsvd:", longOptions)
	for {
		r, ok := g.Next()
		if !ok {
			break
		}
		switch c := r.Opt; c {
		case 'c':
			app |= appCheck
		case optEquals, optSubnetOf, optOverlaps:
			app |= appCompare
			cmpStr = r.Arg
			cmpOp = map[int]int{optEquals: cmpEquals, optSubnetOf: cmpSubnetOf, optOverlaps: cmpOverlaps}[c]
			ncmps++
			ap.exitFailure = 2
		case 'S':
			app |= appSplit
			splitStrs = append(splitStrs, r.Arg)
		case optSplitHosts:
			app |= appSplit
			splitHostsStrs = append(splitHostsStrs, r.Arg)
		case optBinary:
			app |= appShowBits
		case 'd':
			app |= appDeaggregate
			ipStr, haveIP = r.Arg, true
		case 'r':
			app |= appShowInfo
			ap.flags |= flagRandom
			randomStr, haveRandom = r.Arg, true
		case 'i':
			app |= appShowInfo
			ap.flags |= flagShowModernInfo
		case optAllInfo:
			app |= appShowInfo
			ap.flags |= flagShowAllInfo | flagShowModernInfo
		case optClassPrefix:
			ap.flags |= flagAssumeClassPrefix
		case '4':
			ap.flags |= flagIPv4
		case '6':
			ap.flags |= flagIPv6
		case 'o':
			selected |= fieldSelectingOptions['o']
			hostname, haveHost = r.Arg, true
		case optNoDecorate:
			ap.flags |= flagNoDecorate
			noDecorate = true
		case 'j':
			ap.flags |= flagJSON
		case optFormat:
			if formatStr == "" {
				formatStr = r.Arg
			} else if formatStr != r.Arg && otherFormatStr == "" {
				otherFormatStr = r.Arg
			}
		case 's':
			ap.beSilent = true
		case 'v':
			app |= appVersion
		case optUsage:
			if !badOption {
				ap.usage(false)
				ap.exit(0)
			}
		case optHelp:
			if !badOption {
				ap.usage(true)
				ap.exit(0)
			}
		case '?':
			if r.Optopt == '?' && !badOption {
				ap.usage(true)
				ap.exit(0)
			}
			if r.Optopt == optEquals || r.Optopt == optSubnetOf || r.Optopt == optOverlaps {
				ap.exitFailure = 2
			}
			badOption = true
			g.Quiet = true
		default:
			selected |= fieldSelectingOptions[c]
		}
	}

	if selected != 0 {
		app |= appShowInfo
	}
	if badOption {
		ap.usage(true)
		ap.exit(ap.exitFailure)
	}

	if rest := g.Remaining(); len(rest) > 0 {
		if haveIP {
			ap.errorf("ipcalc: superfluous option given\n")
			ap.exit(ap.exitFailure)
		}
		ipStr, haveIP = rest[0], true
		if len(rest) > 1 {
			chptr, haveChptr = rest[1], true
		}
	}

	if formatStr != "" {
		if !ap.selectFormat(formatStr, otherFormatStr, &selected, &format) {
			ap.exit(ap.exitFailure)
		}
	} else if ap.flags&flagJSON != 0 && ap.flags&flagNoDecorate != 0 {
		ap.flags &^= flagNoDecorate
	}

	if selected == 0 || (formatStr == "" && ap.flags&flagJSON != 0) {
		ap.flags |= flagShowModernInfo
	}
	ap.flags |= selected

	if formatStr == "" {
		format = ap.legacyFormat()
	}

	if bits.OnesCount(uint(app)) > 1 || ncmps > 1 {
		ap.errorf("ipcalc: you cannot mix these options\n")
		ap.exit(ap.exitFailure)
	}

	if noDecorate && app&appShowBits != 0 {
		ap.errorf("ipcalc: --no-decorate is not supported with --binary; use --format=value\n")
		ap.exit(ap.exitFailure)
	}

	if formatStr != "" && !ap.checkFormatForApp(format, app) {
		ap.exit(ap.exitFailure)
	}
	if formatStr == "" {
		ap.suggestFormat(app, selected)
	}

	if ap.flags&flagIPv6 != 0 && ap.flags&flagIPv4 != 0 {
		ap.errorf("ipcalc: you cannot specify both IPv4 and IPv6\n")
		ap.exit(ap.exitFailure)
	}

	userFlags = ap.flags

	if ap.flags&flagShowAllInfo != 0 {
		ap.flags |= flagGetGeoIP
	}

	if ap.flags&flagIPv4 == 0 && haveIP && strings.Contains(ipStr, ":") {
		ap.flags |= flagIPv6
	}

	switch app {
	case appVersion:
		fmt.Fprintf(ap.out, "ipcalc %s\n", versionString())
		ap.exit(0)
	case appDeaggregate:
		ap.deaggregate(ipStr)
		ap.exit(0)
	}

	prefix := -1

	if ap.flags&flagRandom != 0 {
		if haveIP {
			ap.errorf("ipcalc: provided superfluous parameter '%s'\n", ipStr)
			ap.exit(ap.exitFailure)
		}
		prefix = ap.strToPrefix(randomStr, true)
		if prefix < 0 {
			fam := "IPv4"
			if ap.flags&flagIPv6 != 0 {
				fam = "IPv6"
			}
			ap.errorf("ipcalc: bad %s prefix: %s\n", fam, randomStr)
			ap.exit(ap.exitFailure)
		}
		s, ok := ap.generateIPNetwork(prefix)
		if !ok {
			ap.errorf("ipcalc: cannot generate network with prefix: %d\n", prefix)
			ap.exit(ap.exitFailure)
		}
		ipStr, haveIP = s, true
		prefix = -1
	}

	if !haveHost && !haveIP {
		ap.errorf("ipcalc: ip address expected\n")
		if !ap.beSilent {
			ap.usage(true)
		}
		ap.exit(ap.exitFailure)
	}

	if haveHost {
		family := 0
		if ap.flags&flagIPv6 != 0 {
			family = 6
		} else if ap.flags&flagIPv4 != 0 {
			family = 4
		}
		s, ok := ap.lookupIP(hostname, family)
		if !ok {
			ap.errorf("ipcalc: could not resolve %s\n", hostname)
			ap.exit(ap.exitFailure)
		}
		ipStr = s
		if ap.flags&flagIPv4 == 0 && strings.Contains(ipStr, ":") {
			ap.flags |= flagIPv6
		}
	}

	havePrefix := false
	if haveChptr {
		if ap.flags&flagIPv6 == 0 {
			prefixStr, havePrefix = chptr, true
		} else {
			ap.errorf("ipcalc: unexpected argument: %s\n", chptr)
			if !ap.beSilent {
				ap.usage(true)
			}
			ap.exit(ap.exitFailure)
		}
	}

	if !havePrefix {
		if i := strings.IndexByte(ipStr, '/'); i >= 0 {
			prefixStr, havePrefix = ipStr[i+1:], true
			ipStr = ipStr[:i]
		}
	}

	if havePrefix {
		prefix = ap.strToPrefix(prefixStr, false)
		if prefix < 0 {
			fam := "IPv4"
			if ap.flags&flagIPv6 != 0 {
				fam = "IPv6"
			}
			ap.errorf("ipcalc: bad %s prefix: %s\n", fam, prefixStr)
			ap.exit(ap.exitFailure)
		}
	}

	var info *ipInfo
	var ok bool
	if ap.flags&flagIPv6 != 0 {
		info, ok = ap.getIPv6Info(ipStr, prefix, ap.flags)
	} else {
		info, ok = ap.getIPv4Info(ipStr, prefix, ap.flags)
	}
	if !ok {
		ap.exit(ap.exitFailure)
	}

	switch app {
	case appSplit:
		if len(splitStrs) > 0 && len(splitHostsStrs) > 0 {
			ap.errorf("ipcalc: --split-hosts cannot be combined with --split\n")
			ap.exit(ap.exitFailure)
		}
		var reqs []splitReq
		for _, h := range splitHostsStrs {
			var ok bool
			if reqs, ok = ap.parseSplitHosts(h, reqs); !ok {
				ap.exit(ap.exitFailure)
			}
		}
		nfill := 0
		for _, sstr := range splitStrs {
			rq, ok := ap.parseSplitReq(sstr)
			if !ok {
				ap.exit(ap.exitFailure)
			}
			if rq.count == 0 {
				if nfill++; nfill > 1 {
					ap.errorf("ipcalc: only one split prefix may be given without a count\n")
					ap.exit(ap.exitFailure)
				}
			}
			reqs = append(reqs, rq)
		}
		ap.showSplitNetworks(reqs, info)
		ap.exit(0)
	case appCheck:
		ap.exit(0)
	case appCompare:
		ap.exit(ap.compareNetworks(cmpOp, info, ap.flags, cmpStr, userFlags))
	}

	_, noColor := ap.env.Getenv("NO_COLOR")
	if ap.env.StdoutTTY && !noColor {
		ap.colors = true
	}

	if app == appShowBits {
		ap.showBinary(info, format)
		ap.exit(0)
	}
	ap.showInfoFields(info, format, haveRandom)
	ap.exit(0)
}
