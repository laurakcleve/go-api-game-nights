package main

import (
	"go-api/database"
	"go-api/handlers"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database.InitDB()

	http.HandleFunc("/", handlers.GetAllPlayedGames)
	http.HandleFunc("/games", handlers.GetAllGames)
	http.HandleFunc("/players", handlers.GetAllPlayers)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
