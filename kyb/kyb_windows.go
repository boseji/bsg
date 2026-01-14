//go:build windows

// kyb_windows.go - Part of the `kyb` Package for Windows Implementation
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
	"syscall"
	"time"
	"unicode"
	"unsafe"
)

var (
	user32    = syscall.NewLazyDLL("user32.dll")
	sendInput = user32.NewProc("SendInput")
)

const (
	INPUT_KEYBOARD  = 1
	KEYEVENTF_KEYUP = 0x0002
)

type KEYBDINPUT struct {
	Vk        uint16
	Scan      uint16
	Flags     uint32
	Time      uint32
	ExtraInfo uintptr
}

type INPUT struct {
	Type uint32
	Ki   KEYBDINPUT
}

func available() bool {
	return true // user32 is always present
}

func sendKey(vk uint16, up bool) {
	flags := uint32(0)
	if up {
		flags = KEYEVENTF_KEYUP
	}

	input := INPUT{
		Type: INPUT_KEYBOARD,
		Ki: KEYBDINPUT{
			Vk:    vk,
			Flags: flags,
		},
	}

	sendInput.Call(
		1,
		uintptr(unsafe.Pointer(&input)),
		unsafe.Sizeof(input),
	)
}

func typeText(text string) error {
	delay := getDelay()

	for _, r := range text {
		if unicode.IsUpper(r) {
			sendKey(0x10, false) // Shift down
		}

		vk := charToVK(r)
		sendKey(vk, false)
		sendKey(vk, true)

		if unicode.IsUpper(r) {
			sendKey(0x10, true) // Shift up
		}

		if delay > 0 {
			time.Sleep(delay)
		}
	}
	return nil
}

func keyPress(key string) error {
	vk := keyStringToVK(key)
	if vk == 0 {
		return ErrNotSupported
	}

	sendKey(vk, false)
	sendKey(vk, true)
	return nil
}

func charToVK(r rune) uint16 {
	if r >= 'a' && r <= 'z' {
		return uint16(r - 'a' + 0x41)
	}
	if r >= 'A' && r <= 'Z' {
		return uint16(r - 'A' + 0x41)
	}
	if r >= '0' && r <= '9' {
		return uint16(r)
	}

	switch r {
	case ' ':
		return 0x20
	case '\n':
		return 0x0D
	default:
		return 0
	}
}

func keyStringToVK(key string) uint16 {
	switch key {
	case "enter":
		return 0x0D
	case "esc":
		return 0x1B
	case "tab":
		return 0x09
	case "ctrl":
		return 0x11
	case "alt":
		return 0x12
	case "shift":
		return 0x10
	default:
		if len(key) == 1 {
			return charToVK(rune(key[0]))
		}
	}
	return 0
}
