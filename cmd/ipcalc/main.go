// SPDX-License-Identifier: GPL-2.0-only
//
// Copyright (c) 1997-2015 Red Hat, Inc.
// Copyright (c) 2026 go-ipcalc contributors
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License, version 2,
// as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// Command ipcalc calculates IPv4 and IPv6 network information.
package main

import (
	"os"

	"github.com/cleonte/go-ipcalc/internal/ipcalc"
)

func main() { os.Exit(ipcalc.Main()) }
