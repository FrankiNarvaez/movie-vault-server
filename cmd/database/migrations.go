package database

import (
	"log"
	database "movie/db"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func dropAllTables(db *sqlx.DB) error {
	log.Println("[!] Dropping entire public schema...")

	dropSchemaQuery := `
		DROP SCHEMA public CASCADE;
		CREATE SCHEMA public;
	`

	if _, err := db.Exec(dropSchemaQuery); err != nil {
		return err
	}

	log.Println("[✓] Public schema dropped and recreated successfully")
	return nil
}

func createTables(db *sqlx.DB) error {
	log.Println("[!] Creating tables...")

	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if _, err = tx.Exec(database.SCHEMA); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	log.Println("[✓] Tables created successfully")
	return nil
}

func Migrate(db *sqlx.DB) {
	// Directamente elimina todo el schema público
	if err := dropAllTables(db); err != nil {
		log.Fatalf("[!!] Error dropping schema: %v", err)
	}

	// Crea tablas
	if err := createTables(db); err != nil {
		log.Fatalf("[!!] Error creating tables: %v", err)
	}

	log.Println("[✓] Migration completed successfully")
}
