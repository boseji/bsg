# btotp Makefile
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

# Application name.
BTOTP_APP_NAME	:= btotp

# Output directory for binaries.
BTOTP_BINDIR	:= build

# Source dir containing main.go
BTOTP_SRCDIR	:= cmd/$(BTOTP_APP_NAME)

# targets
.PHONY: btotp clean_btotp

btotp: $(BTOTP_BINDIR) \
		$(BTOTP_BINDIR)/$(BTOTP_APP_NAME)-windows.exe \
		$(BTOTP_BINDIR)/$(BTOTP_APP_NAME)-linux-amd64 \
		$(BTOTP_BINDIR)/$(BTOTP_APP_NAME)-linux-arm \
		$(BTOTP_BINDIR)/$(BTOTP_APP_NAME)-linux-arm64 \
		$(BTOTP_BINDIR)/$(BTOTP_APP_NAME)-darwin-amd64 \
		$(BTOTP_BINDIR)/$(BTOTP_APP_NAME)-darwin-arm64

# ensure output dir exists (order-only)
$(BTOTP_BINDIR):
	mkdir -p $@

# per-platform builds
$(BTOTP_BINDIR)/$(BTOTP_APP_NAME)-windows.exe: $(BTOTP_SRCDIR)/main.go | $(BTOTP_BINDIR)
	GOOS=windows GOARCH=amd64 go build -o $@ ./$(BTOTP_SRCDIR)

$(BTOTP_BINDIR)/$(BTOTP_APP_NAME)-linux-amd64: $(BTOTP_SRCDIR)/main.go | $(BTOTP_BINDIR)
	GOOS=linux   GOARCH=amd64 go build -o $@ ./$(BTOTP_SRCDIR)

$(BTOTP_BINDIR)/$(BTOTP_APP_NAME)-linux-arm: $(BTOTP_SRCDIR)/main.go | $(BTOTP_BINDIR)
	GOOS=linux   GOARCH=arm   go build -o $@ ./$(BTOTP_SRCDIR)

$(BTOTP_BINDIR)/$(BTOTP_APP_NAME)-linux-arm64: $(BTOTP_SRCDIR)/main.go | $(BTOTP_BINDIR)
	GOOS=linux   GOARCH=arm64 go build -o $@ ./$(BTOTP_SRCDIR)

$(BTOTP_BINDIR)/$(BTOTP_APP_NAME)-darwin-amd64: $(BTOTP_SRCDIR)/main.go | $(BTOTP_BINDIR)
	GOOS=darwin  GOARCH=amd64 go build -o $@ ./$(BTOTP_SRCDIR)

$(BTOTP_BINDIR)/$(BTOTP_APP_NAME)-darwin-arm64: $(BTOTP_SRCDIR)/main.go | $(BTOTP_BINDIR)
	GOOS=darwin  GOARCH=arm64 go build -o $@ ./$(BTOTP_SRCDIR)

clean_btotp:
	rm -rf $(BTOTP_BINDIR)/$(BTOTP_APP_NAME)-*
