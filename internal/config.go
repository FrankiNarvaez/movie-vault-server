package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var TMDBApiKey string

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("[!!] Fail to load .env file")
	}

	TMDBApiKey = fmt.Sprintf("Bearer %s", os.Getenv("TMDB_API_KEY"))
}
