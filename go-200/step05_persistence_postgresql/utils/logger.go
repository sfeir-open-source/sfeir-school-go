package utils

import (
	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	// AppName is the application's name
	AppName  = "todolist"
	AppField = "app"

	// LogStashFormatter is constant used to format logs as logstash format
	LogStashFormatter = "logstash"
	// TextFormatter is constant used to format logs as simple text format
	TextFormatter = "text"
)

// InitLog initializes the logrus logger
func InitLog(logLevel, formatter string) error {

	switch formatter {
	case LogStashFormatter:
		logrus.SetFormatter(logrustash.DefaultFormatter(
			logrus.Fields{AppField: AppName}))
	default:
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
	}

	logrus.SetOutput(os.Stdout)

	level, err := logrus.ParseLevel(logLevel)

	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
		return err
	}

	logrus.SetLevel(level)
	return nil
}
