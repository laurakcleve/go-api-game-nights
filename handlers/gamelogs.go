package handlers

import (
	"encoding/json"
	"fmt"
	"go-api/database"
	"go-api/models"
	"net/http"
)

func GetAllPlayedGames(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT * FROM played_game")
	if err != nil {
		http.Error(w, fmt.Sprintf("Database query error: %v", err), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var logs []models.PlayedGame
	for rows.Next() {
		var log models.PlayedGame
		if err := rows.Scan(&log.ID, &log.Date, &log.GameID, &log.WinnerID); err != nil {
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

func GetAllGames(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT * FROM game")
	if err != nil {
		http.Error(w, fmt.Sprintf("Database query error: %v", err), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var games []models.Game
	for rows.Next() {
		var game models.Game
		if err := rows.Scan(&game.ID, &game.Name); err != nil {
			http.Error(w, fmt.Sprintf("Database scan error: %v", err), http.StatusInternalServerError)
			return
		}
		games = append(games, game)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Database rows error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(games)
	if err != nil {
		http.Error(w, fmt.Sprintf("JSON encoding error, %v", err), http.StatusInternalServerError)
	}
}

func GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT * FROM player")
	if err != nil {
		http.Error(w, fmt.Sprintf("Database query error: %v", err), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var players []models.Player
	for rows.Next() {
		var player models.Player
		if err := rows.Scan(&player.ID, &player.Name); err != nil {
			http.Error(w, fmt.Sprintf("Database scan error: %v", err), http.StatusInternalServerError)
			return
		}
		players = append(players, player)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Database rows error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(players)
	if err != nil {
		http.Error(w, fmt.Sprintf("JSON encoding error, %v", err), http.StatusInternalServerError)
	}
}
