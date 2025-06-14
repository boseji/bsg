// gen.go - Part of the `gen` Package
//
//     ॐ भूर्भुवः स्वः
//     तत्स॑वि॒तुर्वरे॑ण्यं॒
//    भर्गो॑ दे॒वस्य॑ धीमहि।
//   धियो॒ यो नः॑ प्रचो॒दया॑त्॥
//
//
//  बी.स.जी - बोसजी के द्वारा रचित सुरक्षा एवं गोपनीयता हेतु तन्त्राक्ष्।
// ================================================
//
// एक सुरक्षा एवं गोपनीयता केंद्रित तंत्राक्षों का संकलन।
//
// एक रचनात्मक भारतीय उत्पाद ।
//
// bsg - Boseji's Security and Privacy Utilities
//
// A collection of Security and Privacy utilities and some notes for help.
//
// This is **Golang** package collection as well as few utility
// command line programs.
//
// Sources
// -------
// https://github.com/boseji/bsg
//
// License
// -------
//
//   bsg - Boseji's Security and Privacy Utilities.
//   Copyright (C) 2025 by Abhijit Bose (aka. Boseji)
//
//   This program is free software: you can redistribute it and/or modify
//   it under the terms of the GNU General Public License version 2 only
//   as published by the Free Software Foundation.
//
//   This program is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
//
//   You should have received a copy of the GNU General Public License
//   along with this program. If not, see <https://www.gnu.org/licenses/>.
//
//  SPDX-License-Identifier: GPL-2.0-only
//  Full Name: GNU General Public License v2.0 only
//  Please visit <https://spdx.org/licenses/GPL-2.0-only.html> for details.
//

// Package gen provides ways to generate special sequences needed for
// various application to sure transactions and IDs.
package gen

import (
	"crypto/rand"
	"encoding/binary"
	"math/big"
)

// Uint64 generates a secure random uint64 value.
func Uint64() uint64 {
	var b [8]byte
	_, err := rand.Read(b[:])
	if err != nil {
		panic("crypto/rand failed: " + err.Error())
	}
	return binary.LittleEndian.Uint64(b[:])
}

// Uint64N returns a secure random uint64 in [0, n].
func Uint64N(n uint64) uint64 {
	if n == 0 {
		panic("Uint64N: argument must be > 0")
	}
	max := new(big.Int).SetUint64(n)
	r, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic("crypto/rand failed: " + err.Error())
	}
	return r.Uint64()
}

// Uint32 returns a secure random uint32.
func Uint32() uint32 {
	return uint32(Uint64())
}

// Uint32N returns a secure random uint32 in [0, n].
func Uint32N(n uint32) uint32 {
	return uint32(Uint64N(uint64(n)))
}

// Uint returns a secure random uint.
func Uint() uint {
	return uint(Uint64())
}

// UintN returns a secure random uint in [0, n].
func UintN(n uint) uint {
	return uint(Uint64N(uint64(n)))
}

// Int64 returns a secure random int64.
func Int64() int64 {
	return int64(Uint64())
}

// Int64N returns a secure random int64 in [0, n].
func Int64N(n int64) int64 {
	return int64(Uint64N(uint64(n)))
}

// Int32 returns a secure random int32.
func Int32() int32 {
	return int32(Uint64())
}

// Int32N returns a secure random int32 in [0, n].
func Int32N(n int32) int32 {
	return int32(Uint64N(uint64(n)))
}

// Int returns a secure random int.
func Int() int {
	return int(Uint64())
}

// IntN returns a secure random int in [0, n].
func IntN(n int) int {
	return int(Uint64N(uint64(n)))
}

// Float64 returns a secure random float64 in [0.0, 1.0].
func Float64() float64 {
	const bits = 53
	r := Uint64() >> (64 - bits)
	return float64(r) / (1 << bits)
}

// Float32 returns a secure random float32 in [0.0, 1.0].
func Float32() float32 {
	const bits = 24
	r := Uint64() >> (64 - bits)
	return float32(r) / float32(1<<bits)
}

// Shuffle shuffles the indices [0, n) using the provided swap function.
func Shuffle(n int, swap func(i, j int)) {
	for i := n - 1; i > 0; i-- {
		j := IntN(i + 1)
		swap(i, j)
	}
}

// Perm returns a random permutation of the integers [0, n).
func Perm(n int) []int {
	m := make([]int, n)
	for i := 0; i < n; i++ {
		m[i] = i
	}
	Shuffle(n, func(i, j int) {
		m[i], m[j] = m[j], m[i]
	})
	return m
}
