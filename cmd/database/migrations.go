package database

import (
	"log"
	database "movie/db"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func hasTables(db *sqlx.DB) bool {
	var count int
	query := `
		SELECT COUNT(*) 
		FROM information_schema.tables 
		WHERE table_schema = 'public';
	`

	if err := db.Get(&count, query); err != nil {
		log.Fatalf("[!!] Error to verify the tables amount: %v", err)
	}

	return count > 0
}

func Migrate(db *sqlx.DB) {
	if hasTables(db) {
		log.Println("[!] dropping existing tables...")

		if _, err := db.Exec(database.DROP_SCHEMA); err != nil {
			log.Printf("[!!] Error to drop tables: %v", err)
		}
	}

	log.Println("[!] creating tables...")

	if _, err := db.Exec(database.SCHEMA); err != nil {
		log.Printf("[!!] Error to create tables: %v", err)
		return
	}
	log.Println("[!] Tables created...")
}
