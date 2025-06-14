// str_test.go - Test Program `gen` Package
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
	"unicode/utf8"

	"github.com/boseji/bsg/gen"
)

// TestStringValid tests normal operation with common charsets.
func TestStringValid(t *testing.T) {
	charsets := []string{
		"abc",                        // small charset
		"0123456789",                 // digits
		"abcdefghijklmnopqrstuvwxyz", // lowercase
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ", // uppercase
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", // alphanumeric
	}
	lengths := []int{1, 5, 10, 32, 100}

	for _, charset := range charsets {
		for _, length := range lengths {
			result, err := gen.String(charset, length)
			if err != nil {
				t.Errorf("unexpected error for charset='%s' length=%d: %v", charset, length, err)
			}
			if len(result) != length {
				t.Errorf("expected length %d, got %d for result='%s'", length, len(result), result)
			}
			for _, ch := range result {
				if !containsRune(charset, ch) {
					t.Errorf("character '%c' in result not found in charset '%s'", ch, charset)
				}
			}
		}
	}
}

// TestStringEmptyCharset ensures error on empty charset.
func TestStringEmptyCharset(t *testing.T) {
	_, err := gen.String("", 10)
	if err == nil {
		t.Error("expected error for empty charset, got nil")
	}
}

// TestStringZeroLength ensures error on zero length.
func TestStringZeroLength(t *testing.T) {
	_, err := gen.String("abc", 0)
	if err == nil {
		t.Error("expected error for zero length, got nil")
	}
}

// TestStringNegativeLength ensures error on negative length.
func TestStringNegativeLength(t *testing.T) {
	_, err := gen.String("abc", -5)
	if err == nil {
		t.Error("expected error for negative length, got nil")
	}
}

// TestStringRandomness checks randomness across multiple calls.
func TestStringRandomness(t *testing.T) {
	charset := "abcdef"
	length := 16
	count := 10
	results := make(map[string]bool)

	for i := 0; i < count; i++ {
		result, err := gen.String(charset, length)
		if err != nil {
			t.Fatalf("error generating string: %v", err)
		}
		if results[result] {
			t.Errorf("duplicate string generated: %s", result)
		}
		results[result] = true
	}
}

// Helper: checks if rune exists in charset string
func containsRune(charset string, r rune) bool {
	for _, ch := range charset {
		if ch == r {
			return true
		}
	}
	return false
}

// TestStringUTF8Charset ensures support for UTF-8 characters.
func TestStringUTF8Charset(t *testing.T) {
	charset := "😀😃😄😁😆"
	length := 10

	result, err := gen.String(charset, length)
	if err != nil {
		t.Fatalf("unexpected error with utf8 charset: %v", err)
	}

	// Confirm length and validity of characters
	if utf8.RuneCountInString(result) != length {
		t.Errorf("expected %d runes, got %d", length, utf8.RuneCountInString(result))
	}

	for _, ch := range result {
		if !containsRune(charset, ch) {
			t.Errorf("character '%c' in result not found in charset '%s'", ch, charset)
		}
	}
}
