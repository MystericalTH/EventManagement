package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, relying on system environment variables")
	}

	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}
