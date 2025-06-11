package config

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func ConnectDB() *sqlx.DB {
	database_url := os.Getenv("DATABASE_URL")

	var err error
	db, err = sqlx.Connect("postgres", database_url)
	if err != nil {
		log.Fatalf("[!!] Ha ocurrido un error al conectar a la base de datos: %v", err)
	}

	return db
}
