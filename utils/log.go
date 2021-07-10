package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

func InitLogrus() {
	logLevel, ok := os.LookupEnv("LOG_LEVEL")

	// LOG_LEVEL not set, let's default to debug
	if !ok {
		logLevel = "debug"
	}

	// parse string, this is built-in feature of logrus
	logrusLevel, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrusLevel = logrus.DebugLevel
	}

	// set global log level
	logrus.SetLevel(logrusLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)

	logrus.Infoln("Log level: ", logrus.GetLevel())
	logrus.New()
}
