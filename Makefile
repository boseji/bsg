# Project Makefile
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

# bring in the btotp-specific rules
include cmd/btotp/btotp.mk

.PHONY: clean

clean: clean_btotp

