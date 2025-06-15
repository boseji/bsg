// time_test.go - Test Program `gen` Package
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
	"testing"
	"time"

	"github.com/boseji/bsg/gen"
)

func TestBST(t *testing.T) {
	result := gen.BST()

	if result.IsZero() {
		t.Fatal("BST() returned zero time")
	}

	// Check location name
	loc := result.Location()
	if loc == nil || loc.String() != "Asia/Kolkata" {
		t.Errorf("expected location Asia/Kolkata, got %v", loc)
	}

	// Check UTC offset is +05:30
	_, offset := result.Zone()
	expectedOffset := 5*60*60 + 30*60 // 19800 seconds
	if offset != expectedOffset {
		t.Errorf("expected offset +05:30 (19800), got %d", offset)
	}
}

func TestToBST_UTCInput(t *testing.T) {
	input := time.Date(2025, 6, 15, 12, 0, 0, 0, time.UTC)
	got := gen.ToBST(input)

	if got.Location().String() != "Asia/Kolkata" {
		t.Errorf("expected location Asia/Kolkata, got %s", got.Location())
	}

	expectedHour := 17
	expectedMinute := 30

	if got.Hour() != expectedHour || got.Minute() != expectedMinute {
		t.Errorf("expected %02d:%02d, got %02d:%02d", expectedHour, expectedMinute, got.Hour(), got.Minute())
	}

	if !got.UTC().Equal(input.UTC()) {
		t.Errorf("expected same instant, got %v vs %v", got.UTC(), input.UTC())
	}
}

func TestToBST_NonUTCInput(t *testing.T) {
	// New York is UTC-4 in summer (EDT)
	locNY, _ := time.LoadLocation("America/New_York")
	input := time.Date(2025, 6, 15, 7, 0, 0, 0, locNY) // equivalent to 11:00 UTC
	got := gen.ToBST(input)

	// In Asia/Kolkata, this should be 16:30
	expectedHour := 16
	expectedMinute := 30

	if got.Location().String() != "Asia/Kolkata" {
		t.Errorf("expected location Asia/Kolkata, got %s", got.Location())
	}
	if got.Hour() != expectedHour || got.Minute() != expectedMinute {
		t.Errorf("expected %02d:%02d, got %02d:%02d", expectedHour, expectedMinute, got.Hour(), got.Minute())
	}

	if !got.UTC().Equal(input.UTC()) {
		t.Errorf("expected same instant, got %v vs %v", got.UTC(), input.UTC())
	}
}

func TestToBST_ZeroTime(t *testing.T) {
	input := time.Time{}
	got := gen.ToBST(input)

	if !got.IsZero() {
		t.Errorf("expected zero time, got %v", got)
	}
}
