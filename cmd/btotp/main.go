// main.go - Part of the `btotp` Utility
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

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	// Import embed package.
	_ "embed"

	"github.com/boseji/bsg/totp"
)

//go:embed data/secrets.json
var embeddedSecrets []byte

// SecretEntry represents an entry in the JSON file.
type SecretEntry struct {
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

func main() {
	var secretsFile string

	fmt.Println("Boseji's TOTP generator v0.1")
	// Define the file flag with both long and short versions.
	flag.StringVar(&secretsFile, "file", "", "path to the secrets JSON file")
	flag.StringVar(&secretsFile, "f", "", "path to the secrets JSON file (shorthand)")

	// Define help flags.
	var showHelp bool
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showHelp, "h", false, "display help (shorthand)")

	// Custom usage message.
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-file filename]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	// If help flag is provided, display usage and exit.
	if showHelp {
		flag.Usage()
		os.Exit(0)
	}

	var data []byte
	var err error

	if secretsFile != "" {
		// Read from the specified JSON file.
		data, err = os.ReadFile(secretsFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", secretsFile, err)
			os.Exit(1)
		}
	} else {
		// Fallback to embedded secrets.
		data = embeddedSecrets
	}

	// Parse the JSON data.
	var secrets []SecretEntry
	if err = json.Unmarshal(data, &secrets); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// fmt.Printf("Secrets: %+v\n", secrets)

	// Generate and display the TOTP for each secret.
	for _, entry := range secrets {
		totp := totp.GenerateTOTP(entry.Secret)
		fmt.Printf("Name: %-10s TOTP: %s\n", entry.Name, totp)
	}
}
