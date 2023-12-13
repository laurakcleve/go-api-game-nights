package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type GameLog struct {
	ID      int
	Date    time.Time
	Game    string
	Players string
	Winner  string
}

func main() {
	const file string = "gamelogs.db"

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	query := `
			CREATE TABLE IF NOT EXISTS gamelogs (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				date DATETIME,
				game TEXT,
				winner TEXT
			);
		`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	add_query := `
		INSERT INTO gamelogs (date, game, winner)
		VALUES ('2023-12-13 12:00:00', 'Nemesis', 'Laura')
	`

	_, err = db.Exec(add_query)
	if err != nil {
		log.Fatalf("Failed to add row: %v", err)
	}
}
