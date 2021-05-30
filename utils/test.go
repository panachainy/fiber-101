package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

func SetupTest() {
	err := godotenv.Load(".env.test")
	if err != nil {
		fmt.Println("Error loading .env.test file:", err)
	}
}
