// Command ipcalc calculates IPv4 and IPv6 network information.
package main

import (
	"os"

	"github.com/cleonte/go-ipcalc/internal/ipcalc"
)

func main() { os.Exit(ipcalc.Main()) }
