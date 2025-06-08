package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var TMDBApiKey string

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("[!!] Fail to load .env file")
	}

	TMDBApiKey = os.Getenv("TMDB_API_KEY")
}
