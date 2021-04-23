package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetEnvVariable from the env file
func GetEnvVariable(key string) (string, error) {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
		return "", err
	}

	return os.Getenv(key), nil
}