// +build !windows,!nacl,!plan9

package utils

import (
	"io/ioutil"
	"log/syslog"
	"os"

	log "github.com/Sirupsen/logrus"
	logrusSyslog "github.com/Sirupsen/logrus/hooks/syslog"
)

// SwitchLoggingtoSystemLogger tells the logger to send the output to os specific logging system
func SwitchLoggingtoSystemLogger() {
	log.StandardLogger().Hooks = make(log.LevelHooks)
	hook, err := logrusSyslog.NewSyslogHook("", "", syslog.LOG_WARNING, "")
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
