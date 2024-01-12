package handlers

import (
	"go-api/database"
	"go-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPlayedGames(c *gin.Context) {
	var playedGames []models.PlayedGame
	_ = database.DB.Find(&playedGames)
	c.JSON(http.StatusOK, playedGames)
}

func GetAllGames(c *gin.Context) {
	var games []models.Game
	_ = database.DB.Find(&games)
	c.JSON(http.StatusOK, games)
}

func GetAllPlayers(c *gin.Context) {
	var players []models.Player
	_ = database.DB.Find(&players)
	c.JSON(http.StatusOK, players)
}

func AddPlayedGame(c *gin.Context) {
	date := c.PostForm("date")
	gameIdStr := c.PostForm("game_id")
	winnerIdStr := c.PostForm("winner_id")

	gameId, err := strconv.Atoi(gameIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	winnerId, err := strconv.Atoi(winnerIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid winner ID"})
		return
	}

	newPlayedGame := models.PlayedGame{
		Date: date, GameID: gameId, WinnerID: winnerId}

	result := database.DB.Create(&newPlayedGame)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, newPlayedGame)
}
