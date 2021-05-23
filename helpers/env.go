package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvDefault(key, defaultValue string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}
