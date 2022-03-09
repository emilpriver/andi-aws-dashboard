package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadSecrets() {
	err := godotenv.Load()

	if err != nil {
		log.Panic("Error loading .env file")
	}
}
