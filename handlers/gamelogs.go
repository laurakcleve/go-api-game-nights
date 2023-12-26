package handlers

import (
	"fmt"
	"go-api/database"
	"go-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPlayedGames(c *gin.Context) {
	rows, err := database.DB.Query("SELECT * FROM played_game")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	var playedGames []models.PlayedGame
	for rows.Next() {
		var log models.PlayedGame
		if err := rows.Scan(&log.ID, &log.Date, &log.GameID, &log.WinnerID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		playedGames = append(playedGames, log)
	}

	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, playedGames)
}

func GetAllGames(c *gin.Context) {
	rows, err := database.DB.Query("SELECT * FROM game")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	var games []models.Game
	for rows.Next() {
		var game models.Game
		if err := rows.Scan(&game.ID, &game.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		games = append(games, game)
	}

	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, games)
}

func GetAllPlayers(c *gin.Context) {
	rows, err := database.DB.Query("SELECT * FROM player")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	var players []models.Player
	for rows.Next() {
		var player models.Player
		if err := rows.Scan(&player.ID, &player.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		players = append(players, player)
	}

	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, players)
}

func AddPlayedGame(c *gin.Context) {
	date := c.PostForm("date")
	game_id := c.PostForm("game_id")
	winner_id := c.PostForm("winner_id")

	fmt.Println(date)
	fmt.Println(game_id)
	fmt.Println(winner_id)

	stmt, err := database.DB.Prepare(
		"INSERT INTO played_game(date, game_id, winner_id) VALUES(?, ?, ?) RETURNING *")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	defer stmt.Close()

	result, err := stmt.Query(date, game_id, winner_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	fmt.Println(result)
	c.JSON(http.StatusOK, result)
}
