package database

import (
	"log"
	database "movie/db"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Seeds(db *sqlx.DB) {
	log.Println("[!] Inserting seeds...")

	if _, err := db.Exec(database.USERS); err != nil {
		log.Printf("[!!] Error insert seeds: %v", err)
		return
	}
	log.Println("[!] Seeds inserted...")
}
