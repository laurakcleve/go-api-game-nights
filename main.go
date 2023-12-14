package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
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

var db *sql.DB
var once sync.Once
var dbInit sync.WaitGroup

func initDB() (*sql.DB, error) {
	fmt.Print(db)

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

	return db, nil
}

func GetDB() (*sql.DB, error) {
	var err error
	once.Do(func() {
		dbInit.Add(1)
		go func() {
			db, err = initDB()
			if err != nil {
				log.Fatalf("Database init error: %v", err)
				return
			}
			dbInit.Done()
		}()
	})

	dbInit.Wait()
	return db, err
}

func handlerFunction(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM gamelogs")
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
	var err error
	db, err = GetDB()
	if err != nil {
		log.Fatal("Failed to initialize the database: ", err)
		return
	}

	http.HandleFunc("/", handlerFunction)

	log.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}