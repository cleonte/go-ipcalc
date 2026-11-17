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

package u128

import (
	"math/big"
	"testing"
)

var (
	zero = Uint{}
	one  = From64(1)
	max  = Uint{^uint64(0), ^uint64(0)}
	// bit 64 is the lowest bit of Hi.
	two64 = Uint{Hi: 1}
)

func toBig(a Uint) *big.Int {
	v := new(big.Int).SetUint64(a.Hi)
	v.Lsh(v, 64)
	return v.Or(v, new(big.Int).SetUint64(a.Lo))
}

var mod128 = new(big.Int).Lsh(big.NewInt(1), 128)

func fromBig(v *big.Int) Uint {
	v = new(big.Int).Mod(v, mod128)
	lo := new(big.Int).And(v, new(big.Int).SetUint64(^uint64(0)))
	hi := new(big.Int).Rsh(v, 64)
	return Uint{hi.Uint64(), lo.Uint64()}
}

var samples = []Uint{
	zero, one, From64(2), From64(^uint64(0)), two64,
	{Hi: 1, Lo: 1}, {Hi: 1 << 63}, {Hi: ^uint64(0)}, max,
	{Hi: 0x20010db800000000, Lo: 0x0000000000000001},
	{Hi: 0x0123456789abcdef, Lo: 0xfedcba9876543210},
}

func TestFrom64(t *testing.T) {
	if got := From64(7); got != (Uint{0, 7}) {
		t.Errorf("From64(7) = %+v", got)
	}
}

func TestAddSub(t *testing.T) {
	tests := []struct{ a, b, sum Uint }{
		{zero, zero, zero},
		{zero, one, one},
		{From64(^uint64(0)), one, two64},
		{max, one, zero},
		{max, max, Uint{^uint64(0), ^uint64(0) - 1}},
		{two64, max, Uint{0, ^uint64(0)}},
	}
	for _, tt := range tests {
		if got := tt.a.Add(tt.b); got != tt.sum {
			t.Errorf("%+v + %+v = %+v, want %+v", tt.a, tt.b, got, tt.sum)
		}
		if got := tt.sum.Sub(tt.b); got != tt.a {
			t.Errorf("%+v - %+v = %+v, want %+v", tt.sum, tt.b, got, tt.a)
		}
	}
	if got := zero.Sub(one); got != max {
		t.Errorf("0 - 1 = %+v, want max", got)
	}
	if got := two64.Sub(one); got != From64(^uint64(0)) {
		t.Errorf("2^64 - 1 = %+v", got)
	}
}

func TestAgainstBig(t *testing.T) {
	for _, a := range samples {
		for _, b := range samples {
			ba, bb := toBig(a), toBig(b)
			if got, want := a.Add(b), fromBig(new(big.Int).Add(ba, bb)); got != want {
				t.Errorf("%+v + %+v = %+v, want %+v", a, b, got, want)
			}
			if got, want := a.Sub(b), fromBig(new(big.Int).Sub(ba, bb)); got != want {
				t.Errorf("%+v - %+v = %+v, want %+v", a, b, got, want)
			}
			if got, want := a.And(b), fromBig(new(big.Int).And(ba, bb)); got != want {
				t.Errorf("%+v & %+v = %+v, want %+v", a, b, got, want)
			}
			if got, want := a.Or(b), fromBig(new(big.Int).Or(ba, bb)); got != want {
				t.Errorf("%+v | %+v = %+v, want %+v", a, b, got, want)
			}
			if got, want := a.Cmp(b), ba.Cmp(bb); got != want {
				t.Errorf("Cmp(%+v, %+v) = %d, want %d", a, b, got, want)
			}
		}
	}
}

func TestCmp(t *testing.T) {
	tests := []struct {
		a, b Uint
		want int
	}{
		{zero, zero, 0},
		{max, max, 0},
		{zero, one, -1},
		{one, zero, 1},
		{From64(^uint64(0)), two64, -1},
		{two64, From64(^uint64(0)), 1},
		{Uint{1, 0}, Uint{0, ^uint64(0)}, 1},
		{Uint{1, 5}, Uint{1, 6}, -1},
	}
	for _, tt := range tests {
		if got := tt.a.Cmp(tt.b); got != tt.want {
			t.Errorf("Cmp(%+v, %+v) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestIsZero(t *testing.T) {
	if !zero.IsZero() {
		t.Error("zero is not zero")
	}
	for _, a := range []Uint{one, two64, {Hi: 1 << 63}, max} {
		if a.IsZero() {
			t.Errorf("%+v is zero", a)
		}
	}
}

func TestShifts(t *testing.T) {
	for _, a := range samples {
		for _, n := range []uint{0, 1, 31, 32, 63, 64, 65, 96, 127, 128, 129, 1000} {
			wantL := fromBig(new(big.Int).Lsh(toBig(a), n))
			wantR := fromBig(new(big.Int).Rsh(toBig(a), n))
			if n >= 128 {
				wantL, wantR = zero, zero
			}
			if got := a.Lsh(n); got != wantL {
				t.Errorf("%+v << %d = %+v, want %+v", a, n, got, wantL)
			}
			if got := a.Rsh(n); got != wantR {
				t.Errorf("%+v >> %d = %+v, want %+v", a, n, got, wantR)
			}
		}
	}
	if got := one.Lsh(127); got != (Uint{Hi: 1 << 63}) {
		t.Errorf("1 << 127 = %+v", got)
	}
	if got := one.Lsh(64); got != two64 {
		t.Errorf("1 << 64 = %+v", got)
	}
	if got := max.Rsh(127); got != one {
		t.Errorf("max >> 127 = %+v", got)
	}
}

func TestLowMask(t *testing.T) {
	for _, n := range []uint{0, 1, 32, 63, 64, 65, 127, 128, 200} {
		want := max
		if n < 128 {
			want = fromBig(new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), n), big.NewInt(1)))
		}
		if got := LowMask(n); got != want {
			t.Errorf("LowMask(%d) = %+v, want %+v", n, got, want)
		}
	}
	if got := LowMask(0); got != zero {
		t.Errorf("LowMask(0) = %+v", got)
	}
	if got := LowMask(64); got != From64(^uint64(0)) {
		t.Errorf("LowMask(64) = %+v", got)
	}
}

func TestPlusOneString(t *testing.T) {
	tests := []struct {
		a    Uint
		want string
	}{
		{zero, "1"},
		{From64(9), "10"},
		{From64(^uint64(0)), "18446744073709551616"},
		{two64, "18446744073709551617"},
		{max, "340282366920938463463374607431768211456"},
	}
	for _, tt := range tests {
		if got := tt.a.PlusOneString(); got != tt.want {
			t.Errorf("%+v + 1 = %s, want %s", tt.a, got, tt.want)
		}
	}
}
