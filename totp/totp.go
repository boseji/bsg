// totp.go - Part of the `totp` Package
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

// Package totp provides a generic TOTP (Time-based One-Time Password)
// generator that uses functional options to configure parameters such as period,
// digit length, hash algorithm, and the time variable.
package totp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"hash"
	"strings"
	"time"
)

// Options holds configuration parameters for TOTP generation.
type Options struct {
	Period    int              // Time step in seconds (default is 30).
	Digits    int              // Number of digits in the OTP (default is 6).
	Algorithm func() hash.Hash // Hash algorithm constructor (default is sha1.New).
	Time      time.Time        // Optional time to use. If zero, time.Now() is used.
}

// Option is a function that modifies Options.
type Option func(*Options)

// DefaultOptions returns an Options struct populated with default values.
func DefaultOptions() Options {
	return Options{
		Period:    30,
		Digits:    6,
		Algorithm: sha1.New,
	}
}

// WithPeriod sets the time step (in seconds).
func WithPeriod(period int) Option {
	return func(opts *Options) {
		opts.Period = period
	}
}

// WithDigits sets the number of digits for the OTP.
func WithDigits(digits int) Option {
	return func(opts *Options) {
		opts.Digits = digits
	}
}

// WithAlgorithm sets the hash algorithm used for HMAC.
// For example, use sha1.New (default), sha256.New, or sha512.New.
func WithAlgorithm(alg func() hash.Hash) Option {
	return func(opts *Options) {
		opts.Algorithm = alg
	}
}

// WithTime sets a custom time to be used for TOTP generation.
func WithTime(t time.Time) Option {
	return func(opts *Options) {
		opts.Time = t
	}
}

// Generate produces a TOTP code for the provided Base32-encoded secret,
// applying any functional options provided. If an option is omitted,
// default values are used. If no custom time is provided, time.Now() is used.
func Generate(secret string, opts ...Option) (string, error) {
	options := DefaultOptions()
	for _, opt := range opts {
		opt(&options)
	}

	// If no time provided, use current time.
	var t time.Time
	if options.Time.IsZero() {
		t = time.Now()
	} else {
		t = options.Time
	}

	// Normalize the secret: trim whitespace, convert to uppercase,
	// and remove any '=' characters.
	secret = strings.ToUpper(strings.TrimSpace(secret))
	secret = strings.ReplaceAll(secret, "=", "")

	// Decode the Base32 secret using a decoder that expects no padding.
	decoder := base32.StdEncoding.WithPadding(base32.NoPadding)
	key, err := decoder.DecodeString(secret)
	if err != nil {
		return "", fmt.Errorf("error decoding secret: %v", err)
	}

	// Calculate the time counter based on the provided time and period.
	counter := uint64(t.Unix() / int64(options.Period))
	var counterBytes [8]byte
	binary.BigEndian.PutUint64(counterBytes[:], counter)

	// Create an HMAC hash using the selected algorithm.
	mac := hmac.New(options.Algorithm, key)
	mac.Write(counterBytes[:])
	hashResult := mac.Sum(nil)

	// Dynamic truncation per RFC 4226.
	offset := hashResult[len(hashResult)-1] & 0x0F
	if int(offset)+4 > len(hashResult) {
		return "", fmt.Errorf("invalid offset, out of bounds")
	}
	binaryCode := (uint32(hashResult[offset])&0x7F)<<24 |
		(uint32(hashResult[offset+1])&0xFF)<<16 |
		(uint32(hashResult[offset+2])&0xFF)<<8 |
		(uint32(hashResult[offset+3]) & 0xFF)

	// Compute modulus based on the desired number of digits.
	mod := uint32(1)
	for i := 0; i < options.Digits; i++ {
		mod *= 10
	}
	otp := binaryCode % mod

	// Format the OTP with leading zeros if necessary.
	otpStr := fmt.Sprintf("%0*d", options.Digits, otp)
	return otpStr, nil
}

// GenerateTOTP computes a TOTP value based on the provided Base32-encoded secret.
// This intern is equivalent to calling Generate with default options.
// This is included for backward compatibility with the original implementation.
func GenerateTOTP(secret string) string {
	// Normalize the secret: trim whitespace, convert to uppercase, and remove all '=' characters.
	secret = strings.ToUpper(strings.TrimSpace(secret))
	secret = strings.ReplaceAll(secret, "=", "")

	// Create a Base32 decoder configured to expect no padding.
	decoder := base32.StdEncoding.WithPadding(base32.NoPadding)
	key, err := decoder.DecodeString(secret)
	if err != nil {
		// In a production app, you might want to handle the error differently.
		panic(fmt.Sprintf("Error decoding secret '%s': %v", secret, err))
	}

	// Calculate the current time step (30-second intervals)
	timeStep := uint64(time.Now().Unix() / 30)

	// Convert the time step to an 8-byte array (big-endian)
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, timeStep)

	// Create an HMAC-SHA1 hash from the key and time step
	h := hmac.New(sha1.New, key)
	h.Write(buf)
	hash := h.Sum(nil)

	// Dynamic truncation: use the last nibble of the hash as an offset
	offset := hash[len(hash)-1] & 0x0F

	// Extract a 4-byte dynamic binary code starting at the offset
	binaryCode := (uint32(hash[offset])&0x7F)<<24 |
		(uint32(hash[offset+1])&0xFF)<<16 |
		(uint32(hash[offset+2])&0xFF)<<8 |
		(uint32(hash[offset+3]) & 0xFF)

	// Calculate the OTP value (6 digits)
	otp := binaryCode % 1000000

	// Format the OTP with zero-padding to ensure it has 6 digits
	return fmt.Sprintf("%06d", otp)
}
