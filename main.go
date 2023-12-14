package main

import (
	"encoding/json"
	"fmt"
	"go-api/database"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type GameLog struct {
	ID     int
	Date   time.Time
	Game   string
	Winner string
}

func handlerFunction(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT * FROM gamelogs")
	if err != nil {
		http.Error(w, fmt.Sprintf("Database query error: %v", err), http.StatusInternalServerError)
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

}

func main() {
	database.InitDB()

	http.HandleFunc("/", handlerFunction)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
