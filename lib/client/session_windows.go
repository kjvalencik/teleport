// +build windows

package client

import (
	"io"
	"os"
)

func watchWindowChange(sigC chan os.Signal) {
	// this isn't supported on windows. do nothing
}

func catchTerminalStopSignal(shell io.Writer) {
	// Catch Ctrl-Z signal. This isn't supported on windows. do nothing
}
