// hash.go - Part of the `gen` Package
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

package gen

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	//"golang.org/x/crypto/sha3" // Older Dependency
)

// Checked using
// https://emn178.github.io/online-tools/shake128/

// SHA1 computes the SHA-1 hash of input data.
func SHA1(data []byte) []byte {
	h := sha1.New()
	h.Write(data)
	return h.Sum(nil)
}

// SHA256 computes the SHA-256 hash of input data.
func SHA256(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

// SHA384 computes the SHA-384 hash of input data.
func SHA384(data []byte) []byte {
	h := sha512.New384()
	h.Write(data)
	return h.Sum(nil)
}

// SHA512 computes the SHA-512 hash of input data.
func SHA512(data []byte) []byte {
	h := sha512.New()
	h.Write(data)
	return h.Sum(nil)
}

// SHA3_256 computes the SHA3-256 hash of input data.
func SHA3_256(data []byte) []byte {
	h := sha3.New256()
	h.Write(data)
	return h.Sum(nil)
}

// SHA3_512 computes the SHA3-512 hash of input data.
func SHA3_512(data []byte) []byte {
	h := sha3.New512()
	h.Write(data)
	return h.Sum(nil)
}

// SHAKE128 returns a SHAKE128 hash of the given length (in bytes).
func SHAKE128(data []byte, length int) []byte {
	return sha3.SumSHAKE128(data, length)
}

// SHAKE256 returns a SHAKE256 hash of the given length (in bytes).
func SHAKE256(data []byte, length int) []byte {
	return sha3.SumSHAKE256(data, length)
}
