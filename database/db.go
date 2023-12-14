package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
)

var DB *sql.DB

func InitDB() error {
	const file string = "gamelogs.db"

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, "database/seed.sql")

	queryBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to open seed file: %v", err)
	}

	query := string(queryBytes)

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}

	DB = db
	return nil
}
