//go:build darwin

// kyb_darwin.go - Part of the `kyb` Package for MacOS Implementation
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

package kyb

import (
	"os/exec"
	"time"
)

func available() bool {
	_, err := exec.LookPath("osascript")
	return err == nil
}

func typeText(text string) error {
	if !available() {
		return ErrNotSupported
	}

	delay := getDelay()

	for _, r := range text {
		script := `tell application "System Events" to keystroke "` + string(r) + `"`
		if err := exec.Command("osascript", "-e", script).Run(); err != nil {
			return err
		}

		if delay > 0 {
			time.Sleep(delay)
		}
	}
	return nil
}

func keyPress(key string) error {
	if !available() {
		return ErrNotSupported
	}

	script := `tell application "System Events" to key code ` + macKeyCode(key)
	return exec.Command("osascript", "-e", script).Run()
}

func quote(s string) string {
	return `"` + s + `"`
}

func macKeyCode(key string) string {
	switch key {
	case "enter":
		return "36"
	case "tab":
		return "48"
	case "esc":
		return "53"
	case "space":
		return "49"
	default:
		return "0"
	}
}
