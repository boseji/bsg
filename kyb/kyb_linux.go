//go:build linux

// kyb_linux.go - Part of the `kyb` Package for Linux Implementation
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
	"strconv"
)

func available() bool {
	_, err := exec.LookPath("xdotool")
	return err == nil
}

func typeText(text string) error {
	if !available() {
		return ErrNotSupported
	}

	args := []string{"type"}

	if d := getDelay(); d > 0 {
		args = append(args, "--delay", strconv.Itoa(int(d.Milliseconds())))
	}

	args = append(args, text)

	return exec.Command(
		"xdotool", args...,
	).Run()
}

func keyPress(key string) error {
	if !available() {
		return ErrNotSupported
	}

	// Enter key Fix
	if key == "Enter" || key == "enter" || key == "return" {
		key = "Return"
	}

	return exec.Command(
		"xdotool",
		"key",
		key,
	).Run()
}
