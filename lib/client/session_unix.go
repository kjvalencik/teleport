// +build !windows

package client

import (
	"io"
	"os"
	"os/signal"
	"syscall"

	log "github.com/Sirupsen/logrus"
)

func watchWindowChange(sigC chan os.Signal) {
	signal.Notify(sigC, syscall.SIGWINCH)
}

func catchTerminalStopSignal(shell io.Writer) {
	// Catch Ctrl-Z signal
	ctrlZSignal := make(chan os.Signal, 1)
	signal.Notify(ctrlZSignal, syscall.SIGTSTP)
	go func() {
		for {
			<-ctrlZSignal
			_, err := shell.Write([]byte{26})
			if err != nil {
				log.Errorf(err.Error())
			}
		}
	}()
}
