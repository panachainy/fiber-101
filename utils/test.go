package utils

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func SetupTest() {
	err := godotenv.Load(".env.test")
	if err != nil {
		logrus.Errorln("Error loading .env.test file:", err)
	}
}
