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
# SPDX: GPL-3.0-or-later
#
#   bsg - Boseji's Security and Privacy Utilities
#   Copyright (C) 2025 by Abhijit Bose (aka. Boseji)
#
#   This program is free software: you can redistribute it and/or modify 
#   it under the terms of the GNU General Public License as published by the 
#   Free Software Foundation, either version 3 of the License, or 
#   (at your option) any later version.
#
#   This program is distributed in the hope that it will be useful, 
#   but WITHOUT ANY WARRANTY; without even the implied warranty 
#   of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. 
#   See the GNU General Public License for more details.
#
#   You should have received a copy of the GNU General Public License along 
#   with this program. If not, see <https://www.gnu.org/licenses/>.
#

# bring in the btotp-specific rules
include cmd/btotp/btotp.mk

.PHONY: clean

clean: clean_btotp

