// +build windows

package logrus_eventlog

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"
	"golang.org/x/sys/windows/svc/eventlog"
)

// EventlogHook to send logs via syslog.
type EventlogHook struct {
	logger *eventlog.Log
}

const defaultSupportedEvent uint32 = eventlog.Error | eventlog.Warning | eventlog.Info

func NewEventlogHook(name string) (*EventlogHook, error) {
	err := eventlog.InstallAsEventCreate(name, defaultSupportedEvent)
	if err != nil && !strings.HasSuffix(err.Error(), "registry key already exists") {
		return nil, err
	}

	l, err := eventlog.Open(name)
	if err != nil {
		return nil, err
	}

	h := &EventlogHook{logger: l}
	runtime.SetFinalizer(h, func(hook *EventlogHook) {
		hook.logger.Close()
	})

	return h, nil
}

func (hook *EventlogHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}

	switch entry.Level {
	case logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel:
		return hook.logger.Error(uint32(entry.Level), line)
	case logrus.WarnLevel:
		return hook.logger.Warning(uint32(entry.Level), line)
	case logrus.InfoLevel, logrus.DebugLevel:
		return hook.logger.Info(uint32(entry.Level), line)
	}

	return nil
}

func (hook *EventlogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
