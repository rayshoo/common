package common

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var logger *log.Logger

var Fields log.Fields

func GetLogger(envLogLevel string) *log.Logger {
	if envLogLevel == "" {
		envLogLevel = "InfoLevel"
	}
	logLevel, err := log.ParseLevel(envLogLevel)
	if err != nil {
		logLevel = log.InfoLevel
	}

	logger = &log.Logger{
		Out:       os.Stderr,
		Formatter: getCustomFormatter(),
		Level:     logLevel,
	}
	return logger
}

func getCustomFormatter() *log.TextFormatter {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05.000 MST"
	customFormatter.FullTimestamp = true
	return customFormatter
}
