package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var logger *log.Logger

func init() {
	envLogLevel, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
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
}

func GetLogger() *log.Logger {
	return logger
}

func getCustomFormatter() *log.TextFormatter {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05.000 MST"
	customFormatter.FullTimestamp = true
	return customFormatter
}
