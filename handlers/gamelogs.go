package handlers

import (
	"encoding/json"
	"fmt"
	"go-api/database"
	"go-api/models"
	"net/http"
)

func HandlerFunction(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT * FROM gamelogs")
	if err != nil {
		http.Error(w, fmt.Sprintf("Database query error: %v", err), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var logs []models.GameLog
	for rows.Next() {
		var log models.GameLog
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
