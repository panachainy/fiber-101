package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func SetupTest() {
	err := godotenv.Load(".env.test")
	if err != nil {
		log.Println("Error loading .env.test file:", err)
	}
}
