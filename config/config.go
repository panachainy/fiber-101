package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func Get(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Errorln("Error loading .env file:", err)
	}

	return os.Getenv(key)
}
