package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Impure
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// Impure
func GetWeatherAPIKey() string {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		log.Fatalf("WEATHER_API_KEY is not set in the environment variables")
	}
	return apiKey
}
