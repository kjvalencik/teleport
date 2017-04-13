// +build windows

package utils

import (
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
	logrusEventlog "github.com/Sirupsen/logrus/hooks/eventlog"
)

// SwitchLoggingtoSystemLogger tells the logger to send the output to os specific logging system
func SwitchLoggingtoSystemLogger() {
	log.StandardLogger().Hooks = make(log.LevelHooks)
	hook, err := logrusEventlog.NewEventlogHook("Teleport tsh")
	if err != nil {
		// syslog not available
		log.SetOutput(os.Stderr)
		log.Warn("syslog not available. reverting to stderr")
	} else {
		// ... and disable stderr:
		log.AddHook(hook)
		log.SetOutput(ioutil.Discard)
	}
}
