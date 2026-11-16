// Code generated from the output of the C ipcalc on a terminal; DO NOT EDIT.

package ipcalc

var colorCases = []struct {
	name, bin string
	args      []string
	env       map[string]string
}{
	{"summary4", "", []string{"192.168.2.7/24"}, nil},
	{"summary4-single", "", []string{"10.0.0.1"}, nil},
	{"summary6", "", []string{"2001:db8::1/64"}, nil},
	{"allinfo4", "", []string{"--all-info", "192.168.2.7/24"}, nil},
	{"allinfo6", "", []string{"--all-info", "fe80::21b:21ff:fe3a:5c7d/64"}, nil},
	{"allinfo6-single", "", []string{"--all-info", "2001:db8::300:ff:fe00:1"}, nil},
	{"fields", "", []string{"-n", "-m", "-b", "192.168.2.7/24"}, nil},
	{"reverse", "", []string{"--reverse-dns", "--addrspace", "2001:db8::/31"}, nil},
	{"geo", "geo", []string{"--all-info", "192.0.2.5"}, nil},
	{"geo-v6", "geo", []string{"-g", "2001:db8::1"}, nil},
	{"binary4", "", []string{"--binary", "192.168.92.6/22"}, map[string]string{"COLUMNS": "80"}},
	{"binary4-single", "", []string{"--binary", "192.168.92.6"}, map[string]string{"COLUMNS": "80"}},
	{"binary6", "", []string{"--binary", "2001:db8:abcd:12::1/56"}, map[string]string{"COLUMNS": "80"}},
	{"binary6-wide", "", []string{"--binary", "2001:db8:abcd:12::1/56"}, map[string]string{"COLUMNS": "200"}},
	{"binary6-linklocal", "", []string{"--binary", "fe80::1/10"}, map[string]string{"COLUMNS": "80"}},
	{"no-decorate", "", []string{"--no-decorate", "192.168.2.7/24"}, nil},
	{"no-color", "", []string{"192.168.2.7/24"}, map[string]string{"NO_COLOR": "1"}},
	{"no-color-binary", "", []string{"--binary", "192.168.92.6/22"}, map[string]string{"COLUMNS": "80", "NO_COLOR": "1"}},
	{"json", "", []string{"-j", "192.168.2.7/24"}, nil},
	{"format-shell", "", []string{"--format=shell", "-n", "192.168.2.7/24"}, nil},
	{"equals", "", []string{"--equals", "10.0.0.0/8", "10.0.0.0/8"}, nil},
	{"split", "", []string{"-S", "26", "10.0.0.0/24"}, nil},
	{"deaggregate", "", []string{"-d", "10.0.0.1-10.0.0.9"}, nil},
}
