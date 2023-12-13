package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM gamelogs")
		if err != nil {
			http.Error(w, "Database query error", http.StatusInternalServerError)
			return
		}

		defer rows.Close()

		var logs []GameLog
		for rows.Next() {
			var log GameLog
			if err := rows.Scan(&log.ID, &log.Date, &log.Game, &log.Winner); err != nil {
				http.Error(w, fmt.Sprintf("Database scan error: %v", err), http.StatusInternalServerError)
				return
			}
			logs = append(logs, log)
		}

		if err = rows.Err(); err != nil {
			http.Error(w, fmt.Sprintf("Database rows error: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(logs)
		if err != nil {
			http.Error(w, fmt.Sprintf("JSON encoding error, %v", err), http.StatusInternalServerError)
		}

	})

	log.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
