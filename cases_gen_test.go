// Code generated from the ipcalc C project test list; DO NOT EDIT.

package main

var generatedCases = []testCase{
	{name: "Split-Info",
		steps: []step{
			{kind: "failure", args: []string{"-abmnp", "-S", "26", "10.100.1.0/24"}},
		}},
	{name: "SplitHosts-Info",
		steps: []step{
			{kind: "failure", args: []string{"-n", "--split-hosts=10", "10.100.1.0/24"}},
		}},
	{name: "Check-Info",
		steps: []step{
			{kind: "failure", args: []string{"-c", "-abmnp", "10.100.1.1"}},
		}},
	{name: "Check-Split",
		steps: []step{
			{kind: "failure", args: []string{"-c", "-S", "26", "10.100.1.0/24"}},
		}},
	{name: "Check-Host",
		steps: []step{
			{kind: "failure", args: []string{"-c", "-h", "127.0.0.1"}},
		}},
	{name: "Random-Info",
		steps: []step{
			{kind: "failure", args: []string{"-r", "24", "-i", "127.0.0.1"}},
		}},
	{name: "HostnameIPv6Localhost",
		steps: []step{
			{kind: "outfile", args: []string{"-6", "-o", "localhost"}, file: "hostname-localhost-ipv6"},
		}},
	{name: "HostnameIPv4Localhost",
		steps: []step{
			{kind: "outfile", args: []string{"-4", "-o", "localhost"}, file: "hostname-localhost-ipv4"},
		}},
	{name: "HostnameIPv4LocalhostJson",
		steps: []step{
			{kind: "outfile", args: []string{"-j", "-4", "-o", "localhost"}, file: "hostname-localhost-ipv4-json"},
		}},
	{name: "IPIPv6Localhost",
		steps: []step{
			{kind: "outfile", args: []string{"-h", "::1"}, file: "ip-localhost-ipv6"},
		}},
	{name: "IPIPv4Localhost",
		steps: []step{
			{kind: "outfile", args: []string{"-h", "127.0.0.1"}, file: "ip-localhost-ipv4"},
		}},
	{name: "IPIPv4LocalhostJson",
		steps: []step{
			{kind: "outfile", args: []string{"-j", "-h", "127.0.0.1"}, file: "ip-localhost-ipv4-json"},
		}},
	{name: "AssignClassPrefix12",
		steps: []step{
			{kind: "outfile", args: []string{"-abmnp", "12.15.1.5", "--class-prefix"}, file: "12.15.1.5"},
		}},
	{name: "AssignClassPrefix129",
		steps: []step{
			{kind: "outfile", args: []string{"-abmnp", "129.15.31.5", "--class-prefix"}, file: "129.15.31.5"},
		}},
	{name: "AssignClassPrefix193",
		steps: []step{
			{kind: "outfile", args: []string{"-abmnp", "193.92.31.0", "--class-prefix"}, file: "193.92.31.0"},
		}},
	{name: "TestPrefix31",
		steps: []step{
			{kind: "outfile", args: []string{"-abmnp", "192.168.1.5/31"}, file: "192.168.1.5-31"},
		}},
	{name: "TestPrefix24",
		steps: []step{
			{kind: "outfile", args: []string{"-abmnp", "10.10.10.5/24"}, file: "192.168.1.5-24"},
		}},
	{name: "TestPrefix30",
		steps: []step{
			{kind: "outfile", args: []string{"-abmnp", "10.100.4.1/30"}, file: "192.168.1.5-30"},
		}},
	{name: "TestPrefix16",
		steps: []step{
			{kind: "outfile", args: []string{"-abmnp", "10.100.4.1/16"}, file: "192.168.1.5-16"},
		}},
	{name: "TestPrefix8",
		steps: []step{
			{kind: "outfile", args: []string{"-abmnp", "10.10.10.10/8"}, file: "192.168.1.5-8"},
		}},
	{name: "SplitPrefix18",
		steps: []step{
			{kind: "outfile", args: []string{"-S", "18", "10.10.10.10/16"}, file: "split-10.10.10.0-16-18"},
		}},
	{name: "SplitPrefix24",
		steps: []step{
			{kind: "outfile", args: []string{"-S", "24", "10.10.10.0/16"}, file: "split-10.10.10.0-16-24"},
		}},
	{name: "SplitPrefix26",
		steps: []step{
			{kind: "outfile", args: []string{"-S", "26", "192.168.5.45/24"}, file: "split-192.168.5.45-24-26"},
		}},
	{name: "SplitPrefix29",
		steps: []step{
			{kind: "outfile", args: []string{"-S", "29", "192.168.5.0/24"}, file: "split-192.168.5.0-24-29"},
		}},
	{name: "SplitPrefix31",
		steps: []step{
			{kind: "outfile", args: []string{"-S", "31", "192.168.5.0/24"}, file: "split-192.168.5.0-24-31"},
		}},
	{name: "SplitPrefix32",
		steps: []step{
			{kind: "outfile", args: []string{"-S", "32", "192.168.5.0/24"}, file: "split-192.168.5.0-24-32"},
		}},
	{name: "SplitPrefix64",
		steps: []step{
			{kind: "outfile", args: []string{"-S", "64", "2a03:2880:20:4f06:face::/56"}, file: "split-2a03:2880:20:4f06:face::-56-64"},
		}},
	{name: "SplitPrefix128",
		steps: []step{
			{kind: "outfile", args: []string{"-S", "128", "fcfa:b4ca:f1d8:125b:dc00::/127"}, file: "split-fcfa:b4ca:f1d8:125b:dc00::-127-128"},
		}},
	{name: "SplitPrefix120",
		steps: []step{
			{kind: "outfile", args: []string{"-S", "120", "fcfa:b4ca:f1d8:125b:dc00::/112"}, file: "split-fcfa:b4ca:f1d8:125b:dc00::-112-120"},
		}},
	{name: "SplitPrefix0IPv4",
		steps: []step{
			{kind: "outfile", args: []string{"-S", "0", "0.0.0.0/0"}, file: "split-0.0.0.0-0-0"},
		}},
	{name: "SplitPrefix3Of1",
		steps: []step{
			{kind: "outfile", args: []string{"-S", "3", "10.0.0.0/1"}, file: "split-10.0.0.0-1-3"},
		}},
	{name: "SplitPrefix2Of0IPv6",
		steps: []step{
			{kind: "outfile", args: []string{"-S", "2", "::/0"}, file: "split-::-0-2"},
		}},
	{name: "SplitPrefix124EndOfSpace",
		steps: []step{
			{kind: "outfile", args: []string{"-S", "124", "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/120"}, file: "split-ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff-120-124"},
		}},
	{name: "SplitVLSMPrefix58And64",
		steps: []step{
			{kind: "outfile", args: []string{"-S", "1:58", "-S", "64", "2001:db8:1c88:6000::/56"}, file: "split-2001:db8:1c88:6000::-56-1:58-64"},
		}},
	{name: "NoDecorateSplitVLSMPrefix58And64",
		steps: []step{
			{kind: "outfile", args: []string{"--no-decorate", "-S", "1:58", "-S", "64", "2001:db8:1c88:6000::/56"}, file: "nsplit-2001:db8:1c88:6000::-56-1:58-64"},
		}},
	{name: "JsonSplitVLSMPrefix58And64",
		steps: []step{
			{kind: "outfile", args: []string{"-j", "-S", "1:58", "-S", "64", "2001:db8:1c88:6000::/56"}, file: "json-split-2001:db8:1c88:6000::-56-1:58-64"},
		}},
	{name: "SplitVLSMUnused",
		steps: []step{
			{kind: "outfile", args: []string{"-S", "1:28", "-S", "2:27", "192.168.0.0/24"}, file: "split-192.168.0.0-24-1:28-2:27"},
		}},
	{name: "NoDecorateSplitVLSMUnused",
		steps: []step{
			{kind: "outfile", args: []string{"--no-decorate", "-S", "1:28", "-S", "2:27", "192.168.0.0/24"}, file: "nsplit-192.168.0.0-24-1:28-2:27"},
		}},
	{name: "JsonSplitVLSMUnused",
		steps: []step{
			{kind: "outfile", args: []string{"-j", "-S", "1:28", "-S", "2:27", "192.168.0.0/24"}, file: "json-split-192.168.0.0-24-1:28-2:27"},
		}},
	{name: "SplitVLSMFillAligned",
		steps: []step{
			{kind: "outfile", args: []string{"-S", "3:64", "-S", "60", "2001:db8::/56"}, file: "split-2001:db8::-56-3:64-60"},
		}},
	{name: "SplitVLSMSamePrefix",
		steps: []step{
			{kind: "outfile", args: []string{"-S", "2:26", "-S", "26", "192.168.1.0/24"}, file: "split-192.168.1.0-24-2:26-26"},
		}},
	{name: "SplitVLSMNetmask",
		steps: []step{
			{kind: "outfile", args: []string{"-S", "2:255.255.255.192", "192.168.1.0/24"}, file: "split-192.168.1.0-24-2:255.255.255.192"},
		}},
	{name: "SplitVLSMWholeSpace",
		steps: []step{
			{kind: "outfile", args: []string{"-S", "2", "-S", "1:1", "::/0"}, file: "split-::-0-2-1:1"},
		}},
	{name: "SplitVLSMTooSmall",
		steps: []step{
			{kind: "failure", args: []string{"-S", "1:25", "-S", "1:25", "-S", "1:26", "192.168.0.0/24"}},
			{kind: "output", args: []string{"-S", "1:25", "-S", "1:25", "-S", "1:26", "192.168.0.0/24"}, want: "ipcalc: 192.168.0.0/24 is too small for the requested subnets"},
		}},
	{name: "SplitVLSMNoSpaceForFill",
		steps: []step{
			{kind: "failure", args: []string{"-S", "2:25", "-S", "26", "192.168.0.0/24"}},
			{kind: "output", args: []string{"-S", "2:25", "-S", "26", "192.168.0.0/24"}, want: "ipcalc: 192.168.0.0/24 is too small for the requested subnets"},
		}},
	{name: "SplitVLSMTwoFills",
		steps: []step{
			{kind: "failure", args: []string{"-S", "26", "-S", "27", "192.168.0.0/24"}},
			{kind: "output", args: []string{"-S", "26", "-S", "27", "192.168.0.0/24"}, want: "ipcalc: only one split prefix may be given without a count"},
		}},
	{name: "SplitVLSMZeroCount",
		steps: []step{
			{kind: "failure", args: []string{"-S", "0:26", "192.168.0.0/24"}},
			{kind: "output", args: []string{"-S", "0:26", "192.168.0.0/24"}, want: "ipcalc: bad split count: 0:26"},
		}},
	{name: "SplitVLSMBadCount",
		steps: []step{
			{kind: "failure", args: []string{"-S", "x:26", "192.168.0.0/24"}},
			{kind: "output", args: []string{"-S", "x:26", "192.168.0.0/24"}, want: "ipcalc: bad split count: x:26"},
		}},
	{name: "SplitVLSMCountOverflow",
		steps: []step{
			{kind: "failure", args: []string{"-S", "99999999999999999999:26", "192.168.0.0/24"}},
			{kind: "output", args: []string{"-S", "99999999999999999999:26", "192.168.0.0/24"}, want: "ipcalc: bad split count: 99999999999999999999:26"},
		}},
	{name: "SplitVLSMBadPrefix",
		steps: []step{
			{kind: "failure", args: []string{"-S", "1:abc", "192.168.0.0/24"}},
			{kind: "output", args: []string{"-S", "1:abc", "192.168.0.0/24"}, want: "ipcalc: bad IPv4 prefix: 1:abc"},
		}},
	{name: "SplitHosts",
		steps: []step{
			{kind: "outfile", args: []string{"--split-hosts=10,20,30", "192.168.0.0/24"}, file: "split-hosts-192.168.0.0-24-10,20,30"},
		}},
	{name: "NoDecorateSplitHosts",
		steps: []step{
			{kind: "outfile", args: []string{"--no-decorate", "--split-hosts=10,20,30", "192.168.0.0/24"}, file: "nsplit-hosts-192.168.0.0-24-10,20,30"},
		}},
	{name: "JsonSplitHosts",
		steps: []step{
			{kind: "outfile", args: []string{"-j", "--split-hosts=10,20,30", "192.168.0.0/24"}, file: "json-split-hosts-192.168.0.0-24-10,20,30"},
		}},
	{name: "SplitHostsSmallest",
		steps: []step{
			{kind: "outfile", args: []string{"--split-hosts=1,2,62,63", "192.168.0.0/24"}, file: "split-hosts-192.168.0.0-24-1,2,62,63"},
		}},
	{name: "SplitHostsRepeated",
		steps: []step{
			{kind: "outfile", args: []string{"--split-hosts=1,2", "--split-hosts=62,63", "192.168.0.0/24"}, file: "split-hosts-192.168.0.0-24-1,2,62,63"},
		}},
	{name: "SplitHostsExactFit",
		steps: []step{
			{kind: "outfile", args: []string{"--no-decorate", "--split-hosts=126,62,62", "192.168.0.0/24"}, file: "nsplit-hosts-192.168.0.0-24-126,62,62"},
		}},
	{name: "SplitHostsIPv6",
		steps: []step{
			{kind: "outfile", args: []string{"--split-hosts=1,2,3,256", "2001:db8::/112"}, file: "split-hosts-2001:db8::-112-1,2,3,256"},
		}},
	{name: "JsonSplitHostsWholeIPv4",
		steps: []step{
			{kind: "outfile", args: []string{"-j", "--split-hosts=4294967294", "0.0.0.0/0"}, file: "json-split-hosts-0.0.0.0-0-4294967294"},
		}},
	{name: "SplitHostsTooSmall",
		steps: []step{
			{kind: "failure", args: []string{"--split-hosts=200,100", "192.168.0.0/24"}},
			{kind: "output", args: []string{"--split-hosts=200,100", "192.168.0.0/24"}, want: "ipcalc: 192.168.0.0/24 is too small for the requested subnets"},
		}},
	{name: "SplitHostsLargerThanNetwork",
		steps: []step{
			{kind: "failure", args: []string{"--split-hosts=300", "192.168.0.0/24"}},
			{kind: "output", args: []string{"--split-hosts=300", "192.168.0.0/24"}, want: "ipcalc: 192.168.0.0/24 is too small for the requested subnets"},
		}},
	{name: "SplitHostsTooManyHosts",
		steps: []step{
			{kind: "failure", args: []string{"--split-hosts=4294967295", "0.0.0.0/0"}},
			{kind: "output", args: []string{"--split-hosts=4294967295", "0.0.0.0/0"}, want: "ipcalc: too many hosts requested: 4294967295"},
		}},
	{name: "SplitHostsWithSplit",
		steps: []step{
			{kind: "failure", args: []string{"--split-hosts=10", "-S", "26", "192.168.0.0/24"}},
			{kind: "output", args: []string{"--split-hosts=10", "-S", "26", "192.168.0.0/24"}, want: "ipcalc: --split-hosts cannot be combined with --split"},
		}},
	{name: "SplitHostsZero",
		steps: []step{
			{kind: "failure", args: []string{"--split-hosts=0", "192.168.0.0/24"}},
			{kind: "output", args: []string{"--split-hosts=0", "192.168.0.0/24"}, want: "ipcalc: bad host count: 0"},
		}},
	{name: "SplitHostsEmptyItem",
		steps: []step{
			{kind: "failure", args: []string{"--split-hosts=1,,2", "192.168.0.0/24"}},
			{kind: "output", args: []string{"--split-hosts=1,,2", "192.168.0.0/24"}, want: "ipcalc: bad host count: 1,,2"},
		}},
	{name: "SplitHostsOverflow",
		steps: []step{
			{kind: "failure", args: []string{"--split-hosts=18446744073709551616", "2001:db8::/64"}},
			{kind: "output", args: []string{"--split-hosts=18446744073709551616", "2001:db8::/64"}, want: "ipcalc: bad host count: 18446744073709551616"},
		}},
	{name: "NoDecorateSplitPrefix18",
		steps: []step{
			{kind: "outfile", args: []string{"--no-decorate", "-S", "18", "10.10.10.10/16"}, file: "nsplit-10.10.10.0-16-18"},
		}},
	{name: "NoDecorateSplitPrefix24",
		steps: []step{
			{kind: "outfile", args: []string{"--no-decorate", "-S", "24", "10.10.10.0/16"}, file: "nsplit-10.10.10.0-16-24"},
		}},
	{name: "NoDecorateSplitPrefix26",
		steps: []step{
			{kind: "outfile", args: []string{"--no-decorate", "-S", "26", "192.168.5.45/24"}, file: "nsplit-192.168.5.45-24-26"},
		}},
	{name: "NoDecorateSplitPrefix29",
		steps: []step{
			{kind: "outfile", args: []string{"--no-decorate", "-S", "29", "192.168.5.0/24"}, file: "nsplit-192.168.5.0-24-29"},
		}},
	{name: "NoDecorateSplitPrefix31",
		steps: []step{
			{kind: "outfile", args: []string{"--no-decorate", "-S", "31", "192.168.5.0/24"}, file: "nsplit-192.168.5.0-24-31"},
		}},
	{name: "NoDecorateSplitPrefix32",
		steps: []step{
			{kind: "outfile", args: []string{"--no-decorate", "-S", "32", "192.168.5.0/24"}, file: "nsplit-192.168.5.0-24-32"},
		}},
	{name: "NoDecorateSplitPrefix64",
		steps: []step{
			{kind: "outfile", args: []string{"--no-decorate", "-S", "64", "2a03:2880:20:4f06:face::/56"}, file: "nsplit-2a03:2880:20:4f06:face::-56-64"},
		}},
	{name: "NoDecorateSplitPrefix128",
		steps: []step{
			{kind: "outfile", args: []string{"--no-decorate", "-S", "128", "fcfa:b4ca:f1d8:125b:dc00::/127"}, file: "nsplit-fcfa:b4ca:f1d8:125b:dc00::-127-128"},
		}},
	{name: "NoDecorateSplitPrefix120",
		steps: []step{
			{kind: "outfile", args: []string{"--no-decorate", "-S", "120", "fcfa:b4ca:f1d8:125b:dc00::/112"}, file: "nsplit-fcfa:b4ca:f1d8:125b:dc00::-112-120"},
		}},
	{name: "AllocateAddressSpacePrefix24",
		steps: []step{
			{kind: "outfile", args: []string{"--addrspace", "-abmnp", "193.92.150.3/24"}, file: "193.92.150.3-24"},
		}},
	{name: "AllocateAddressSpacePrefix64",
		steps: []step{
			{kind: "outfile", args: []string{"--addrspace", "-abmnp", "fd95:6be5:0ae0:84a5::/64"}, file: "fd95:6be5:0ae0:84a5::-64"},
		}},
	{name: "AllocateAddressSpacePrefix48",
		steps: []step{
			{kind: "outfile", args: []string{"--addrspace", "-abmnp", "fd0b:a336:4e7d::/48"}, file: "fd0b:a336:4e7d::-48"},
		}},
	{name: "TestHumanReadableGenericInfoNoPrefix",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "2a03:2880:20:4f06:face:b00c:0:1"}, file: "i-2a03:2880:20:4f06:face:b00c:0:1"},
		}},
	{name: "TestHumanReadableGenericInfoInternetIPv4",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "3.130.45.15"}, file: "i-3.130.45.15"},
		}},
	{name: "TestHumanReadableThisNetwork",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "0.0.0.1"}, file: "i-0.0.0.1"},
		}},
	{name: "TestHumanReadablePrivateUse1",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "10.0.0.1"}, file: "i-10.0.0.1"},
		}},
	{name: "TestHumanReadablePrivateUse2",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "172.16.0.1"}, file: "i-172.16.0.1"},
		}},
	{name: "TestHumanReadablePrivateUse3",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "172.31.255.254"}, file: "i-172.31.255.254"},
		}},
	{name: "TestHumanReadablePrivateUse4",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "192.168.0.1"}, file: "i-192.168.0.1"},
		}},
	{name: "TestHumanReadableSharedAddressSpace",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "100.64.0.1"}, file: "i-100.64.0.1"},
		}},
	{name: "TestHumanReadableLoopback",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "127.0.0.1"}, file: "i-127.0.0.1"},
		}},
	{name: "TestHumanReadableLoopback2",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "127.0.1.1"}, file: "i-127.0.1.1"},
		}},
	{name: "TestHumanReadableLinkLocal",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "169.254.0.1"}, file: "i-169.254.0.1"},
		}},
	{name: "TestHumanReadableIETFProtocolAssignmentsIPv4",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "192.0.0.254"}, file: "i-192.0.0.254"},
		}},
	{name: "TestHumanReadableServiceContinuityPrefixIPv4",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "192.0.0.6"}, file: "i-192.0.0.6"},
		}},
	{name: "TestHumanReadableDummyAddressIPv4",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "192.0.0.8"}, file: "i-192.0.0.8"},
		}},
	{name: "TestHumanReadablePortControlProtocolAnycastIPv4",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "192.0.0.9"}, file: "i-192.0.0.9"},
		}},
	{name: "TestHumanReadableTraversalUsingRelaysAroundNATAnycast",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "192.0.0.10"}, file: "i-192.0.0.10"},
		}},
	{name: "TestHumanReadableNAT64DNS64Discovery1",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "192.0.0.170"}, file: "i-192.0.0.170"},
		}},
	{name: "TestHumanReadableNAT64DNS64Discovery2",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "192.0.0.171"}, file: "i-192.0.0.171"},
		}},
	{name: "TestHumanReadableAS112v4",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "192.31.196.1"}, file: "i-192.31.196.1"},
		}},
	{name: "TestHumanReadableDeprecated6to4RelayAnycast",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "192.88.99.1"}, file: "i-192.88.99.1"},
		}},
	{name: "TestHumanReadableAMTIPv4",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "192.52.193.5"}, file: "i-192.52.193.5"},
		}},
	{name: "TestHumanReadableDirectDelegationAS112ServiceIPv4",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "192.175.48.12"}, file: "i-192.175.48.12"},
		}},
	{name: "TestHumanReadableTestNet1",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "192.0.2.1"}, file: "i-192.0.2.1"},
		}},
	{name: "TestHumanReadableTestNet2",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "198.51.100.1"}, file: "i-198.51.100.1"},
		}},
	{name: "TestHumanReadableTestNet3",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "203.0.113.1"}, file: "i-203.0.113.1"},
		}},
	{name: "TestHumanReadableBenchmarking1",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "198.18.0.1"}, file: "i-198.18.0.1"},
		}},
	{name: "TestHumanReadableBenchmarking2",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "198.19.201.71"}, file: "i-198.19.201.71"},
		}},
	{name: "TestHumanReadableBenchmarking3",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "198.19.255.254"}, file: "i-198.19.255.254"},
		}},
	{name: "TestHumanReadableMulticast",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "224.0.0.1"}, file: "i-224.0.0.1"},
		}},
	{name: "TestHumanReadableReserved1",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "240.0.0.1"}, file: "i-240.0.0.1"},
		}},
	{name: "TestHumanReadableReserved2",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "250.7.93.1"}, file: "i-250.7.93.1"},
		}},
	{name: "TestHumanReadableReserved3",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "255.255.255.254"}, file: "i-255.255.255.254"},
		}},
	{name: "TestHumanReadableLimitedBroadcast",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "255.255.255.255"}, file: "i-255.255.255.255"},
		}},
	{name: "TestHumanReadableTranslat96",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "64:ff9b::2100:951a"}, file: "i-64:ff9b::2100:951a"},
		}},
	{name: "TestHumanReadableTranslat48",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "64:ff9b:1:ea21:4316:55::1"}, file: "i-64:ff9b:1:ea21:4316:55::1"},
		}},
	{name: "TestHumanReadableTEREDO",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "2001:0000:50::d78a:0:c"}, file: "i-2001:0:50::d78a:0:c"},
		}},
	{name: "TestHumanReadablePortControlProtocolAnycast",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "2001:1::1"}, file: "i-2001:1::1"},
		}},
	{name: "TestHumanReadableTraversalUsingRelays",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "2001:1::2"}, file: "i-2001:1::2"},
		}},
	{name: "TestHumanReadableBenchmarking",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "2001:2:0:a::5:b7f"}, file: "i-2001:2:0:a::5:b7f"},
		}},
	{name: "TestHumanReadableAMT",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "2001:3:89:71::ab:1"}, file: "i-2001:3:89:71::ab:1"},
		}},
	{name: "TestHumanReadableAS112v6",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "2001:4:112::6483:bd12:9010"}, file: "i-2001:4:112::6483:bd12:9010"},
		}},
	{name: "TestHumanReadableORCHID",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "2001:1e:ab:7::"}, file: "i-2001:1e:ab:7::"},
		}},
	{name: "TestHumanReadableORCHIDv2",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "2001:2c:ab:7::"}, file: "i-2001:2c:ab:7::"},
		}},
	{name: "TestHumanReadableDocumentation",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "2001:db8:329::4a"}, file: "i-2001:db8:329::4a"},
		}},
	{name: "TestHumanReadableIETFProtocolAssignments",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "2001:190:70::e86a:0:b"}, file: "i-2001:190:70::e86a:0:b"},
		}},
	{name: "TestHumanReadable6to4",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "2002::34"}, file: "i-2002::34"},
		}},
	{name: "TestHumanReadableDirectDelegationAS112Service",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "2620:4f:8000::7372:ad1"}, file: "i-2620:4f:8000::7372:ad1"},
		}},
	{name: "TestHumanReadableUniqueLocalUnicast",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "fc00:0:a001::bcdf:4"}, file: "i-fc00:0:a001::bcdf:4"},
		}},
	{name: "TestHumanReadableLinkScopedUnicast",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "fe80::b:0:7acd:0:5"}, file: "i-fe80::b:0:7acd:0:5"},
		}},
	{name: "TestHumanReadableThisHostOnThisNetwork",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "0.0.0.0"}, file: "i-0.0.0.0"},
		}},
	{name: "TestHumanReadable6a44Relay",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "192.88.99.2"}, file: "i-192.88.99.2"},
		}},
	{name: "AddrspaceIPv4LargerThanBlock",
		steps: []step{
			{kind: "output", args: []string{"--addrspace", "192.168.0.0/8"}, want: "ADDRSPACE=Internet"},
		}},
	{name: "AddrspaceIPv4LargerThanPrivateBlock",
		steps: []step{
			{kind: "output", args: []string{"--addrspace", "10.0.0.0/7"}, want: "ADDRSPACE=Internet"},
		}},
	{name: "AddrspaceIPv4LargerThanThisNetwork",
		steps: []step{
			{kind: "output", args: []string{"--addrspace", "0.0.0.0/7"}, want: "ADDRSPACE=Internet"},
		}},
	{name: "TestHumanReadableDiscardOnly",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "100::1"}, file: "i-100::1"},
		}},
	{name: "AddrspaceNotDiscardOnly",
		steps: []step{
			{kind: "output", args: []string{"--addrspace", "1000::1"}, want: "ADDRSPACE=Reserved"},
		}},
	{name: "TestHumanReadableDummyIPv6Prefix",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "100:0:0:1::5"}, file: "i-100:0:0:1::5"},
		}},
	{name: "TestHumanReadableDNSSDSRPAnycast",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "2001:1::3"}, file: "i-2001:1::3"},
		}},
	{name: "TestHumanReadableDRIPDET",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "2001:30::1"}, file: "i-2001:30::1"},
		}},
	{name: "TestHumanReadableDocumentation3fff",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "3fff::1"}, file: "i-3fff::1"},
		}},
	{name: "TestHumanReadableSRv6SIDs",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "5f00::1"}, file: "i-5f00::1"},
		}},
	{name: "AddrspaceIPv6LargerThanBlock",
		steps: []step{
			{kind: "output", args: []string{"--addrspace", "2002::/15"}, want: "ADDRSPACE=\"Global Unicast\""},
		}},
	{name: "TestHumanReadableGenericInfoPrefix48",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "fd0b:a336:4e7d::/48"}, file: "i-fd0b:a336:4e7d::-48"},
		}},
	{name: "TestHelpCommand",
		steps: []step{
			{kind: "success", args: []string{"--help"}},
		}},
	{name: "GeoInfoAccepted",
		steps: []step{
			{kind: "success", args: []string{"-g", "127.0.0.1"}},
		}},
	{name: "GeoInfoIPv4",
		steps: []step{
			{kind: "outfile", bin: "geo", args: []string{"-g", "192.0.2.5"}, file: "geo-192.0.2.5"},
		}},
	{name: "GeoInfoIPv6",
		steps: []step{
			{kind: "outfile", bin: "geo", args: []string{"-g", "2001:db8::1"}, file: "geo-2001:db8::1"},
		}},
	{name: "GeoInfoNoCountry",
		steps: []step{
			{kind: "outfile", bin: "geo", args: []string{"-g", "198.51.100.200"}, file: "geo-198.51.100.200"},
		}},
	{name: "GeoInfoRegisteredCountry",
		steps: []step{
			{kind: "outfile", bin: "geo", args: []string{"-g", "198.51.100.5"}, file: "geo-198.51.100.5"},
		}},
	{name: "GeoInfoRepresentedCountry",
		steps: []step{
			{kind: "outfile", bin: "geo", args: []string{"-g", "198.51.100.130"}, file: "geo-198.51.100.130"},
		}},
	{name: "GeoInfoASNOnly",
		steps: []step{
			{kind: "outfile", bin: "geo", args: []string{"-g", "203.0.113.1"}, file: "geo-203.0.113.1"},
		}},
	{name: "GeoInfoAllInfo",
		steps: []step{
			{kind: "outfile", bin: "geo", args: []string{"--all-info", "192.0.2.5"}, file: "geo-all-info-192.0.2.5"},
		}},
	{name: "GeoInfoNotFound",
		steps: []step{
			{kind: "outfile", bin: "geo", args: []string{"--all-info", "10.0.0.1"}, file: "geo-all-info-10.0.0.1"},
		}},
	{name: "GeoInfoJson",
		steps: []step{
			{kind: "outfile", bin: "geo", args: []string{"--format=json", "-g", "2001:db8::1"}, file: "geo-json-2001:db8::1"},
		}},
	{name: "GeoInfoNetworkWithinRecords",
		steps: []step{
			{kind: "outfile", bin: "geo", args: []string{"-g", "192.0.2.0/24"}, file: "geo-192.0.2.0-24"},
		}},
	{name: "GeoInfoNetworkLargerThanCity",
		steps: []step{
			{kind: "outfile", bin: "geo", args: []string{"-g", "192.0.2.77/23"}, file: "geo-192.0.2.77-23"},
		}},
	{name: "GeoInfoNetworkLargerThanCityIPv6",
		steps: []step{
			{kind: "outfile", bin: "geo", args: []string{"-g", "2001:db8::/32"}, file: "geo-2001:db8::-32"},
		}},
	{name: "GeoInfoNetworkLargerThanRecords",
		steps: []step{
			{kind: "outfile", bin: "geo", args: []string{"--all-info", "192.0.0.0/16"}, file: "geo-all-info-192.0.0.0-16"},
		}},
	{name: "TestShortHelpCommand",
		steps: []step{
			{kind: "success", args: []string{"-?"}},
		}},
	{name: "UnknownLongOption",
		steps: []step{
			{kind: "failure", args: []string{"--bogus", "1.2.3.4"}},
		}},
	{name: "UnknownShortOption",
		steps: []step{
			{kind: "failure", args: []string{"-y", "1.2.3.4"}},
		}},
	{name: "MissingOptionArgument",
		steps: []step{
			{kind: "failure", args: []string{"-S"}},
		}},
	{name: "ValidateLoopbackIPv4",
		steps: []step{
			{kind: "success", args: []string{"-c", "127.0.0.1"}},
		}},
	{name: "ValidateLoopbackExplicitIPv6",
		steps: []step{
			{kind: "success", args: []string{"-c", "-6", "::1"}},
		}},
	{name: "ValidateLoopbackIPv6",
		steps: []step{
			{kind: "success", args: []string{"-c", "::1"}},
		}},
	{name: "ValidatePrivateIPv4",
		steps: []step{
			{kind: "success", args: []string{"-c", "192.168.1.1"}},
		}},
	{name: "ValidateAbbreviatedGlobalIPv6",
		steps: []step{
			{kind: "success", args: []string{"-c", "-6", "2a01:198:200:300::2"}},
		}},
	{name: "ValidateExpandedGlobalIPv6",
		steps: []step{
			{kind: "success", args: []string{"-c", "-6", "2a01:198:200:300:0000:0000:0000:2"}},
		}},
	{name: "ValidateCompleteGlobalIPv6",
		steps: []step{
			{kind: "success", args: []string{"-c", "-6", "2a01:0198:0200:0300:0000:0000:0000:0002"}},
		}},
	{name: "ValidateLoopbackIPv6WithPrefix",
		steps: []step{
			{kind: "success", args: []string{"-c", "-6", "::1/128"}},
		}},
	{name: "ValidateReservedIPv6WithPrefix",
		steps: []step{
			{kind: "success", args: []string{"-c", "-6", "fec0::1:0:0:c0a8:8002/64"}},
		}},
	{name: "NoValidateInvalidHexIPv6",
		steps: []step{
			{kind: "failure", args: []string{"-c", "-6", "gggg::"}},
		}},
	{name: "NoAllowIPv4AndIPv6",
		steps: []step{
			{kind: "failure", args: []string{"-4", "-6", "2a01:198:200:300:0000:0000:0000:2"}},
		}},
	{name: "NoValidateIPv6InIPv4Mode",
		steps: []step{
			{kind: "failure", args: []string{"-c", "-4", "2a01:198:200:300:0000:0000:0000:2"}},
		}},
	{name: "NoValidateIPv4AndIPv6WithIPv4Address",
		steps: []step{
			{kind: "failure", args: []string{"-c", "-4", "-6", "127.0.0.1"}},
		}},
	{name: "NoValidateInvalidPrefixIPv6",
		steps: []step{
			{kind: "failure", args: []string{"-c", "::1/999"}},
		}},
	{name: "NoValidateNegativePrefix",
		steps: []step{
			{kind: "failure", args: []string{"-m", "192.168.1.1/-1"}},
		}},
	{name: "NoValidateIPv6PrefixInIPv4Mode",
		steps: []step{
			{kind: "failure", args: []string{"-m", "192.168.1.1/64"}},
		}},
	{name: "NoValidateInvalidPrefixIPv4",
		steps: []step{
			{kind: "failure", args: []string{"-m", "192.168.1.1/99999"}},
		}},
	{name: "CalculateBroadcastAddressFromPrefixIPv4",
		steps: []step{
			{kind: "output", args: []string{"-b", "192.168.1.1/24"}, want: "BROADCAST=192.168.1.255"},
		}},
	{name: "CalculateBroadcastAddressFromNetmaskIPv4",
		steps: []step{
			{kind: "output", args: []string{"-b", "192.168.1.1", "255.255.255.0"}, want: "BROADCAST=192.168.1.255"},
		}},
	{name: "CalculateNetmaskPrefix24",
		steps: []step{
			{kind: "output", args: []string{"-m", "192.168.1.1/24"}, want: "NETMASK=255.255.255.0"},
		}},
	{name: "CalculateNetmaskPrefix22",
		steps: []step{
			{kind: "output", args: []string{"-m", "172.16.59.222/22"}, want: "NETMASK=255.255.252.0"},
		}},
	{name: "CalculateNetmaskPrefix1",
		steps: []step{
			{kind: "output", args: []string{"-m", "10.0.0.0/1"}, want: "NETMASK=128.0.0.0"},
		}},
	{name: "CalculateNetmaskNoPrefix",
		steps: []step{
			{kind: "output", args: []string{"-m", "192.168.1.1"}, want: "NETMASK=255.255.255.255"},
		}},
	{name: "CalculatePrefix24FromNetmask",
		steps: []step{
			{kind: "output", args: []string{"-p", "192.168.1.1", "255.255.255.0"}, want: "PREFIX=24"},
		}},
	{name: "CalculatePrefix32FromNetmask",
		steps: []step{
			{kind: "output", args: []string{"-p", "192.168.1.1", "255.255.255.255"}, want: "PREFIX=32"},
		}},
	{name: "CalculatePrefix0FromNetmask",
		steps: []step{
			{kind: "output", args: []string{"-p", "192.168.1.1", "0.0.0.0"}, want: "PREFIX=0"},
		}},
	{name: "CalculatePrefix22FromNetmask",
		steps: []step{
			{kind: "output", args: []string{"-p", "172.16.59.222", "255.255.252.0"}, want: "PREFIX=22"},
		}},
	{name: "CalculateClassNetmask24",
		steps: []step{
			{kind: "output", args: []string{"--class-prefix", "-m", "192.168.1.1"}, want: "NETMASK=255.255.255.0"},
		}},
	{name: "CalculateClassNetmask8",
		steps: []step{
			{kind: "output", args: []string{"--class-prefix", "-m", "10.1.2.3"}, want: "NETMASK=255.0.0.0"},
		}},
	{name: "CalculateClassNetmask16",
		steps: []step{
			{kind: "output", args: []string{"--class-prefix", "-m", "129.22.4.3"}, want: "NETMASK=255.255.0.0"},
		}},
	{name: "CalculateNetworkFromPrefix32",
		steps: []step{
			{kind: "output", args: []string{"-n", "192.168.1.1/32"}, want: "NETWORK=192.168.1.1"},
		}},
	{name: "CalculateNetworkFromPrefix0",
		steps: []step{
			{kind: "output", args: []string{"-n", "192.168.1.1/0"}, want: "NETWORK=0.0.0.0"},
		}},
	{name: "AddressesPrefix0IPv4",
		steps: []step{
			{kind: "output", args: []string{"--addresses", "0.0.0.0/0"}, want: "ADDRESSES=4294967294"},
		}},
	{name: "AddressesPrefix1IPv4",
		steps: []step{
			{kind: "output", args: []string{"--addresses", "10.0.0.0/1"}, want: "ADDRESSES=2147483646"},
		}},
	{name: "AddressesPrefix0IPv6",
		steps: []step{
			{kind: "output", args: []string{"--addresses", "::/0"}, want: "ADDRESSES=340282366920938463463374607431768211456"},
		}},
	{name: "LookupReverseDNSFromPrefix32IPv4",
		steps: []step{
			{kind: "output", args: []string{"--reverse-dns", "193.92.150.3/32"}, want: "REVERSEDNS=3.150.92.193.in-addr.arpa."},
		}},
	{name: "LookupReverseDNSFromPrefix26IPv4",
		steps: []step{
			{kind: "output", args: []string{"--reverse-dns", "193.92.150.3/26"}, want: "REVERSEDNS=0-63.150.92.193.in-addr.arpa."},
		}},
	{name: "LookupReverseDNSFromPrefix16IPv4",
		steps: []step{
			{kind: "output", args: []string{"--reverse-dns", "193.92.150.3/16"}, want: "REVERSEDNS=92.193.in-addr.arpa."},
		}},
	{name: "LookupReverseDNSFromPrefix12IPv4",
		steps: []step{
			{kind: "output", args: []string{"--reverse-dns", "193.92.150.3/12"}, want: "REVERSEDNS=80-95.193.in-addr.arpa."},
		}},
	{name: "LookupReverseDNSFromPrefix4IPv4",
		steps: []step{
			{kind: "output", args: []string{"--reverse-dns", "193.92.150.3/4"}, want: "REVERSEDNS=192-207.in-addr.arpa."},
		}},
	{name: "LookupReverseDNSFromPrefix0IPv4",
		steps: []step{
			{kind: "output", args: []string{"--reverse-dns", "193.92.150.3/0"}, want: "REVERSEDNS=in-addr.arpa."},
		}},
	{name: "LookupReverseDNS4321IPv6",
		steps: []step{
			{kind: "output", args: []string{"--reverse-dns", "4321:0:1:2:3:4:567:89ab"}, want: "REVERSEDNS=b.a.9.8.7.6.5.0.4.0.0.0.3.0.0.0.2.0.0.0.1.0.0.0.0.0.0.0.1.2.3.4.ip6.arpa."},
		}},
	{name: "LookupReverseDNSFirstIPv6",
		steps: []step{
			{kind: "output", args: []string{"--reverse-dns", "1::"}, want: "REVERSEDNS=0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.1.0.0.0.ip6.arpa."},
		}},
	{name: "LookupReverseDNSLoopbackIPv6",
		steps: []step{
			{kind: "output", args: []string{"--reverse-dns", "::1"}, want: "REVERSEDNS=1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.ip6.arpa."},
		}},
	{name: "LookupReverseDNSFromPrefix128IPv6",
		steps: []step{
			{kind: "output", args: []string{"--reverse-dns", "1234::4321/128"}, want: "REVERSEDNS=1.2.3.4.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.4.3.2.1.ip6.arpa."},
		}},
	{name: "LookupReverseDNSFromPrefix124IPv6",
		steps: []step{
			{kind: "output", args: []string{"--reverse-dns", "1234::4321/124"}, want: "REVERSEDNS=2.3.4.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.4.3.2.1.ip6.arpa."},
		}},
	{name: "LookupReverseDNSFromPrefix120IPv6",
		steps: []step{
			{kind: "output", args: []string{"--reverse-dns", "1234::4321/120"}, want: "REVERSEDNS=3.4.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.4.3.2.1.ip6.arpa."},
		}},
	{name: "LookupReverseDNSFromPrefix0IPv6",
		steps: []step{
			{kind: "output", args: []string{"--reverse-dns", "::/0"}, want: "REVERSEDNS=ip6.arpa."},
		}},
	{name: "LookupReverseDNSFromPrefix31IPv6",
		steps: []step{
			{kind: "output", args: []string{"--reverse-dns", "2001:db8::/31"}, want: "REVERSEDNS=\"8.b.d.0.1.0.0.2.ip6.arpa. 9.b.d.0.1.0.0.2.ip6.arpa.\""},
		}},
	{name: "LookupReverseDNSFromPrefix33IPv6",
		steps: []step{
			{kind: "output", args: []string{"--reverse-dns", "2001:db8:8000::/33"}, want: "REVERSEDNS=\"8.8.b.d.0.1.0.0.2.ip6.arpa. 9.8.b.d.0.1.0.0.2.ip6.arpa. a.8.b.d.0.1.0.0.2.ip6.arpa. b.8.b.d.0.1.0.0.2.ip6.arpa. c.8.b.d.0.1.0.0.2.ip6.arpa. d.8.b.d.0.1.0.0.2.ip6.arpa. e.8.b.d.0.1.0.0.2.ip6.arpa. f.8.b.d.0.1.0.0.2.ip6.arpa.\""},
		}},
	{name: "LookupReverseDNSFromPrefix34IPv6",
		steps: []step{
			{kind: "output", args: []string{"--reverse-dns", "2001:db8:4000::/34"}, want: "REVERSEDNS=\"4.8.b.d.0.1.0.0.2.ip6.arpa. 5.8.b.d.0.1.0.0.2.ip6.arpa. 6.8.b.d.0.1.0.0.2.ip6.arpa. 7.8.b.d.0.1.0.0.2.ip6.arpa.\""},
		}},
	{name: "LookupReverseDNSFromPrefix127IPv6",
		steps: []step{
			{kind: "output", args: []string{"--reverse-dns", "1234::4321/127"}, want: "REVERSEDNS=\"0.2.3.4.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.4.3.2.1.ip6.arpa. 1.2.3.4.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.4.3.2.1.ip6.arpa.\""},
		}},
	{name: "LookupReverseDNSUnmaskedPrefix35IPv6",
		steps: []step{
			{kind: "output", args: []string{"--reverse-dns", "2001:db8:3fff::1/35"}, want: "REVERSEDNS=\"2.8.b.d.0.1.0.0.2.ip6.arpa. 3.8.b.d.0.1.0.0.2.ip6.arpa.\""},
		}},
	{name: "LookupReverseDNSNoDecoratePrefix37IPv6",
		steps: []step{
			{kind: "outfile", args: []string{"--reverse-dns", "--no-decorate", "2001:db8:abcd::/37"}, file: "reverse-dns-no-decorate-2001:db8:abcd::-37"},
		}},
	{name: "AllInfoReverseDNSPrefix31IPv6",
		steps: []step{
			{kind: "outfile", args: []string{"--all-info", "2001:db8::/31"}, file: "all-info-2001:db8::-31"},
		}},
	{name: "JsonAllInfoReverseDNSPrefix31IPv6",
		steps: []step{
			{kind: "outfile", args: []string{"--all-info", "--json", "2001:db8::/31"}, file: "json-all-info-2001:db8::-31"},
		}},
	{name: "JsonAllInfoReverseDNSPrefix32IPv6",
		steps: []step{
			{kind: "outfile", args: []string{"--all-info", "--json", "2001:db8::/32"}, file: "json-all-info-2001:db8::-32"},
		}},
	{name: "JsonAllInfoReverseDNSPrefix24IPv4",
		steps: []step{
			{kind: "outfile", args: []string{"--all-info", "--json", "192.168.2.7/24"}, file: "json-all-info-192.168.2.7-24"},
		}},
	{name: "SpecificInfoOutput",
		steps: []step{
			{kind: "outfile", args: []string{"-abmnp", "--minaddr", "--maxaddr", "--addresses", "192.168.2.6/24"}, file: "192.168.2.6"},
		}},
	{name: "SpecificInfoOutputNoDecorate",
		steps: []step{
			{kind: "outfile", args: []string{"-abmnp", "--minaddr", "--maxaddr", "--addresses", "--no-decorate", "192.168.2.7/24"}, file: "192.168.2.7"},
		}},
	{name: "SpecificInfoOutputNoDecorateQuotedAddrspace",
		steps: []step{
			{kind: "output", args: []string{"--addrspace", "--no-decorate", "10.0.0.1"}, want: "\"Private Use\""},
		}},
	{name: "JsonSplitPrefix24",
		steps: []step{
			{kind: "outfile", args: []string{"-j", "-S", "24", "10.10.10.0/16"}, file: "json-split-10.10.10.0-16-24"},
		}},
	{name: "JsonSplitPrefix26",
		steps: []step{
			{kind: "outfile", args: []string{"-j", "-S", "26", "192.168.5.45/24"}, file: "json-split-192.168.5.45-24-26"},
		}},
	{name: "JsonNoDecorateSplitPrefix24",
		steps: []step{
			{kind: "outfile", args: []string{"-j", "--no-decorate", "-S", "24", "10.10.10.0/16"}, file: "json-nsplit-10.10.10.0-16-24"},
		}},
	{name: "JsonNoDecorateSplitPrefix26",
		steps: []step{
			{kind: "outfile", args: []string{"-j", "--no-decorate", "-S", "26", "192.168.5.45/24"}, file: "json-nsplit-192.168.5.45-24-26"},
		}},
	{name: "JsonSplitPrefix64",
		steps: []step{
			{kind: "outfile", args: []string{"-j", "-S", "64", "2a03:2880:20:4f06:face::/56"}, file: "json-nsplit-2a03:2880:20:4f06:face::-56-64"},
		}},
	{name: "JsonNoDecorateSplitPrefix64",
		steps: []step{
			{kind: "outfile", args: []string{"-j", "--no-decorate", "-S", "64", "2a03:2880:20:4f06:face::/56"}, file: "json-nsplit-2a03:2880:20:4f06:face::-56-64"},
		}},
	{name: "JsonTestHumanReadableGenericInfoPrefix56",
		steps: []step{
			{kind: "outfile", args: []string{"-j", "-i", "2a03:2880:20:4f06:face:b00c:0:1/56"}, file: "json-i-2a03:2880:20:4f06:face:b00c:0:1-56"},
		}},
	{name: "JsonTestHumanReadableSpecificInfo",
		steps: []step{
			{kind: "outfile", args: []string{"-j", "-n", "2a03:2880:20:4f06:face:b00c:0:1/56"}, file: "json-i-2a03:2880:20:4f06:face:b00c:0:1-56"},
		}},
	{name: "AllInfo",
		steps: []step{
			{kind: "outfile", args: []string{"--all-info", "192.168.2.7/24"}, file: "all-info-192.168.2.7"},
		}},
	{name: "AllInfoClassDLowest",
		steps: []step{
			{kind: "outfile", args: []string{"--all-info", "224.0.0.1"}, file: "all-info-224.0.0.1"},
		}},
	{name: "AllInfoClassDHighestOctet",
		steps: []step{
			{kind: "outfile", args: []string{"--all-info", "239.1.1.1"}, file: "all-info-239.1.1.1"},
		}},
	{name: "AllInfoClassE",
		steps: []step{
			{kind: "outfile", args: []string{"--all-info", "240.0.0.1"}, file: "all-info-240.0.0.1"},
		}},
	{name: "AllInfoIIDLinkLocalUniversal",
		steps: []step{
			{kind: "outfile", args: []string{"--all-info", "fe80::21b:21ff:fe3a:5c7d/64"}, file: "all-info-fe80::21b:21ff:fe3a:5c7d-64"},
		}},
	{name: "JsonAllInfoIIDLinkLocalUniversal",
		steps: []step{
			{kind: "outfile", args: []string{"--all-info", "--json", "fe80::21b:21ff:fe3a:5c7d/64"}, file: "json-all-info-fe80::21b:21ff:fe3a:5c7d-64"},
		}},
	{name: "NoDecorateAllInfoIIDLinkLocalUniversal",
		steps: []step{
			{kind: "outfile", args: []string{"--all-info", "--no-decorate", "fe80::21b:21ff:fe3a:5c7d/64"}, file: "no-decorate-all-info-fe80::21b:21ff:fe3a:5c7d-64"},
		}},
	{name: "AllInfoIIDLinkLocalLocal",
		steps: []step{
			{kind: "outfile", args: []string{"--all-info", "fe80::ff:fe00:1"}, file: "all-info-fe80::ff:fe00:1"},
		}},
	{name: "AllInfoIIDGroup",
		steps: []step{
			{kind: "outfile", args: []string{"--all-info", "2001:db8::300:ff:fe00:1"}, file: "all-info-2001:db8::300:ff:fe00:1"},
		}},
	{name: "AllInfoIIDNotEUI64",
		steps: []step{
			{kind: "outfile", args: []string{"--all-info", "2001:db8::f00d:1234:5678:9abc/64"}, file: "all-info-2001:db8::f00d:1234:5678:9abc-64"},
		}},
	{name: "AllInfoIIDNetwork",
		steps: []step{
			{kind: "outfile", args: []string{"--all-info", "fe80::/64"}, file: "all-info-fe80::-64"},
		}},
	{name: "AllInfoIIDMapped",
		steps: []step{
			{kind: "outfile", args: []string{"--all-info", "::ffff:192.0.2.1"}, file: "all-info-::ffff:192.0.2.1"},
		}},
	{name: "AllInfoIIDMulticast",
		steps: []step{
			{kind: "outfile", args: []string{"--all-info", "ff02::21b:21ff:fe3a:5c7d"}, file: "all-info-ff02::21b:21ff:fe3a:5c7d"},
		}},
	{name: "InfoIIDLinkLocal",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "fe80::21b:21ff:fe3a:5c7d/64"}, file: "i-fe80::21b:21ff:fe3a:5c7d-64"},
		}},
	{name: "Info",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "192.168.2.7/24"}, file: "info-192.168.2.7"},
		}},
	{name: "InfoImplicit",
		steps: []step{
			{kind: "outfile", args: []string{"192.168.2.7/24"}, file: "info-192.168.2.7"},
		}},
	{name: "DeaggregateNoDecorate",
		steps: []step{
			{kind: "outfile", args: []string{"--no-decorate", "-d", "192.168.2.7-192.168.2.13"}, file: "deaggregate-192.168.2.7-192.168.2.13"},
		}},
	{name: "Deaggregate",
		steps: []step{
			{kind: "outfile", args: []string{"-d", "192.168.2.0-192.168.3.255"}, file: "deaggregate-192.168.2.0-192.168.3.255"},
		}},
	{name: "DeaggregateLargeNum",
		steps: []step{
			{kind: "outfile", args: []string{"-d", "240.0.0.0-255.255.255.255"}, file: "deaggregate-240.0.0.0-255.255.255.255"},
		}},
	{name: "DeaggregateLargeNum2",
		steps: []step{
			{kind: "outfile", args: []string{"-d", "127.0.0.0-255.255.255.255"}, file: "deaggregate-127.0.0.0-255.255.255.255"},
		}},
	{name: "DeaggregateLargeNum3",
		steps: []step{
			{kind: "outfile", args: []string{"-d", "0.0.0.0-255.255.255.255"}, file: "deaggregate-0.0.0.0-255.255.255.255"},
		}},
	{name: "DeaggregateSpaces",
		steps: []step{
			{kind: "outfile", args: []string{"-d", "  192.168.2.0   - 192.168.3.255   "}, file: "deaggregate-192.168.2.0-192.168.3.255"},
		}},
	{name: "DeaggregateJson",
		steps: []step{
			{kind: "outfile", args: []string{"-j", "-d", "192.168.2.33-192.168.3.2"}, file: "deaggregate-192.168.2.33-192.168.3.2"},
		}},
	{name: "DeaggregateIPv6NoDecorate",
		steps: []step{
			{kind: "outfile", args: []string{"--no-decorate", "-d", "2a03:2880:20:4f06:face::0-2a03:2880:20:4f06:face::fffe"}, file: "deaggregate-2a03:2880:20:4f06:face::0-2a03:2880:20:4f06:face::fffe"},
		}},
	{name: "DeaggregateIPv6",
		steps: []step{
			{kind: "outfile", args: []string{"-d", "fcd3:57d1:733:c18f:b498:25e1:788:0-fcd3:57d1:733:c18f:b498:25e1:788:ffff"}, file: "deaggregate-fcd3:57d1:733:c18f:b498:25e1:788:0-fcd3:57d1:733:c18f:b498:25e1:788:ffff"},
		}},
	{name: "DeaggregateIPv6Json",
		steps: []step{
			{kind: "outfile", args: []string{"-j", "-d", "fcd3:57d1:733:c18f:b498:25e1:788:f-fcd3:57d1:733:c18f:b498:25e1:789:ffa9"}, file: "deaggregate-fcd3:57d1:733:c18f:b498:25e1:788:f-fcd3:57d1:733:c18f:b498:25e1:789:ffa9"},
		}},
	{name: "DeaggregateIPv6All",
		steps: []step{
			{kind: "outfile", args: []string{"-d", "::-ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"}, file: "deaggregate-::-ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"},
		}},
	{name: "DeaggregateIPv6UpperHalf",
		steps: []step{
			{kind: "outfile", args: []string{"-d", "8000::-ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"}, file: "deaggregate-8000::-ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"},
		}},
	{name: "DeaggregateIPv6LastAddress",
		steps: []step{
			{kind: "outfile", args: []string{"-d", "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff-ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"}, file: "deaggregate-ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff-ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"},
		}},
	{name: "DeaggregateIPv6LowerHalf",
		steps: []step{
			{kind: "outfile", args: []string{"-d", "::-8000::"}, file: "deaggregate-::-8000::"},
		}},
	{name: "StandardInfo",
		steps: []step{
			{kind: "outfile", args: []string{"192.168.2.0/24"}, file: "standard-192.168.2.0-24"},
		}},
	{name: "NoDecorateInfo",
		steps: []step{
			{kind: "outfile", args: []string{"--no-decorate", "192.168.2.0/24"}, file: "no-decorate-192.168.2.0-24"},
		}},
	{name: "NoDecorateAllInfo",
		steps: []step{
			{kind: "outfile", args: []string{"--all-info", "--no-decorate", "192.168.2.0/24"}, file: "no-decorate-all-info-192.168.2.0-24"},
		}},
	{name: "EqualsHostBits",
		steps: []step{
			{kind: "status", args: []string{"192.168.7.15/24", "--equals=192.168.7.5/24"}, code: 0},
			{kind: "output", args: []string{"192.168.7.15/24", "--equals=192.168.7.5/24"}, want: "The networks are equal"},
		}},
	{name: "EqualsNetmask",
		steps: []step{
			{kind: "status", args: []string{"192.168.7.15", "255.255.255.0", "--equals=192.168.7.5/255.255.255.0"}, code: 0},
			{kind: "output", args: []string{"192.168.7.15", "255.255.255.0", "--equals=192.168.7.5/255.255.255.0"}, want: "The networks are equal"},
		}},
	{name: "EqualsClassPrefix",
		steps: []step{
			{kind: "status", args: []string{"--class-prefix", "10.0.0.0/8", "--equals=10.9.9.9"}, code: 0},
			{kind: "output", args: []string{"--class-prefix", "10.0.0.0/8", "--equals=10.9.9.9"}, want: "The networks are equal"},
		}},
	{name: "EqualsJson",
		steps: []step{
			{kind: "status", args: []string{"-j", "192.168.7.15/24", "--equals=192.168.7.5/24"}, code: 0},
			{kind: "output", args: []string{"-j", "192.168.7.15/24", "--equals=192.168.7.5/24"}, want: ""},
		}},
	{name: "EqualsNoDecorate",
		steps: []step{
			{kind: "status", args: []string{"--no-decorate", "192.168.7.15/24", "--equals=192.168.7.5/25"}, code: 1},
			{kind: "output", args: []string{"--no-decorate", "192.168.7.15/24", "--equals=192.168.7.5/25"}, want: ""},
		}},
	{name: "SubnetOfNoDecorate",
		steps: []step{
			{kind: "status", args: []string{"--no-decorate", "192.168.0.1/24", "--subnet-of=192.168.1.15/23"}, code: 0},
			{kind: "output", args: []string{"--no-decorate", "192.168.0.1/24", "--subnet-of=192.168.1.15/23"}, want: ""},
		}},
	{name: "OverlapsSilent",
		steps: []step{
			{kind: "status", args: []string{"-s", "10.0.0.0/8", "--overlaps=::/0"}, code: 1},
			{kind: "output", args: []string{"-s", "10.0.0.0/8", "--overlaps=::/0"}, want: ""},
		}},
	{name: "EqualsHost32",
		steps: []step{
			{kind: "status", args: []string{"192.168.1.15", "--equals=192.168.1.15/32"}, code: 0},
			{kind: "output", args: []string{"192.168.1.15", "--equals=192.168.1.15/32"}, want: "The networks are equal"},
		}},
	{name: "EqualsHostBitsIPv6",
		steps: []step{
			{kind: "status", args: []string{"2001:db8::1/64", "--equals=2001:db8::ffff/64"}, code: 0},
			{kind: "output", args: []string{"2001:db8::1/64", "--equals=2001:db8::ffff/64"}, want: "The networks are equal"},
		}},
	{name: "EqualsPrefix0",
		steps: []step{
			{kind: "status", args: []string{"1.2.3.4/0", "--equals=0.0.0.0/0"}, code: 0},
			{kind: "output", args: []string{"1.2.3.4/0", "--equals=0.0.0.0/0"}, want: "The networks are equal"},
		}},
	{name: "EqualsDifferentPrefix",
		steps: []step{
			{kind: "status", args: []string{"192.168.0.0/24", "--equals=192.168.0.0/23"}, code: 1},
			{kind: "output", args: []string{"192.168.0.0/24", "--equals=192.168.0.0/23"}, want: "The networks are not equal"},
		}},
	{name: "EqualsDisjoint",
		steps: []step{
			{kind: "status", args: []string{"192.168.1.15/24", "--equals=192.168.7.1/24"}, code: 1},
			{kind: "output", args: []string{"192.168.1.15/24", "--equals=192.168.7.1/24"}, want: "The networks are not equal"},
		}},
	{name: "EqualsDifferentPrefixIPv6",
		steps: []step{
			{kind: "status", args: []string{"2001:db8::/64", "--equals=2001:db8::/65"}, code: 1},
			{kind: "output", args: []string{"2001:db8::/64", "--equals=2001:db8::/65"}, want: "The networks are not equal"},
		}},
	{name: "SubnetOf",
		steps: []step{
			{kind: "status", args: []string{"192.168.0.1/24", "--subnet-of=192.168.1.15/23"}, code: 0},
			{kind: "output", args: []string{"192.168.0.1/24", "--subnet-of=192.168.1.15/23"}, want: "192.168.0.0/24 is a subnet of 192.168.0.0/23"},
		}},
	{name: "SubnetOfEqual",
		steps: []step{
			{kind: "status", args: []string{"10.0.0.0/8", "--subnet-of=10.0.0.0/8"}, code: 0},
			{kind: "output", args: []string{"10.0.0.0/8", "--subnet-of=10.0.0.0/8"}, want: "10.0.0.0/8 is a subnet of 10.0.0.0/8"},
		}},
	{name: "SubnetOfPartialByte",
		steps: []step{
			{kind: "status", args: []string{"10.64.0.0/10", "--subnet-of=10.0.0.0/9"}, code: 0},
			{kind: "output", args: []string{"10.64.0.0/10", "--subnet-of=10.0.0.0/9"}, want: "10.64.0.0/10 is a subnet of 10.0.0.0/9"},
		}},
	{name: "SubnetOfPrefix0",
		steps: []step{
			{kind: "status", args: []string{"10.1.2.3", "--subnet-of=0.0.0.0/0"}, code: 0},
			{kind: "output", args: []string{"10.1.2.3", "--subnet-of=0.0.0.0/0"}, want: "10.1.2.3/32 is a subnet of 0.0.0.0/0"},
		}},
	{name: "SubnetOfIPv6",
		steps: []step{
			{kind: "status", args: []string{"2001:db8:1::/48", "--subnet-of=2001:db8::/32"}, code: 0},
			{kind: "output", args: []string{"2001:db8:1::/48", "--subnet-of=2001:db8::/32"}, want: "2001:db8:1::/48 is a subnet of 2001:db8::/32"},
		}},
	{name: "SubnetOfPrefix0IPv6",
		steps: []step{
			{kind: "status", args: []string{"::1", "--subnet-of=::/0"}, code: 0},
			{kind: "output", args: []string{"::1", "--subnet-of=::/0"}, want: "::1/128 is a subnet of ::/0"},
		}},
	{name: "SubnetOfSupernet",
		steps: []step{
			{kind: "status", args: []string{"192.168.1.15/23", "--subnet-of=192.168.0.1/24"}, code: 1},
			{kind: "output", args: []string{"192.168.1.15/23", "--subnet-of=192.168.0.1/24"}, want: "192.168.0.0/23 is not a subnet of 192.168.0.0/24"},
		}},
	{name: "SubnetOfSiblingPartialByte",
		steps: []step{
			{kind: "status", args: []string{"10.128.0.0/10", "--subnet-of=10.0.0.0/9"}, code: 1},
			{kind: "output", args: []string{"10.128.0.0/10", "--subnet-of=10.0.0.0/9"}, want: "10.128.0.0/10 is not a subnet of 10.0.0.0/9"},
		}},
	{name: "SubnetOfPrefix0Supernet",
		steps: []step{
			{kind: "status", args: []string{"0.0.0.0/0", "--subnet-of=0.0.0.0/1"}, code: 1},
			{kind: "output", args: []string{"0.0.0.0/0", "--subnet-of=0.0.0.0/1"}, want: "0.0.0.0/0 is not a subnet of 0.0.0.0/1"},
		}},
	{name: "SubnetOfSupernetIPv6",
		steps: []step{
			{kind: "status", args: []string{"2001:db8::/32", "--subnet-of=2001:db8:1::/48"}, code: 1},
			{kind: "output", args: []string{"2001:db8::/32", "--subnet-of=2001:db8:1::/48"}, want: "2001:db8::/32 is not a subnet of 2001:db8:1::/48"},
		}},
	{name: "OverlapsSupernet",
		steps: []step{
			{kind: "status", args: []string{"192.168.1.15/23", "--overlaps=192.168.0.1/24"}, code: 0},
			{kind: "output", args: []string{"192.168.1.15/23", "--overlaps=192.168.0.1/24"}, want: "The networks overlap"},
		}},
	{name: "OverlapsSubnet",
		steps: []step{
			{kind: "status", args: []string{"192.168.0.1/24", "--overlaps=192.168.1.15/23"}, code: 0},
			{kind: "output", args: []string{"192.168.0.1/24", "--overlaps=192.168.1.15/23"}, want: "The networks overlap"},
		}},
	{name: "OverlapsHost",
		steps: []step{
			{kind: "status", args: []string{"10.0.0.1", "--overlaps=10.0.0.1"}, code: 0},
			{kind: "output", args: []string{"10.0.0.1", "--overlaps=10.0.0.1"}, want: "The networks overlap"},
		}},
	{name: "OverlapsPrefix0IPv6",
		steps: []step{
			{kind: "status", args: []string{"::/0", "--overlaps=2001:db8::1"}, code: 0},
			{kind: "output", args: []string{"::/0", "--overlaps=2001:db8::1"}, want: "The networks overlap"},
		}},
	{name: "OverlapsDisjoint",
		steps: []step{
			{kind: "status", args: []string{"192.168.1.15/24", "--overlaps=192.168.7.1/24"}, code: 1},
			{kind: "output", args: []string{"192.168.1.15/24", "--overlaps=192.168.7.1/24"}, want: "The networks do not overlap"},
		}},
	{name: "OverlapsAdjacent",
		steps: []step{
			{kind: "status", args: []string{"192.168.0.0/24", "--overlaps=192.168.1.0/24"}, code: 1},
			{kind: "output", args: []string{"192.168.0.0/24", "--overlaps=192.168.1.0/24"}, want: "The networks do not overlap"},
		}},
	{name: "OverlapsAdjacentIPv6",
		steps: []step{
			{kind: "status", args: []string{"2001:db8::/33", "--overlaps=2001:db8:8000::/33"}, code: 1},
			{kind: "output", args: []string{"2001:db8::/33", "--overlaps=2001:db8:8000::/33"}, want: "The networks do not overlap"},
		}},
	{name: "OverlapsMixedFamily",
		steps: []step{
			{kind: "status", args: []string{"10.0.0.0/8", "--overlaps=::/0"}, code: 1},
			{kind: "output", args: []string{"10.0.0.0/8", "--overlaps=::/0"}, want: "The networks do not overlap"},
		}},
	{name: "EqualsMappedIPv4",
		steps: []step{
			{kind: "status", args: []string{"::ffff:10.0.0.1", "--equals=10.0.0.1"}, code: 1},
			{kind: "output", args: []string{"::ffff:10.0.0.1", "--equals=10.0.0.1"}, want: "The networks are not equal"},
		}},
	{name: "EqualsBadAddress",
		steps: []step{
			{kind: "status", args: []string{"10.0.0.0/8", "--equals=1.2.3"}, code: 2},
		}},
	{name: "SubnetOfBadPrefixIPv6",
		steps: []step{
			{kind: "status", args: []string{"::/0", "--subnet-of=::/129"}, code: 2},
		}},
	{name: "OverlapsBadNetmask",
		steps: []step{
			{kind: "status", args: []string{"10.0.0.0/8", "--overlaps=10.0.0.0/255.0.255.0"}, code: 2},
		}},
	{name: "OverlapsForcedFamily",
		steps: []step{
			{kind: "status", args: []string{"-4", "10.0.0.0/8", "--overlaps=::/0"}, code: 2},
		}},
	{name: "EqualsBadInputAddress",
		steps: []step{
			{kind: "status", args: []string{"1.2.3", "--equals=1.2.3.0/24"}, code: 2},
		}},
	{name: "Equals-Overlaps",
		steps: []step{
			{kind: "status", args: []string{"10.0.0.0/8", "--equals=10.0.0.0/8", "--overlaps=10.0.0.0/8"}, code: 2},
		}},
	{name: "Equals-Equals",
		steps: []step{
			{kind: "status", args: []string{"10.0.0.0/8", "--equals=10.0.0.0/8", "--equals=10.0.0.0/8"}, code: 2},
		}},
	{name: "Equals-Info",
		steps: []step{
			{kind: "status", args: []string{"--equals=10.0.0.0/8", "-n", "10.0.0.0/8"}, code: 2},
		}},
	{name: "Check-Equals",
		steps: []step{
			{kind: "status", args: []string{"-c", "10.0.0.0/8", "--equals=10.0.0.0/8"}, code: 2},
		}},
	{name: "Equals-IPv4-IPv6",
		steps: []step{
			{kind: "status", args: []string{"-4", "-6", "10.0.0.0/8", "--equals=10.0.0.0/8"}, code: 2},
		}},
	{name: "Overlaps-UnknownOption",
		steps: []step{
			{kind: "status", args: []string{"10.0.0.0/8", "--overlaps=10.0.0.0/8", "--bogus"}, code: 2},
		}},
	{name: "UnknownOption-Overlaps",
		steps: []step{
			{kind: "status", args: []string{"--bogus", "10.0.0.0/8", "--overlaps=10.0.0.0/8"}, code: 2},
		}},
	{name: "EqualsMissingArgument",
		steps: []step{
			{kind: "status", args: []string{"10.0.0.0/8", "--equals"}, code: 2},
		}},
	{name: "SubnetOfNoAddress",
		steps: []step{
			{kind: "status", args: []string{"--subnet-of=10.0.0.0/8"}, code: 2},
		}},
	{name: "UnknownOptionStatus",
		steps: []step{
			{kind: "status", args: []string{"--bogus", "1.2.3.4"}, code: 1},
		}},
	{name: "CidrIPv4",
		steps: []step{
			{kind: "status", args: []string{"--cidr", "192.168.123.7/24"}, code: 0},
			{kind: "output", args: []string{"--cidr", "192.168.123.7/24"}, want: "CIDR=192.168.123.0/24"},
		}},
	{name: "CidrNoDecorateIPv6",
		steps: []step{
			{kind: "status", args: []string{"--cidr", "--no-decorate", "2001:db8::1/64"}, code: 0},
			{kind: "output", args: []string{"--cidr", "--no-decorate", "2001:db8::1/64"}, want: "2001:db8::/64"},
		}},
	{name: "CidrSingleHost",
		steps: []step{
			{kind: "status", args: []string{"--cidr", "192.168.1.1"}, code: 0},
			{kind: "output", args: []string{"--cidr", "192.168.1.1"}, want: "CIDR=192.168.1.1/32"},
		}},
	{name: "CidrSingleHostIPv6",
		steps: []step{
			{kind: "status", args: []string{"--cidr", "::1"}, code: 0},
			{kind: "output", args: []string{"--cidr", "::1"}, want: "CIDR=::1/128"},
		}},
	{name: "CidrFieldOrder",
		steps: []step{
			{kind: "status", args: []string{"-n", "--cidr", "-p", "10.1.2.3/8"}, code: 0},
			{kind: "output", args: []string{"-n", "--cidr", "-p", "10.1.2.3/8"}, want: "NETWORK=10.0.0.0\nPREFIX=8\nCIDR=10.0.0.0/8"},
		}},
	{name: "RandomCidr",
		steps: []step{
			{kind: "success", args: []string{"-r", "24", "--cidr", "--no-decorate"}},
		}},
	{name: "WildcardIPv4",
		steps: []step{
			{kind: "status", args: []string{"--wildcard", "192.168.123.7/24"}, code: 0},
			{kind: "output", args: []string{"--wildcard", "192.168.123.7/24"}, want: "WILDCARD=0.0.0.255"},
		}},
	{name: "WildcardSingleHost",
		steps: []step{
			{kind: "status", args: []string{"--wildcard", "192.168.1.1"}, code: 0},
			{kind: "output", args: []string{"--wildcard", "192.168.1.1"}, want: "WILDCARD=0.0.0.0"},
		}},
	{name: "WildcardPrefix0",
		steps: []step{
			{kind: "status", args: []string{"--wildcard", "0.0.0.0/0"}, code: 0},
			{kind: "output", args: []string{"--wildcard", "0.0.0.0/0"}, want: "WILDCARD=255.255.255.255"},
		}},
	{name: "WildcardPrefix1",
		steps: []step{
			{kind: "status", args: []string{"--wildcard", "10.0.0.1/1"}, code: 0},
			{kind: "output", args: []string{"--wildcard", "10.0.0.1/1"}, want: "WILDCARD=127.255.255.255"},
		}},
	{name: "WildcardPrefix8",
		steps: []step{
			{kind: "status", args: []string{"--wildcard", "10.0.0.1/8"}, code: 0},
			{kind: "output", args: []string{"--wildcard", "10.0.0.1/8"}, want: "WILDCARD=0.255.255.255"},
		}},
	{name: "WildcardPrefix12",
		steps: []step{
			{kind: "status", args: []string{"--wildcard", "172.16.5.4/12"}, code: 0},
			{kind: "output", args: []string{"--wildcard", "172.16.5.4/12"}, want: "WILDCARD=0.15.255.255"},
		}},
	{name: "WildcardPrefix16",
		steps: []step{
			{kind: "status", args: []string{"--wildcard", "172.16.5.4/16"}, code: 0},
			{kind: "output", args: []string{"--wildcard", "172.16.5.4/16"}, want: "WILDCARD=0.0.255.255"},
		}},
	{name: "WildcardPrefix25",
		steps: []step{
			{kind: "status", args: []string{"--wildcard", "192.168.1.130/25"}, code: 0},
			{kind: "output", args: []string{"--wildcard", "192.168.1.130/25"}, want: "WILDCARD=0.0.0.127"},
		}},
	{name: "WildcardPrefix27",
		steps: []step{
			{kind: "status", args: []string{"--wildcard", "192.168.1.130/27"}, code: 0},
			{kind: "output", args: []string{"--wildcard", "192.168.1.130/27"}, want: "WILDCARD=0.0.0.31"},
		}},
	{name: "WildcardPrefix31",
		steps: []step{
			{kind: "status", args: []string{"--wildcard", "192.168.1.130/31"}, code: 0},
			{kind: "output", args: []string{"--wildcard", "192.168.1.130/31"}, want: "WILDCARD=0.0.0.1"},
		}},
	{name: "WildcardNetmask",
		steps: []step{
			{kind: "status", args: []string{"--wildcard", "10.1.2.3/255.255.255.252"}, code: 0},
			{kind: "output", args: []string{"--wildcard", "10.1.2.3/255.255.255.252"}, want: "WILDCARD=0.0.0.3"},
		}},
	{name: "WildcardNoDecorate",
		steps: []step{
			{kind: "status", args: []string{"--wildcard", "--no-decorate", "172.16.5.4/12"}, code: 0},
			{kind: "output", args: []string{"--wildcard", "--no-decorate", "172.16.5.4/12"}, want: "0.15.255.255"},
		}},
	{name: "WildcardIPv6",
		steps: []step{
			{kind: "status", args: []string{"--wildcard", "2001:db8::1/64"}, code: 0},
			{kind: "output", args: []string{"--wildcard", "2001:db8::1/64"}, want: ""},
		}},
	{name: "WildcardPrefix0IPv6",
		steps: []step{
			{kind: "status", args: []string{"--wildcard", "::/0"}, code: 0},
			{kind: "output", args: []string{"--wildcard", "::/0"}, want: ""},
		}},
	{name: "WildcardFieldOrder",
		steps: []step{
			{kind: "status", args: []string{"--wildcard", "-m", "-n", "10.1.2.3/255.255.240.0"}, code: 0},
			{kind: "output", args: []string{"--wildcard", "-m", "-n", "10.1.2.3/255.255.240.0"}, want: "NETWORK=10.1.0.0\nNETMASK=255.255.240.0\nWILDCARD=0.0.15.255"},
		}},
	{name: "InfoWildcard",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "--wildcard", "192.168.2.7/24"}, file: "info-wildcard-192.168.2.7-24"},
		}},
	{name: "JsonWildcard",
		steps: []step{
			{kind: "outfile", args: []string{"-j", "--wildcard", "192.168.2.7/24"}, file: "json-wildcard-192.168.2.7-24"},
		}},
	{name: "InfoWildcardIPv6",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "--wildcard", "2001:db8::/32"}, file: "info-wildcard-2001:db8::-32"},
		}},
	{name: "RandomWildcard",
		steps: []step{
			{kind: "success", args: []string{"-r", "24", "--wildcard", "--no-decorate"}},
		}},
	{name: "FieldOrderNetworkNetmask",
		steps: []step{
			{kind: "status", args: []string{"-m", "-n", "10.1.2.3/8"}, code: 0},
			{kind: "output", args: []string{"-m", "-n", "10.1.2.3/8"}, want: "NETWORK=10.0.0.0\nNETMASK=255.0.0.0"},
		}},
	{name: "AddressAndLookupHost",
		steps: []step{
			{kind: "status", args: []string{"-a", "-o", "localhost", "-4"}, code: 0},
			{kind: "output", args: []string{"-a", "-o", "localhost", "-4"}, want: "ADDRESS=127.0.0.1"},
		}},
	{name: "HostnameHostPrefix",
		steps: []step{
			{kind: "status", args: []string{"-h", "127.0.0.1/32"}, code: 0},
			{kind: "output", args: []string{"-h", "127.0.0.1/32"}, want: "HOSTNAME=localhost"},
		}},
	{name: "InfoReverseDNS",
		steps: []step{
			{kind: "outfile", args: []string{"-i", "--reverse-dns", "192.168.2.7/24"}, file: "info-reverse-dns-192.168.2.7-24"},
		}},
	{name: "JsonReverseDNS",
		steps: []step{
			{kind: "outfile", args: []string{"-j", "--reverse-dns", "192.168.2.7/24"}, file: "json-reverse-dns-192.168.2.7-24"},
		}},
	{name: "HostnameNetwork",
		steps: []step{
			{kind: "status", args: []string{"-h", "127.0.0.1/8"}, code: 1},
			{kind: "output", args: []string{"-h", "127.0.0.1/8"}, want: "ipcalc: -h requires a single address, not a network: 127.0.0.1/8"},
		}},
	{name: "HostnameNetworkJson",
		steps: []step{
			{kind: "status", args: []string{"-j", "-h", "127.0.0.1/8"}, code: 1},
			{kind: "output", args: []string{"-j", "-h", "127.0.0.1/8"}, want: "ipcalc: -h requires a single address, not a network: 127.0.0.1/8"},
		}},
	{name: "HostnameNetworkIPv6",
		steps: []step{
			{kind: "status", args: []string{"-h", "::1/64"}, code: 1},
			{kind: "output", args: []string{"-h", "::1/64"}, want: "ipcalc: -h requires a single address, not a network: ::1/64"},
		}},
	{name: "HostnameNetworkSilent",
		steps: []step{
			{kind: "status", args: []string{"-s", "-h", "127.0.0.1/8"}, code: 1},
			{kind: "output", args: []string{"-s", "-h", "127.0.0.1/8"}, want: ""},
		}},
	{name: "HostnameNetworkFormatJson",
		steps: []step{
			{kind: "status", args: []string{"--format=json", "-h", "127.0.0.1/8"}, code: 1},
			{kind: "output", args: []string{"--format=json", "-h", "127.0.0.1/8"}, want: "ipcalc: -h requires a single address, not a network: 127.0.0.1/8"},
		}},
	{name: "FormatJsonFields",
		steps: []step{
			{kind: "status", args: []string{"--format=json", "-m", "-n", "192.168.2.7/24"}, code: 0},
			{kind: "output", args: []string{"--format=json", "-m", "-n", "192.168.2.7/24"}, want: "{\n  \"NETWORK\":\"192.168.2.0\",\n  \"NETMASK\":\"255.255.255.0\"\n}"},
		}},
	{name: "FormatHumanFields",
		steps: []step{
			{kind: "status", args: []string{"--format=human", "-m", "-n", "192.168.2.7/24"}, code: 0},
			{kind: "output", args: []string{"--format=human", "-m", "-n", "192.168.2.7/24"}, want: "Network:\t192.168.2.0/24\nNetmask:\t255.255.255.0 = 24"},
		}},
	{name: "FormatJsonReverseDNS",
		steps: []step{
			{kind: "status", args: []string{"--format=json", "--reverse-dns", "2001:db8::/31"}, code: 0},
			{kind: "output", args: []string{"--format=json", "--reverse-dns", "2001:db8::/31"}, want: "{\n  \"REVERSEDNS\":[\n    \"8.b.d.0.1.0.0.2.ip6.arpa.\",\n    \"9.b.d.0.1.0.0.2.ip6.arpa.\"]\n}"},
		}},
	{name: "FormatJsonHostname",
		steps: []step{
			{kind: "status", args: []string{"--format=json", "-h", "127.0.0.1"}, code: 0},
			{kind: "output", args: []string{"--format=json", "-h", "127.0.0.1"}, want: "{\n  \"HOSTNAME\":\"localhost\"\n}"},
		}},
	{name: "FormatJsonLookupHost",
		steps: []step{
			{kind: "status", args: []string{"--format=json", "-4", "-o", "localhost"}, code: 0},
			{kind: "output", args: []string{"--format=json", "-4", "-o", "localhost"}, want: "{\n  \"ADDRESS\":\"127.0.0.1\"\n}"},
		}},
	{name: "FormatJsonWithJ",
		steps: []step{
			{kind: "status", args: []string{"-j", "--format=json", "-n", "10.1.2.3/8"}, code: 0},
			{kind: "output", args: []string{"-j", "--format=json", "-n", "10.1.2.3/8"}, want: "{\n  \"NETWORK\":\"10.0.0.0\"\n}"},
		}},
	{name: "FormatValueWithNoDecorate",
		steps: []step{
			{kind: "status", args: []string{"--no-decorate", "--format=value", "-n", "10.1.2.3/8"}, code: 0},
			{kind: "output", args: []string{"--no-decorate", "--format=value", "-n", "10.1.2.3/8"}, want: "10.0.0.0"},
		}},
	{name: "FormatHumanEquals",
		steps: []step{
			{kind: "status", args: []string{"--format=human", "192.168.7.15/24", "--equals=192.168.7.5/24"}, code: 0},
			{kind: "output", args: []string{"--format=human", "192.168.7.15/24", "--equals=192.168.7.5/24"}, want: "The networks are equal"},
		}},
	{name: "FormatCheck",
		steps: []step{
			{kind: "status", args: []string{"-c", "--format=shell", "127.0.0.1"}, code: 0},
			{kind: "output", args: []string{"-c", "--format=shell", "127.0.0.1"}, want: ""},
		}},
	{name: "FormatHumanSummary",
		steps: []step{
			{kind: "outfile", args: []string{"--format=human", "192.168.2.0/24"}, file: "standard-192.168.2.0-24"},
		}},
	{name: "FormatJsonAllInfo",
		steps: []step{
			{kind: "outfile", args: []string{"--format=json", "--all-info", "192.168.2.7/24"}, file: "json-all-info-192.168.2.7-24"},
		}},
	{name: "FormatShellSummary",
		steps: []step{
			{kind: "outfile", args: []string{"--format=shell", "192.168.2.7/24"}, file: "shell-192.168.2.7-24"},
		}},
	{name: "FormatJsonInfoLookupHost",
		steps: []step{
			{kind: "outfile", args: []string{"--format=json", "-i", "-4", "-o", "localhost"}, file: "hostname-localhost-ipv4-json"},
		}},
	{name: "FormatValueSplit",
		steps: []step{
			{kind: "outfile", args: []string{"--format=value", "-S", "26", "192.168.5.45/24"}, file: "nsplit-192.168.5.45-24-26"},
		}},
	{name: "FormatJsonSplit",
		steps: []step{
			{kind: "outfile", args: []string{"--format=json", "-S", "26", "192.168.5.45/24"}, file: "json-split-192.168.5.45-24-26"},
		}},
	{name: "FormatHumanSplit",
		steps: []step{
			{kind: "outfile", args: []string{"--format=human", "-S", "26", "192.168.5.45/24"}, file: "split-192.168.5.45-24-26"},
		}},
	{name: "FormatValueDeaggregate",
		steps: []step{
			{kind: "outfile", args: []string{"--format=value", "-d", "192.168.2.7-192.168.2.13"}, file: "deaggregate-192.168.2.7-192.168.2.13"},
		}},
	{name: "FormatShellFields",
		steps: []step{
			{kind: "equal", args: []string{"--format=shell", "-m", "-n", "10.1.2.3/8"}, args2: []string{"-m", "-n", "10.1.2.3/8"}},
		}},
	{name: "FormatValueFields",
		steps: []step{
			{kind: "equal", args: []string{"--format=value", "-m", "-n", "10.1.2.3/8"}, args2: []string{"--no-decorate", "-m", "-n", "10.1.2.3/8"}},
		}},
	{name: "FormatUnknown",
		steps: []step{
			{kind: "status", args: []string{"--format=xml", "10.0.0.1"}, code: 1},
			{kind: "output", args: []string{"--format=xml", "10.0.0.1"}, want: "ipcalc: unknown output format: xml; use human, shell, json or value"},
		}},
	{name: "FormatUnknownSilent",
		steps: []step{
			{kind: "status", args: []string{"-s", "--format=xml", "10.0.0.1"}, code: 1},
			{kind: "output", args: []string{"-s", "--format=xml", "10.0.0.1"}, want: ""},
		}},
	{name: "FormatTwice",
		steps: []step{
			{kind: "status", args: []string{"--format=json", "--format=shell", "10.0.0.1"}, code: 1},
			{kind: "output", args: []string{"--format=json", "--format=shell", "10.0.0.1"}, want: "ipcalc: conflicting output formats: --format=json and --format=shell"},
		}},
	{name: "FormatConflictJson",
		steps: []step{
			{kind: "status", args: []string{"-j", "--format=shell", "10.0.0.1"}, code: 1},
			{kind: "output", args: []string{"-j", "--format=shell", "10.0.0.1"}, want: "ipcalc: conflicting output formats: -j and --format=shell"},
		}},
	{name: "FormatConflictNoDecorate",
		steps: []step{
			{kind: "status", args: []string{"--no-decorate", "--format=json", "-m", "10.0.0.1"}, code: 1},
			{kind: "output", args: []string{"--no-decorate", "--format=json", "-m", "10.0.0.1"}, want: "ipcalc: conflicting output formats: --no-decorate and --format=json"},
		}},
	{name: "FormatInfoWithField",
		steps: []step{
			{kind: "status", args: []string{"--format=json", "-i", "-m", "10.0.0.1"}, code: 1},
			{kind: "output", args: []string{"--format=json", "-i", "-m", "10.0.0.1"}, want: "ipcalc: -i and --all-info cannot be combined with specific field options"},
		}},
	{name: "FormatAllInfoWithField",
		steps: []step{
			{kind: "status", args: []string{"--format=human", "--all-info", "-n", "10.0.0.1/24"}, code: 1},
			{kind: "output", args: []string{"--format=human", "--all-info", "-n", "10.0.0.1/24"}, want: "ipcalc: -i and --all-info cannot be combined with specific field options"},
		}},
	{name: "FormatValueSummary",
		steps: []step{
			{kind: "status", args: []string{"--format=value", "10.0.0.1/24"}, code: 1},
			{kind: "output", args: []string{"--format=value", "10.0.0.1/24"}, want: "ipcalc: --format=value needs specific field options; use --format=shell or --format=json"},
		}},
	{name: "FormatShellSplit",
		steps: []step{
			{kind: "status", args: []string{"--format=shell", "-S", "26", "10.0.0.0/24"}, code: 1},
			{kind: "output", args: []string{"--format=shell", "-S", "26", "10.0.0.0/24"}, want: "ipcalc: --format=shell is not supported with --split, --split-hosts or --deaggregate; use --format=value"},
		}},
	{name: "FormatShellDeaggregate",
		steps: []step{
			{kind: "status", args: []string{"--format=shell", "-d", "10.0.0.1-10.0.0.9"}, code: 1},
			{kind: "output", args: []string{"--format=shell", "-d", "10.0.0.1-10.0.0.9"}, want: "ipcalc: --format=shell is not supported with --split, --split-hosts or --deaggregate; use --format=value"},
		}},
	{name: "FormatJsonEquals",
		steps: []step{
			{kind: "status", args: []string{"--format=json", "10.0.0.0/8", "--equals=10.0.0.0/8"}, code: 2},
			{kind: "output", args: []string{"--format=json", "10.0.0.0/8", "--equals=10.0.0.0/8"}, want: "ipcalc: --format=json is not supported with --equals, --subnet-of or --overlaps"},
		}},
	{name: "FormatValueSubnetOf",
		steps: []step{
			{kind: "status", args: []string{"--format=value", "10.0.0.0/8", "--subnet-of=10.0.0.0/7"}, code: 2},
			{kind: "output", args: []string{"--format=value", "10.0.0.0/8", "--subnet-of=10.0.0.0/7"}, want: "ipcalc: --format=value is not supported with --equals, --subnet-of or --overlaps"},
		}},
	{name: "FormatShellOverlaps",
		steps: []step{
			{kind: "status", args: []string{"--format=shell", "10.0.0.0/8", "--overlaps=10.0.0.0/7"}, code: 2},
			{kind: "output", args: []string{"--format=shell", "10.0.0.0/8", "--overlaps=10.0.0.0/7"}, want: "ipcalc: --format=shell is not supported with --equals, --subnet-of or --overlaps"},
		}},
	{name: "BitsAddress",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "192.168.92.6"}, file: "bits-192.168.92.6"},
		}},
	{name: "BitsHostBits",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "192.168.92.6/22"}, file: "bits-192.168.92.6-22"},
		}},
	{name: "BitsPrefix8",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "10.0.0.0/8"}, file: "bits-10.0.0.0-8"},
		}},
	{name: "BitsPrefix31",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "192.168.1.4/31"}, file: "bits-192.168.1.4-31"},
		}},
	{name: "BitsPrefix1",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "128.0.0.0/1"}, file: "bits-128.0.0.0-1"},
		}},
	{name: "BitsPrefix0",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "0.0.0.0/0"}, file: "bits-0.0.0.0-0"},
		}},
	{name: "BitsDetails",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "192.168.92.6/24"}, file: "bits-192.168.92.6-24"},
		}},
	{name: "BitsAddressHostMin",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "192.168.92.1/24"}, file: "bits-192.168.92.1-24"},
		}},
	{name: "BitsAddressBroadcast",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "192.168.92.255/24"}, file: "bits-192.168.92.255-24"},
		}},
	{name: "BitsAddressPrefix31",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "192.168.1.5/31"}, file: "bits-192.168.1.5-31"},
		}},
	{name: "BitsAddressIPv6",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "2001:db8::1"}, file: "bits-2001:db8::1"},
		}},
	{name: "BitsHostBitsIPv6",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "2001:db8:abcd:12::1/56"}, file: "bits-2001:db8:abcd:12::1-56"},
		}},
	{name: "BitsPrefix64IPv6",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "2001:db8::/64"}, file: "bits-2001:db8::-64"},
		}},
	{name: "BitsPrefix62IPv6",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "2001:db8:abcd:10::/62"}, file: "bits-2001:db8:abcd:10::-62"},
		}},
	{name: "BitsPrefix96IPv6",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "2001:db8::/96"}, file: "bits-2001:db8::-96"},
		}},
	{name: "BitsPrefix65IPv6",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "2001:db8::/65"}, file: "bits-2001:db8::-65"},
		}},
	{name: "BitsPrefix0IPv6",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "::/0"}, file: "bits-::-0"},
		}},
	{name: "BitsPrefix127IPv6",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "2001:db8::/127"}, file: "bits-2001:db8::-127"},
		}},
	{name: "BitsUniqueLocalIPv6",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "fd10:da05:b5ac::/48"}, file: "bits-fd10:da05:b5ac::-48"},
		}},
	{name: "BitsLinkLocalIPv6",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "fe80::/64"}, file: "bits-fe80::-64"},
		}},
	{name: "BitsLinkLocalPrefix10IPv6",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "fe80::/10"}, file: "bits-fe80::-10"},
		}},
	{name: "BitsMulticastIPv6",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "ff02::/16"}, file: "bits-ff02::-16"},
		}},
	{name: "BitsFormatHuman",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "--format=human", "192.168.92.6/22"}, file: "bits-192.168.92.6-22"},
		}},
	{name: "BitsJson",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "--format=json", "192.168.92.0/22"}, file: "bits-json-192.168.92.0-22"},
		}},
	{name: "BitsJsonIPv6",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "--format=json", "2001:db8::1"}, file: "bits-json-2001:db8::1"},
		}},
	{name: "BitsJsonNetworkIPv6",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "--format=json", "2001:db8:abcd:12::1/56"}, file: "bits-json-2001:db8:abcd:12::1-56"},
		}},
	{name: "BitsShell",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "--format=shell", "192.168.92.6/22"}, file: "bits-shell-192.168.92.6-22"},
		}},
	{name: "BitsShellPrefix8",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "--format=shell", "10.0.0.0/8"}, file: "bits-shell-10.0.0.0-8"},
		}},
	{name: "BitsShellPrefix31",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "--format=shell", "192.168.1.5/31"}, file: "bits-shell-192.168.1.5-31"},
		}},
	{name: "BitsShellAddress",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "--format=shell", "192.168.92.6"}, file: "bits-shell-192.168.92.6"},
		}},
	{name: "BitsShellIPv6",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "--format=shell", "2001:db8:abcd:12::1/56"}, file: "bits-shell-2001:db8:abcd:12::1-56"},
		}},
	{name: "BitsShellAddressIPv6",
		env: map[string]string{"COLUMNS": "80"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "--format=shell", "2001:db8::1"}, file: "bits-shell-2001:db8::1"},
		}},
	{name: "BitsWideAddressIPv6",
		env: map[string]string{"COLUMNS": "135"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "2001:db8::1"}, file: "bits-wide-2001:db8::1"},
		}},
	{name: "BitsWideHostBitsIPv6",
		env: map[string]string{"COLUMNS": "135"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "2001:db8:abcd:12::1/56"}, file: "bits-wide-2001:db8:abcd:12::1-56"},
		}},
	{name: "BitsWidePrefix96IPv6",
		env: map[string]string{"COLUMNS": "200"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "2001:db8::/96"}, file: "bits-wide-2001:db8::-96"},
		}},
	{name: "BitsWidePrefix0IPv6",
		env: map[string]string{"COLUMNS": "135"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "::/0"}, file: "bits-wide-::-0"},
		}},
	{name: "BitsNarrowIPv6",
		env: map[string]string{"COLUMNS": "134"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "2001:db8:abcd:12::1/56"}, file: "bits-2001:db8:abcd:12::1-56"},
		}},
	{name: "BitsColumnsInvalid",
		env: map[string]string{"COLUMNS": "abc"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "2001:db8:abcd:12::1/56"}, file: "bits-2001:db8:abcd:12::1-56"},
		}},
	{name: "BitsColumnsZero",
		env: map[string]string{"COLUMNS": "0"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "2001:db8:abcd:12::1/56"}, file: "bits-2001:db8:abcd:12::1-56"},
		}},
	{name: "BitsWideIPv4",
		env: map[string]string{"COLUMNS": "200"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "192.168.92.6/22"}, file: "bits-192.168.92.6-22"},
		}},
	{name: "BitsShellPrefix0IPv6",
		env: map[string]string{"COLUMNS": "200"},
		steps: []step{
			{kind: "outfile", args: []string{"--binary", "--format=shell", "::/0"}, file: "bits-shell-::-0"},
		}},
	{name: "BitsAddressValue",
		env: map[string]string{"COLUMNS": "200"},
		steps: []step{
			{kind: "status", args: []string{"--binary", "--format=value", "192.168.92.6"}, code: 0},
			{kind: "output", args: []string{"--binary", "--format=value", "192.168.92.6"}, want: "11000000101010000101110000000110"},
		}},
	{name: "BitsNetworkValue",
		env: map[string]string{"COLUMNS": "200"},
		steps: []step{
			{kind: "status", args: []string{"--binary", "--format=value", "192.168.92.6/22"}, code: 0},
			{kind: "output", args: []string{"--binary", "--format=value", "192.168.92.6/22"}, want: "11000000101010000101110000000000"},
		}},
	{name: "BitsNetmaskValue",
		env: map[string]string{"COLUMNS": "200"},
		steps: []step{
			{kind: "status", args: []string{"--binary", "--format=value", "192.168.92.6", "255.255.252.0"}, code: 0},
			{kind: "output", args: []string{"--binary", "--format=value", "192.168.92.6", "255.255.252.0"}, want: "11000000101010000101110000000000"},
		}},
	{name: "BitsClassPrefixValue",
		env: map[string]string{"COLUMNS": "200"},
		steps: []step{
			{kind: "status", args: []string{"--binary", "--format=value", "--class-prefix", "10.1.2.3"}, code: 0},
			{kind: "output", args: []string{"--binary", "--format=value", "--class-prefix", "10.1.2.3"}, want: "00001010000000000000000000000000"},
		}},
	{name: "BitsPrefix0Value",
		env: map[string]string{"COLUMNS": "200"},
		steps: []step{
			{kind: "status", args: []string{"--binary", "--format=value", "0.0.0.0/0"}, code: 0},
			{kind: "output", args: []string{"--binary", "--format=value", "0.0.0.0/0"}, want: "00000000000000000000000000000000"},
		}},
	{name: "BitsNetworkIPv6Value",
		env: map[string]string{"COLUMNS": "200"},
		steps: []step{
			{kind: "status", args: []string{"--binary", "--format=value", "2001:db8::/32"}, code: 0},
			{kind: "output", args: []string{"--binary", "--format=value", "2001:db8::/32"}, want: "00100000000000010000110110111000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		}},
	{name: "BitsAddressIPv6Value",
		env: map[string]string{"COLUMNS": "200"},
		steps: []step{
			{kind: "status", args: []string{"--binary", "--format=value", "2001:db8::1"}, code: 0},
			{kind: "output", args: []string{"--binary", "--format=value", "2001:db8::1"}, want: "00100000000000010000110110111000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001"},
		}},
	{name: "BitsPrefix0IPv6Value",
		env: map[string]string{"COLUMNS": "200"},
		steps: []step{
			{kind: "status", args: []string{"--binary", "--format=value", "::/0"}, code: 0},
			{kind: "output", args: []string{"--binary", "--format=value", "::/0"}, want: "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		}},
	{name: "BitsPrefix32Value",
		steps: []step{
			{kind: "equal", args: []string{"--binary", "--format=value", "192.168.92.6/32"}, args2: []string{"--binary", "--format=value", "192.168.92.6"}},
		}},
	{name: "BitsPrefix32",
		steps: []step{
			{kind: "equal", args: []string{"--binary", "192.168.92.6/32"}, args2: []string{"--binary", "192.168.92.6"}},
		}},
	{name: "BitsJsonLegacy",
		steps: []step{
			{kind: "equal", args: []string{"--binary", "-j", "192.168.92.0/22"}, args2: []string{"--binary", "--format=json", "192.168.92.0/22"}},
		}},
	{name: "BitsNoDecorate",
		steps: []step{
			{kind: "status", args: []string{"--binary", "--no-decorate", "192.168.92.6/22"}, code: 1},
			{kind: "output", args: []string{"--binary", "--no-decorate", "192.168.92.6/22"}, want: "ipcalc: --no-decorate is not supported with --binary; use --format=value"},
		}},
	{name: "BitsJsonNoDecorate",
		steps: []step{
			{kind: "status", args: []string{"--binary", "-j", "--no-decorate", "192.168.92.0/22"}, code: 1},
			{kind: "output", args: []string{"--binary", "-j", "--no-decorate", "192.168.92.0/22"}, want: "ipcalc: --no-decorate is not supported with --binary; use --format=value"},
		}},
	{name: "BitsNoDecorateFormatValue",
		steps: []step{
			{kind: "status", args: []string{"--binary", "--no-decorate", "--format=value", "192.168.92.6/22"}, code: 1},
			{kind: "output", args: []string{"--binary", "--no-decorate", "--format=value", "192.168.92.6/22"}, want: "ipcalc: --no-decorate is not supported with --binary; use --format=value"},
		}},
	{name: "BitsNoDecorateSilent",
		steps: []step{
			{kind: "status", args: []string{"--binary", "-s", "--no-decorate", "192.168.92.6/22"}, code: 1},
			{kind: "output", args: []string{"--binary", "-s", "--no-decorate", "192.168.92.6/22"}, want: ""},
		}},
	{name: "BitsBadAddress",
		steps: []step{
			{kind: "status", args: []string{"--binary", "1.2.3.400"}, code: 1},
		}},
	{name: "BitsBadPrefixIPv4",
		steps: []step{
			{kind: "status", args: []string{"--binary", "1.2.3.4/33"}, code: 1},
		}},
	{name: "BitsBadPrefixIPv6",
		steps: []step{
			{kind: "status", args: []string{"--binary", "2001:db8::/129"}, code: 1},
		}},
	{name: "BitsForcedFamily",
		steps: []step{
			{kind: "status", args: []string{"--binary", "-6", "1.2.3.4"}, code: 1},
		}},
	{name: "Bits-Info",
		steps: []step{
			{kind: "status", args: []string{"--binary", "-i", "10.0.0.0/8"}, code: 1},
		}},
	{name: "Bits-AllInfo",
		steps: []step{
			{kind: "status", args: []string{"--binary", "--all-info", "10.0.0.0/8"}, code: 1},
		}},
	{name: "Bits-Network",
		steps: []step{
			{kind: "status", args: []string{"--binary", "-n", "10.0.0.0/8"}, code: 1},
		}},
	{name: "Bits-Cidr",
		steps: []step{
			{kind: "status", args: []string{"--binary", "--cidr", "10.0.0.0/8"}, code: 1},
		}},
	{name: "Bits-Check",
		steps: []step{
			{kind: "status", args: []string{"--binary", "-c", "10.0.0.0/8"}, code: 1},
		}},
	{name: "Bits-Split",
		steps: []step{
			{kind: "status", args: []string{"--binary", "-S", "24", "10.0.0.0/8"}, code: 1},
		}},
	{name: "Bits-SplitHosts",
		steps: []step{
			{kind: "status", args: []string{"--binary", "--split-hosts=10", "10.0.0.0/8"}, code: 1},
		}},
	{name: "Bits-Deaggregate",
		steps: []step{
			{kind: "status", args: []string{"--binary", "-d", "10.0.0.0-10.0.0.255"}, code: 1},
		}},
	{name: "Bits-Random",
		steps: []step{
			{kind: "status", args: []string{"--binary", "-r", "24"}, code: 1},
		}},
	{name: "Bits-Equals",
		steps: []step{
			{kind: "status", args: []string{"--binary", "10.0.0.0/8", "--equals=10.0.0.0/8"}, code: 2},
		}},
	{name: "BitsFormatConflict",
		steps: []step{
			{kind: "status", args: []string{"--binary", "-j", "--format=shell", "10.0.0.0/8"}, code: 1},
		}},
}
