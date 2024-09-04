package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken string
	GithubToken string
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
	return Config{
		TelegramToken: os.Getenv("TOKEN"),
		GithubToken: os.Getenv("GITHUB_TOKEN"),
	}
}