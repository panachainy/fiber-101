package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Get(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	return os.Getenv(key)
}
