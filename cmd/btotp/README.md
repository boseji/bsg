>
> ॐᳬ᳞ भूर्भुवः स्वः
> 
> तत्स॑वि॒तुर्वरे॑ण्यं॒
> 
> भर्गो॑ दे॒वस्य॑ धीमहि।
> 
> धियो॒ यो नः॑ प्रचो॒दया॑त्॥
>

#  बी.स.जी - बोसजी के द्वारा रचित सुरक्षा एवं गोपनीयता हेतु तन्त्राक्ष्।

> एक सुरक्षा एवं गोपनीयता केंद्रित तंत्राक्षों का संकलन। 

***एक रचनात्मक भारतीय उत्पाद ।***

## bsg - Boseji's Security and Privacy Utilities

A collection of Security and Privacy utilities and some notes for help.

This is **Golang** package collection as well as few utility command line programs.

## btotp - Boseji's Time-based One-Time Password Utility

A simple command-line tool written in Go that generates Time-based One-Time Passwords (TOTP)
using HMAC-SHA1 with a 30-second time step and 6-digit output. The program supports multiple
secrets and can read from an external JSON file or fall back to embedded secrets using Go's
standard `embed` package.

---

### Overview
This tool computes TOTP codes using SHA1, a 30‑second interval, and produces 6‑digit codes.
It supports multiple secrets (each with a descriptive name) that can be loaded either from an
external JSON file or from an embedded file. 

---

### Prerequisites
- **Go 1.24 or later** (required for the `embed` package)
- **GNU Make** (to use the provided Makefile)

---

### Building

For your current platform:
```bash
go build -o btotp ./cmd/btotp
```

For cross-compilation (set `GOOS` and `GOARCH` accordingly), for example, for Windows (amd64):
```bash
GOOS=windows GOARCH=amd64 go build -o btotp-windows.exe ./cmd/btotp
```
---
### Usage

When run, the program generates and displays TOTP codes for each secret found in the JSON data.

### Using Embedded Secrets

If no external file is specified, the embedded `data/secrets.json` is used:

```bash
./btotp
```

### Specifying an External Secrets File

Use the `-file` or `-f` flag to specify a JSON file:

```bash
./btotp -f=path/to/your_secrets.json
```

### Displaying Help

For usage information, run:

```bash
./btotp -h
```

or

```bash
./btotp --help
```
---
### `secrets.json` File Format

The secrets file must be a JSON array of objects. Each object must include:

- **name**: A descriptive name for the secret
- **secret**: A Base32-encoded secret string

Example:

```json
[
  {
    "name": "Google",
    "secret": "JBSWY3DPEHPK3PXP"
  },
  {
    "name": "GitHub",
    "secret": "KRSXG5DSNRXW4=="
  }
]
```

**Notes:**

- The program normalizes each secret by trimming whitespace, converting it to uppercase,
  and removing all `=` characters.
- Ensure the Base32-encoded secrets are valid and do not include extra spaces or invalid
  characters.

---
### Troubleshooting Tips

### Illegal Base32 Data Error

If you encounter an error like `illegal base32 data at input byte ...`, verify that your
secret strings are correctly formatted. Remove any extraneous whitespace or invalid characters.
Although the program normalizes the secrets, double-check the input if issues persist.

### File Not Found

If the external JSON file cannot be read, ensure the path provided with `-file` or `-f`
is correct and accessible.

### Go Version Issues

This project requires Go 1.16+ due to the use of the `embed` package. Check your Go version with:
```bash
go version
```

### Cross-Compilation Errors

If errors occur during cross-compilation, ensure that your Go environment is correctly configured
for cross-compiling. Some target platforms might require additional tools or settings.

----

## Acknowledgments

- [RFC 6238](https://tools.ietf.org/html/rfc6238) for the TOTP specification
- Thanks to the Go standard library for robust packages supporting cryptography, JSON handling,
  and file embedding

## License

`SPDX: GPL-3.0-or-later`

Sources: <https://github.com/boseji/bsg>

`bsg` - Boseji's Security and Privacy Utilities.

Copyright (C) 2025 by Abhijit Bose (aka. Boseji)

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.

