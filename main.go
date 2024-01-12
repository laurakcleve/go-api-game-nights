package main

import (
	"go-api/database"
	"go-api/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database.InitDB()

	router := gin.Default()

	router.GET("/", handlers.GetAllPlayedGames)
	router.POST("/", handlers.AddPlayedGame)
	router.GET("/games", handlers.GetAllGames)
	router.GET("/players", handlers.GetAllPlayers)

	router.Run()
}
