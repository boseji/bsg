// totp_test.go - Part of the `totp` Package
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
// SPDX: GPL-3.0-or-later
//
//   bsg - Boseji's Security and Privacy Utilities.
//   Copyright (C) 2025 by Abhijit Bose (aka. Boseji)
//
//   This program is free software: you can redistribute it and/or modify
//   it under the terms of the GNU General Public License as published by the
//   Free Software Foundation, either version 3 of the License, or
//   (at your option) any later version.
//
//   This program is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty
//   of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
//   See the GNU General Public License for more details.
//
//   You should have received a copy of the GNU General Public License along
//   with this program. If not, see <https://www.gnu.org/licenses/>.
//

package totp

import (
	"crypto/sha256"
	"fmt"
	"testing"
	"time"
)

// TestGenerate runs several test cases to verify the behavior of Generate.
func TestGenerate(t *testing.T) {
	// The Base32-encoded secret below is the encoding of "12345678901234567890".
	// It is used in RFC 6238 test vectors.
	secret := "GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ"

	// Define test cases.
	tests := []struct {
		name        string
		secret      string
		opts        []Option
		expectedOTP string
		expectError bool
	}{
		{
			name:        "default options with time=59 (SHA1, 6-digit)",
			secret:      secret,
			opts:        []Option{WithTime(time.Unix(59, 0))},
			expectedOTP: "287082",
		},
		{
			name:        "default options with time=1111111109 (SHA1, 6-digit)",
			secret:      secret,
			opts:        []Option{WithTime(time.Unix(1111111109, 0))},
			expectedOTP: "081804",
		},
		{
			name:        "8-digit OTP with time=59 (SHA1, 8-digit)",
			secret:      secret,
			opts:        []Option{WithTime(time.Unix(59, 0)), WithDigits(8)},
			expectedOTP: "94287082",
		},
		{
			name:        "default 6-digit OTP with SHA256 with time=59",
			secret:      secret,
			opts:        []Option{WithTime(time.Unix(59, 0)), WithAlgorithm(sha256.New)},
			expectedOTP: "247374", //"918194",
		},
		{
			name:        "invalid secret returns error",
			secret:      "INVALIDSECRET!",
			opts:        []Option{},
			expectError: true,
		},
	}

	// Run test cases.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			otp, err := Generate(tc.secret, tc.opts...)
			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error, got nil with OTP %s", otp)
				}
				return
			}
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if otp != tc.expectedOTP {
				t.Errorf("OTP mismatch: expected %s, got %s", tc.expectedOTP, otp)
			}
		})
	}
}

// ExampleGenerate demonstrates using Generate with default options and a custom time.
func ExampleGenerate() {
	// The Base32-encoded secret is that of "12345678901234567890".
	secret := "GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ"

	// Use a fixed time (Unix time 59) so the output is predictable.
	otp, _ := Generate(secret, WithTime(time.Unix(59, 0)))
	fmt.Println(otp)
	// Output:
	// 287082
}
