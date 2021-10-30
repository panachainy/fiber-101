package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func Get(key string, envPath ...string) string {
	path := ".env"

	if len(envPath) > 0 {
		path = envPath[0]
	}

	err := godotenv.Load(path)
	if err != nil {
		logrus.Debugln("[CONFIG] Error loading .env file:", err)
	}

	return os.Getenv(key)
}
