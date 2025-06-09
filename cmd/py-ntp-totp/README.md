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

## py-ntp-totp - Boseji's Python based NTP synced TOTP Utility

This is a Python-based **TOTP (Time-based One-Time Password)** generator that uses **NTP (Network Time Protocol)** to fetch accurate UTC time and generate codes compatible with **Google Authenticator** and similar apps.

It avoids system time drift by querying `pool.ntp.org` directly via raw UDP (no external packages needed), and falls back to system time if NTP fails.

----

## ✅ Features

- Supports multiple TOTP accounts
- Uses SHA1 (Google Authenticator-compatible)
- Fetches UTC time via NTP (UDP, no external libraries)
- Falls back to local system UTC if NTP is unreachable
- Prints current TOTP for each account

---

## Requirements

`Python 3.x`

No third-party dependencies are needed.

---

## Usage

1. Clone or copy the script.
2. Define your TOTP secrets in the `accounts` list.
3. Run the script:

```bash
python3 ntp_totp.py
```

---

## Example `accounts` List

```
accounts = [
    {"name": "gmail1", "totp": "JBSWY3DPEHPK3PXP"},
    {"name": "github", "totp": "GEZDGNBVGY3TQOJQ"},
    {"name": "banking", "totp": "MZXW6YTBOI======"}
]
```

All TOTP secrets must be **base32-encoded** strings.

---

## How it Works

1. Sends a 48-byte UDP packet to `pool.ntp.org`
2. Parses the NTP response and converts to Unix epoch
3. If NTP fails, falls back to `time.time()`
4. Uses standard TOTP algorithm (RFC 6238) with SHA1
5. Sleeps momentarily to enhance resolution.

---

## Output

```
****************************************
- Boseji's TOTP Generator Program v0.1 -
****************************************

NTP Time Sync
[INFO] NTP UTC Time   : 2025-06-09 16:23:38 UTC
[INFO] Local Epoch    : 2025-06-09 16:23:38 UTC
[INFO] Clock Drift    : 0.905 seconds

TOTP Codes:
gmail1              : 236109
github              : 830869
banking             : 976865
```

---

## Files

- `ntp_totp.py`: Main script
- `README.md`: This documentation

---

## References

- [RFC 6238 – TOTP Algorithm](https://datatracker.ietf.org/doc/html/rfc6238)
- [RFC 5905 – NTP Protocol](https://datatracker.ietf.org/doc/html/rfc5905)
- [Google Authenticator TOTP format](https://github.com/google/google-authenticator)

---

## License

This project is released under the GNU General Public License v2. See the [LICENSE](../../LICENSE.txt) file for details.

Sources: <https://github.com/boseji/bsg>

`bsg` - Boseji's Security and Privacy Utilities.

Copyright (C) 2025 by Abhijit Bose (aka. Boseji)

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License version 2 only
as published by the Free Software Foundation.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

You should have received a copy of the GNU General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.

SPDX-License-Identifier: `GPL-2.0-only`

Full Name: `GNU General Public License v2.0 only`

Please visit <https://spdx.org/licenses/GPL-2.0-only.html> for details.


