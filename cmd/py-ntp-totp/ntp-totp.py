#!/usr/bin/env python3

# ntp-totp.py - Python NTP based TOTP generator
#
#
# bsg - Boseji's Security and Privacy Utilities
#
# Sources
# -------
# https://github.com/boseji/bsg
#
# License
# -------
#
#   bsg - Boseji's Security and Privacy Utilities
#   Copyright (C) 2025 by Abhijit Bose (aka. Boseji)
#
#   This program is free software: you can redistribute it and/or modify
#   it under the terms of the GNU General Public License version 2 only
#   as published by the Free Software Foundation.
#
#   This program is distributed in the hope that it will be useful,
#   but WITHOUT ANY WARRANTY; without even the implied warranty of
#   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. 
#
#   You should have received a copy of the GNU General Public License
#   along with this program. If not, see <https://www.gnu.org/licenses/>.
#
#  SPDX-License-Identifier: GPL-2.0-only
#  Full Name: GNU General Public License v2.0 only
#  Please visit <https://spdx.org/licenses/GPL-2.0-only.html> for details.
#


import hmac
import hashlib
import base64
import struct
import socket
import time

# --- NTP time fetch using raw UDP ---
def get_ntp_utc_time(server="pool.ntp.org", timeout=2, debug=False):
    NTP_PORT = 123
    NTP_PACKET = b'\x1b' + 47 * b'\0'
    NTP_DELTA = 2208988800  # NTP epoch (1900) to UNIX epoch (1970)

    try:
        with socket.socket(socket.AF_INET, socket.SOCK_DGRAM) as s:
            s.settimeout(timeout)
            s.sendto(NTP_PACKET, (server, NTP_PORT))
            msg, _ = s.recvfrom(48)

        if len(msg) != 48:
            raise ValueError("Incomplete NTP response")

        # Unpack transmit timestamp (bytes 40-43: seconds)
        transmit_time = struct.unpack("!I", msg[40:44])[0]
        ntp_utc = transmit_time - NTP_DELTA

        # Show clear distinction
        local_epoch = time.time()
        if debug:
            print(f"[INFO] NTP UTC Time   : {time.strftime('%Y-%m-%d %H:%M:%S UTC', time.gmtime(ntp_utc))}")
            print(f"[INFO] Local Epoch    : {time.strftime('%Y-%m-%d %H:%M:%S UTC', time.gmtime(local_epoch))}")
            print(f"[INFO] Clock Drift    : {local_epoch - ntp_utc:.3f} seconds")

        return int(ntp_utc)

    except Exception as e:
        print(f"[WARNING] NTP UDP request failed: {e}. Falling back to system UTC time.")
        return int(time.time())  # Already UTC
    

# --- TOTP generator ---
def generate_totp(secret_base32, digits=6, period=30, algo=hashlib.sha1, time_source=get_ntp_utc_time):
    key = base64.b32decode(secret_base32, casefold=True)
    current_time = time_source()
    counter = current_time // period
    counter_bytes = struct.pack(">Q", counter)
    hmac_hash = hmac.new(key, counter_bytes, algo).digest()
    offset = hmac_hash[-1] & 0x0F
    truncated = hmac_hash[offset:offset + 4]
    code = struct.unpack(">I", truncated)[0] & 0x7fffffff
    otp = code % (10 ** digits)
    return str(otp).zfill(digits)

# --- Dictionary of TOTP entries ---
accounts = [
    {"name": "gmail1", "totp": "JBSWY3DPEHPK3PXP"},
    {"name": "github", "totp": "GEZDGNBVGY3TQOJQ"},
    {"name": "banking", "totp": "MZXW6YTBOI======"},
]

# --- Generate for all ---
def generate_all(accounts):
    print("\nTOTP Codes:")
    for account in accounts:
        try:
            code = generate_totp(account["totp"])
            print(f"{account['name']:20}: {code}")
            time.sleep(1) # Slow down the pings by 1 second
        except Exception as e:
            print(f"{account['name']:20}: [ERROR] {e}")

if __name__ == "__main__":
    print("****************************************")
    print("- Boseji's TOTP Generator Program v0.1 -")
    print("****************************************")
    print()
    #secret = "JBSWY3DPEHPK3PXP"  # example key
    #otp = generate_totp(secret)
    #print("Generated TOTP:", otp)
    print("NTP Time Sync")
    get_ntp_utc_time(debug=True)
    generate_all(accounts)
