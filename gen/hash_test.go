// hash_test.go - Test Program `gen` Package
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
	"encoding/hex"
	"testing"

	"github.com/boseji/bsg/gen"
)

// Helper to convert hex string to []byte
func fromHex(t *testing.T, hexStr string) []byte {
	b, err := hex.DecodeString(hexStr)
	if err != nil {
		t.Fatalf("invalid hex string: %v", err)
	}
	return b
}

func TestSHA1(t *testing.T) {
	input := []byte("abc")
	expected := fromHex(t, "a9993e364706816aba3e25717850c26c9cd0d89d")
	output := gen.SHA1(input)
	if string(output) != string(expected) {
		t.Errorf("SHA1 failed: got %x, expected %x", output, expected)
	}
}

func TestSHA256(t *testing.T) {
	input := []byte("abc")
	expected := fromHex(t,
		"ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad",
	)
	output := gen.SHA256(input)
	if string(output) != string(expected) {
		t.Errorf("SHA256 failed: got %x, expected %x", output, expected)
	}
}

func TestSHA384(t *testing.T) {
	input := []byte("abc")
	expected := fromHex(t,
		"cb00753f45a35e8bb5a03d699ac65007272c32ab0eded163"+
			"1a8b605a43ff5bed8086072ba1e7cc2358baeca134c825a7",
	)
	output := gen.SHA384(input)
	if string(output) != string(expected) {
		t.Errorf("SHA384 failed: got %x, expected %x", output, expected)
	}
}

func TestSHA512(t *testing.T) {
	input := []byte("abc")
	expected := fromHex(t,
		"ddaf35a193617abacc417349ae20413112e6fa4e89a97ea2"+
			"0a9eeee64b55d39a2192992a274fc1a836ba3c23a3feebbd"+
			"454d4423643ce80e2a9ac94fa54ca49f",
	)
	output := gen.SHA512(input)
	if string(output) != string(expected) {
		t.Errorf("SHA512 failed: got %x, expected %x", output, expected)
	}
}

func TestSHA3_256(t *testing.T) {
	input := []byte("abc")
	expected := fromHex(t,
		"3a985da74fe225b2045c172d6bd390bd855f086e3e9d525b46bfe24511431532",
	)
	output := gen.SHA3_256(input)
	if string(output) != string(expected) {
		t.Errorf("SHA3-256 failed: got %x, expected %x", output, expected)
	}
}

func TestSHA3_512(t *testing.T) {
	input := []byte("abc")
	expected := fromHex(t,
		"b751850b1a57168a5693cd924b6b096e08f621827444f70d"+
			"884f5d0240d2712e10e116e9192af3c91a7ec57647e39340"+
			"57340b4cf408d5a56592f8274eec53f0",
	)
	output := gen.SHA3_512(input)
	if string(output) != string(expected) {
		t.Errorf("SHA3-512 failed: got %x, expected %x", output, expected)
	}
}

func TestSHAKE128(t *testing.T) {
	input := []byte("abc")
	expected := fromHex(t,
		"5881092dd818bf5cf8a3ddb793fbcba74097d5c526a6d35f97b83351940f2cc8",
	)
	output := gen.SHAKE128(input, 32)
	if string(output) != string(expected) {
		t.Errorf("SHAKE128 failed:\n  got:      %x\n  expected: %x", output, expected)
	}
}

func TestSHAKE256(t *testing.T) {
	input := []byte("abc")
	expected := fromHex(t,
		"483366601360a8771c6863080cc4114d8db44530f8f1e1ee"+
			"4f94ea37e78b5739d5a15bef186a5386c75744c0527e1faa"+
			"9f8726e462a12a4feb06bd8801e751e4",
	)
	output := gen.SHAKE256(input, 64)
	if string(output) != string(expected) {
		t.Errorf("SHAKE256 failed:\n  got:      %x\n  expected: %x", output, expected)
	}
}
