package database

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var db *sqlx.DB

func Connect() *sqlx.DB {
	godotenv.Load()
	database_url := os.Getenv("DATABASE_URL")
	var err error
	db, err = sqlx.Connect("postgres", database_url)
	if err != nil {
		log.Fatalf("[!!] Ha ocurrido un error al conectar a la base de datos: %v", err)
	}

	return db
}
