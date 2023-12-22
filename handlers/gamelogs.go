package handlers

import (
	"encoding/json"
	"fmt"
	"go-api/database"
	"go-api/models"
	"net/http"
	"time"
)

func GetAllPlayedGames(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
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

	case http.MethodPost:
		var inputData models.PlayedGameInput
		err := json.NewDecoder(r.Body).Decode(&inputData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		fmt.Printf("Input data: %v", inputData)

		parsedDate, err := time.Parse(time.RFC3339, inputData.Date)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing date: %v", err), http.StatusInternalServerError)
		}
		fmt.Printf("Parsed date: %v", parsedDate)

		stmt, err := database.DB.Prepare(
			"INSERT INTO played_game(date, game_id, winner_id) VALUES(?, ?, ?)")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error preparing db query: %v", err), http.StatusInternalServerError)
		}
		defer stmt.Close()

		_, err = stmt.Exec(inputData.Date, inputData.GameID, inputData.WinnerID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error executing db insert: %v", err), http.StatusInternalServerError)
		}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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
