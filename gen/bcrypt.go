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
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	DefaultBcryptCost int = 12
	NominalBcryptCost int = bcrypt.DefaultCost
	MinBcryptCost     int = bcrypt.MinCost
	MaxBcryptCost     int = bcrypt.MaxCost
)

// BcryptHash helps to Hash a password using the Bcrypt Algorithm
// With a pre-determined cost value.
func BcryptHash(password string) (string, error) {
	// DefaultCost = 10 (adjust higher for slower hashing)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),
		DefaultBcryptCost)
	return string(bytes), err
}

// BcryptHashC helps to Hash a password using the Bcrypt Algorithm
// With a supplied cost value.
func BcryptHashC(password string, cost int) (string, error) {
	if cost < MinBcryptCost {
		return "", fmt.Errorf("cost too low")
	}
	if cost > MaxBcryptCost {
		return "", fmt.Errorf("cost too low")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

// BcryptCheck helps to Verify a password matches into the Hash or not
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
