// kyb.go - Part of the `kyb` Package
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

// Package kyb is a small, OS-aware Go package for simulating keyboard input on
// Windows, Linux, and macOS using a single, unified API.
package kyb

import (
	"errors"
	"sync"
	"time"
)

var ErrNotSupported = errors.New("keyboard simulation not supported")

var (
	delayMu  sync.RWMutex
	keyDelay time.Duration
)

// SetDelay sets the default delay between keystrokes
func SetDelay(d time.Duration) {
	delayMu.Lock()
	keyDelay = d
	delayMu.Unlock()
}

// getDelay is used internally by OS backends
func getDelay() time.Duration {
	delayMu.RLock()
	defer delayMu.RUnlock()
	return keyDelay
}

// Type types a full string
func Type(text string) error {
	return typeText(text)
}

// TypeFast disables delay temporarily
func TypeFast(text string) error {
	delayMu.Lock()
	old := keyDelay
	keyDelay = 0
	delayMu.Unlock()

	err := typeText(text)

	delayMu.Lock()
	keyDelay = old
	delayMu.Unlock()

	return err
}

// KeyPress presses and releases a key (e.g. "ctrl+c", "alt+tab", "a")
func KeyPress(key string) error {
	return keyPress(key)
}

// Available checks if backend dependencies exist
func Available() bool {
	return available()
}
