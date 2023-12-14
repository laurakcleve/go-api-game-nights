package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB() error {
	const file string = "gamelogs.db"

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	query := `
		  DROP TABLE gamelogs;
			CREATE TABLE gamelogs (
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

	DB = db
	return nil
}
