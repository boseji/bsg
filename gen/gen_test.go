// gen_test.go - Test Program `gen` Package
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

package gen_test

import (
	"testing"

	"github.com/boseji/bsg/gen"
)

func TestIntN(t *testing.T) {
	for i := 0; i < 100; i++ {
		v := gen.IntN(100)
		if v < 0 || v >= 100 {
			t.Errorf("gen.IntN out of range: %d", v)
		}
	}
}

func TestFloat64(t *testing.T) {
	for i := 0; i < 100; i++ {
		f := gen.Float64()
		if f < 0.0 || f >= 1.0 {
			t.Errorf("gen.Float64 out of range: %f", f)
		}
	}
}

func TestFloat32(t *testing.T) {
	for i := 0; i < 100; i++ {
		f := gen.Float32()
		if f < 0.0 || f >= 1.0 {
			t.Errorf("gen.Float32 out of range: %f", f)
		}
	}
}

func TestPerm(t *testing.T) {
	n := 10
	perm := gen.Perm(n)
	if len(perm) != n {
		t.Fatalf("gen.Perm returned slice of incorrect length: %d", len(perm))
	}
	seen := make(map[int]bool)
	for _, v := range perm {
		if v < 0 || v >= n {
			t.Errorf("Value out of range in permutation: %d", v)
		}
		if seen[v] {
			t.Errorf("Duplicate value in permutation: %d", v)
		}
		seen[v] = true
	}
}

func TestShuffle(t *testing.T) {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	gen.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})
	// Ensure all original elements are still present
	seen := make(map[int]bool)
	for _, v := range data {
		seen[v] = true
	}
	for i := 0; i < 10; i++ {
		if !seen[i] {
			t.Errorf("Value missing after shuffle: %d", i)
		}
	}
}

func TestBounds(t *testing.T) {
	_ = gen.Int()
	_ = gen.Int32()
	_ = gen.Int64()
	_ = gen.Uint()
	_ = gen.Uint32()
	_ = gen.Uint64()
	_ = gen.UintN(100)
	_ = gen.Uint32N(100)
	_ = gen.Uint64N(100)
	_ = gen.Int32N(100)
	_ = gen.Int64N(100)
	_ = gen.IntN(100)
}

func TestIntNZeroPanics(t *testing.T) {
	defer func() {
		recover() // Must recover from panic
	}()
	_ = gen.IntN(0)
	t.Errorf("Expected panic for gen.IntN(0)")
}
