// bcrypt.go - Part of the `gen` Package
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
//   Copyright (C) 2025-2026 by Abhijit Bose (aka. Boseji)
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
	"testing"

	"golang.org/x/crypto/bcrypt"
)

// TestBcryptHash tests the basic BcryptHash function
func TestBcryptHash(t *testing.T) {
	tests := []struct {
		name     string
		password string
	}{
		{
			name:     "simple password",
			password: "mypassword",
		},
		{
			name:     "empty password",
			password: "",
		},
		{
			name:     "long password",
			password: "this is a very long password with special characters!@#$%^&*()",
		},
		{
			name:     "unicode password",
			password: "пароль密码🔐",
		},
		{
			name:     "password with spaces",
			password: "pass word with spaces",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := BcryptHash(tt.password)
			if err != nil {
				t.Fatalf("BcryptHash() error = %v, want nil", err)
			}

			if hash == "" {
				t.Error("BcryptHash() returned empty hash")
			}

			// Verify the hash is valid and matches the password
			if !BcryptCheck(tt.password, hash) {
				t.Error("BcryptCheck() failed to verify the hash")
			}
		})
	}
}

// TestBcryptHashC tests the BcryptHashC function with different cost values
func TestBcryptHashC(t *testing.T) {
	password := "testpassword"

	tests := []struct {
		name    string
		cost    int
		wantErr bool
	}{
		{
			name:    "minimum cost",
			cost:    bcrypt.MinCost,
			wantErr: false,
		},
		{
			name:    "default cost",
			cost:    bcrypt.DefaultCost,
			wantErr: false,
		},
		/*{ // Takes too much of Time so ignored
			name:    "maximum cost",
			cost:    bcrypt.MaxCost,
			wantErr: false,
		},*/
		{
			name:    "custom cost 10",
			cost:    10,
			wantErr: false,
		},
		{
			name:    "custom cost 14",
			cost:    14,
			wantErr: false,
		},
		{
			name:    "invalid cost too low",
			cost:    bcrypt.MinCost - 1,
			wantErr: true,
		},
		{
			name:    "invalid cost too high",
			cost:    bcrypt.MaxCost + 1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := BcryptHashC(password, tt.cost)

			if (err != nil) != tt.wantErr {
				t.Fatalf("BcryptHashC() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && hash == "" {
				t.Error("BcryptHashC() returned empty hash")
			}

			if !tt.wantErr && !BcryptCheck(password, hash) {
				t.Error("BcryptCheck() failed to verify hash with custom cost")
			}
		})
	}
}

// TestBcryptCheck tests the BcryptCheck function
func TestBcryptCheck(t *testing.T) {
	password := "correctpassword"
	hash, _ := BcryptHash(password)

	tests := []struct {
		name     string
		password string
		hash     string
		want     bool
	}{
		{
			name:     "correct password",
			password: password,
			hash:     hash,
			want:     true,
		},
		{
			name:     "incorrect password",
			password: "wrongpassword",
			hash:     hash,
			want:     false,
		},
		{
			name:     "empty password with valid hash",
			password: "",
			hash:     hash,
			want:     false,
		},
		{
			name:     "correct password with empty hash",
			password: password,
			hash:     "",
			want:     false,
		},
		{
			name:     "case sensitive check",
			password: "CorrectPassword",
			hash:     hash,
			want:     false,
		},
		{
			name:     "malformed hash",
			password: password,
			hash:     "notavalidhash",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BcryptCheck(tt.password, tt.hash)
			if got != tt.want {
				t.Errorf("BcryptCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestBcryptHashConsistency tests that BcryptHash uses DefaultBcryptCost
func TestBcryptHashConsistency(t *testing.T) {
	password := "testpassword"

	hash1, _ := BcryptHash(password)
	hash2, _ := BcryptHashC(password, DefaultBcryptCost)

	// Both should be verifiable (they may not be equal due to salt)
	if !BcryptCheck(password, hash1) || !BcryptCheck(password, hash2) {
		t.Error("BcryptHash() and BcryptHashC() with DefaultBcryptCost should both be valid")
	}
}

// TestBcryptHashDifferentSalts tests that hashes are different due to salt
func TestBcryptHashDifferentSalts(t *testing.T) {
	password := "samepassword"

	hash1, _ := BcryptHash(password)
	hash2, _ := BcryptHash(password)

	if hash1 == hash2 {
		t.Error("BcryptHash() should produce different hashes due to random salt")
	}

	// But both should verify correctly
	if !BcryptCheck(password, hash1) || !BcryptCheck(password, hash2) {
		t.Error("Both hashes should verify with the correct password")
	}
}
