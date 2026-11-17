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

// Package u128 implements the 128-bit unsigned arithmetic that ipcalc needs
// for IPv6 addresses.
package u128

import (
	"math/big"
	"math/bits"
)

// Uint holds an unsigned 128-bit integer, such as an address of either
// family; IPv4 uses the low 32 bits.
type Uint struct{ Hi, Lo uint64 }

// From64 returns v as a Uint.
func From64(v uint64) Uint { return Uint{0, v} }

// Add returns a+b modulo 2^128.
func (a Uint) Add(b Uint) Uint {
	lo, c := bits.Add64(a.Lo, b.Lo, 0)
	hi, _ := bits.Add64(a.Hi, b.Hi, c)
	return Uint{hi, lo}
}

// Sub returns a-b modulo 2^128.
func (a Uint) Sub(b Uint) Uint {
	lo, bo := bits.Sub64(a.Lo, b.Lo, 0)
	hi, _ := bits.Sub64(a.Hi, b.Hi, bo)
	return Uint{hi, lo}
}

// Cmp returns -1, 0 or 1 as a is less than, equal to or greater than b.
func (a Uint) Cmp(b Uint) int {
	switch {
	case a.Hi != b.Hi:
		if a.Hi < b.Hi {
			return -1
		}
		return 1
	case a.Lo != b.Lo:
		if a.Lo < b.Lo {
			return -1
		}
		return 1
	}
	return 0
}

// IsZero reports whether a is zero.
func (a Uint) IsZero() bool { return a.Hi == 0 && a.Lo == 0 }

// And returns the bitwise AND of a and b.
func (a Uint) And(b Uint) Uint { return Uint{a.Hi & b.Hi, a.Lo & b.Lo} }

// Or returns the bitwise OR of a and b.
func (a Uint) Or(b Uint) Uint { return Uint{a.Hi | b.Hi, a.Lo | b.Lo} }

// Lsh returns a shifted left by n bits; a shift of 128 or more gives zero.
func (a Uint) Lsh(n uint) Uint {
	switch {
	case n == 0:
		return a
	case n < 64:
		return Uint{a.Hi<<n | a.Lo>>(64-n), a.Lo << n}
	case n < 128:
		return Uint{a.Lo << (n - 64), 0}
	}
	return Uint{}
}

// Rsh returns a shifted right by n bits; a shift of 128 or more gives zero.
func (a Uint) Rsh(n uint) Uint {
	switch {
	case n == 0:
		return a
	case n < 64:
		return Uint{a.Hi >> n, a.Lo>>n | a.Hi<<(64-n)}
	case n < 128:
		return Uint{0, a.Hi >> (n - 64)}
	}
	return Uint{}
}

// LowMask returns a value with the n low bits set; n of 128 or more sets all.
func LowMask(n uint) Uint {
	if n >= 128 {
		return Uint{^uint64(0), ^uint64(0)}
	}
	return From64(1).Lsh(n).Sub(From64(1))
}

// PlusOneString returns a+1 in decimal without overflowing.
func (a Uint) PlusOneString() string {
	v := new(big.Int).SetUint64(a.Hi)
	v.Lsh(v, 64)
	v.Or(v, new(big.Int).SetUint64(a.Lo))
	return v.Add(v, big.NewInt(1)).String()
}
