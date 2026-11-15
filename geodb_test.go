package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"net/netip"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"testing"
)

// A small MaxMind DB writer (https://maxmind.github.io/MaxMind-DB/), so that
// the geo tests do not depend on the databases installed on the system. It
// builds a database from the JSON files in testdata:
//
//	{"database_type": "...", "networks": {"192.0.2.0/24": {...}, ...}}
//
// Integers are written as the smallest unsigned type that holds them, as
// MaxMind does (e.g. accuracy_radius as uint16).

const (
	mmdbTypeUTF8   = 2
	mmdbTypeDouble = 3
	mmdbTypeUint16 = 5
	mmdbTypeUint32 = 6
	mmdbTypeMap    = 7
	mmdbTypeUint64 = 9
	mmdbTypeArray  = 11
)

type mmdbUint struct {
	typ   int
	value uint64
}

func mmdbControl(typ, size int) []byte {
	var out, ext []byte
	switch {
	case size < 29:
		out = []byte{byte(size)}
	case size < 29+256:
		out, ext = []byte{29}, []byte{byte(size - 29)}
	case size < 285+65536:
		out, ext = []byte{30}, binary.BigEndian.AppendUint16(nil, uint16(size-285))
	default:
		v := size - 65821
		out, ext = []byte{31}, []byte{byte(v >> 16), byte(v >> 8), byte(v)}
	}
	if typ <= 7 {
		return append([]byte{byte(typ<<5) | out[0]}, ext...)
	}
	return append([]byte{out[0], byte(typ - 7)}, ext...)
}

func mmdbEncodeUint(typ int, v uint64) []byte {
	var data []byte
	for ; v > 0; v >>= 8 {
		data = append([]byte{byte(v)}, data...)
	}
	return append(mmdbControl(typ, len(data)), data...)
}

