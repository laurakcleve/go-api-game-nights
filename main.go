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

	http.HandleFunc("/", handlers.HandlerFunction)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
