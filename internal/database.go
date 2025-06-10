package config

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func ConnectDB() *sqlx.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dsn :=
		"host=" + host +
			" port=" + port +
			" user=" + user +
			" password=" + password +
			" dbname=" + dbname +
			" sslmode=disable"

	var err error
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("[!!] Ha ocurrido un error al conectar a la base de datos: %v", err)
	}

	return db
}