func mmdbEncode(v any) []byte {
	switch x := v.(type) {
	case mmdbUint:
		return mmdbEncodeUint(x.typ, x.value)
	case json.Number:
		if s := x.String(); !containsAny(s, ".eE") {
			n, err := x.Int64()
			if err != nil || n < 0 {
				panic("bad integer " + s)
			}
			switch {
			case n < 1<<16:
				return mmdbEncodeUint(mmdbTypeUint16, uint64(n))
			case n < 1<<32:
				return mmdbEncodeUint(mmdbTypeUint32, uint64(n))
			}
			return mmdbEncodeUint(mmdbTypeUint64, uint64(n))
		}
		f, _ := x.Float64()
		return binary.BigEndian.AppendUint64(mmdbControl(mmdbTypeDouble, 8), math.Float64bits(f))
	case string:
		return append(mmdbControl(mmdbTypeUTF8, len(x)), x...)
	case map[string]any:
		keys := make([]string, 0, len(x))
		for k := range x {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		out := mmdbControl(mmdbTypeMap, len(x))
		for _, k := range keys {
			out = append(out, mmdbEncode(k)...)
			out = append(out, mmdbEncode(x[k])...)
		}
		return out
	case []any:
		out := mmdbControl(mmdbTypeArray, len(x))
		for _, e := range x {
			out = append(out, mmdbEncode(e)...)
		}
		return out
	}
	panic(fmt.Sprintf("cannot encode %T", v))
}

type mmdbNode struct {
	child [2]any // nil, *mmdbNode or mmdbData
}

type mmdbData int

func containsAny(s, chars string) bool {
	for _, c := range s {
		for _, d := range chars {
			if c == d {
				return true
			}
		}
	}
	return false
}

// mmdbInsert adds a leaf for the network; IPv4 networks go under ::/96,
// where libmaxminddb looks them up in an IPv6 tree.
func mmdbInsert(root *mmdbNode, network string, data mmdbData) error {
	p, err := netip.ParsePrefix(network)
	if err != nil {
		return err
	}
	a := p.Addr()
	plen := p.Bits()
	if a.Is4() {
		plen += 96
		b := a.As4()
		a = netip.AddrFrom16([16]byte{12: b[0], 13: b[1], 14: b[2], 15: b[3]})
	}
	bytes16 := a.As16()
	node := root
	for depth := 0; depth < plen; depth++ {
		bit := (bytes16[depth/8] >> (7 - uint(depth%8))) & 1
		if depth == plen-1 {
			if node.child[bit] != nil {
				return fmt.Errorf("%s overlaps another network", network)
			}
			node.child[bit] = data
			break
		}
		switch c := node.child[bit].(type) {
		case nil:
			n := &mmdbNode{}
			node.child[bit] = n
			node = n
		case *mmdbNode:
			node = c
		default:
			return fmt.Errorf("%s overlaps another network", network)
		}
	}
	return nil
}

func buildMMDB(desc []byte) ([]byte, error) {
	var d struct {
		DatabaseType string                     `json:"database_type"`
		Networks     map[string]json.RawMessage `json:"networks"`
	}
	if err := json.Unmarshal(desc, &d); err != nil {
		return nil, err
	}
	networks := make([]string, 0, len(d.Networks))
	for n := range d.Networks {
		networks = append(networks, n)
	}
	sort.Strings(networks)

	root := &mmdbNode{}
	var dataSection []byte
	for _, n := range networks {
		if err := mmdbInsert(root, n, mmdbData(len(dataSection))); err != nil {
			return nil, err
		}
		dec := json.NewDecoder(bytes.NewReader(d.Networks[n]))
		dec.UseNumber()
		var rec any
		if err := dec.Decode(&rec); err != nil {
			return nil, err
		}
		dataSection = append(dataSection, mmdbEncode(rec)...)
	}

	nodes := []*mmdbNode{root}
	for i := 0; i < len(nodes); i++ {
		for _, c := range nodes[i].child {
			if n, ok := c.(*mmdbNode); ok {
				nodes = append(nodes, n)
			}
		}
	}
	index := map[*mmdbNode]int{}
	for i, n := range nodes {
		index[n] = i
	}
	count := len(nodes)
	record := func(c any) uint32 {
		switch x := c.(type) {
		case nil:
			return uint32(count)
		case mmdbData:
			return uint32(count + 16 + int(x))
		}
		return uint32(index[c.(*mmdbNode)])
	}
	var tree []byte
	for _, n := range nodes {
		for _, c := range n.child {
			r := record(c)
			tree = append(tree, byte(r>>16), byte(r>>8), byte(r))
		}
	}

	meta := mmdbEncode(map[string]any{
		"node_count":                  mmdbUint{mmdbTypeUint32, uint64(count)},
		"record_size":                 mmdbUint{mmdbTypeUint16, 24},
		"ip_version":                  mmdbUint{mmdbTypeUint16, 6},
		"database_type":               d.DatabaseType,
		"languages":                   []any{"en"},
		"binary_format_major_version": mmdbUint{mmdbTypeUint16, 2},
		"binary_format_minor_version": mmdbUint{mmdbTypeUint16, 0},
		"build_epoch":                 mmdbUint{mmdbTypeUint64, 1},
		"description":                 map[string]any{"en": "ipcalc test database"},
	})

	var out []byte
	out = append(out, tree...)
	out = append(out, make([]byte, 16)...)
	out = append(out, dataSection...)
	out = append(out, "\xab\xcd\xefMaxMind.com"...)
	out = append(out, meta...)
	return out, nil
}

var (
	geoOnce sync.Once
	geoDir  string
	geoErr  error
)

// geoDatabases writes the test databases once and returns their directory.
func geoDatabases(t *testing.T) string {
	t.Helper()
	geoOnce.Do(func() {
		geoDir, geoErr = os.MkdirTemp("", "ipcalc-geo")
		if geoErr != nil {
			return
		}
		for _, db := range []string{"country", "city", "asn"} {
			var desc []byte
			if desc, geoErr = os.ReadFile(filepath.Join("testdata", "geo-"+db+".json")); geoErr != nil {
				return
			}
			var mmdb []byte
			if mmdb, geoErr = buildMMDB(desc); geoErr != nil {
				return
			}
			if geoErr = os.WriteFile(filepath.Join(geoDir, "geo-"+db+".mmdb"), mmdb, 0o644); geoErr != nil {
				return
			}
		}
	})
	if geoErr != nil {
		t.Fatal(geoErr)
	}
	return geoDir
}
