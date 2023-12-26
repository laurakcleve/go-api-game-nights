package main

import (
	"encoding/json"
	"go-api/database"
	"go-api/handlers"
	"go-api/models"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", handlers.GetAllPlayedGames)
	router.POST("/", handlers.AddPlayedGame)
	router.GET("/games", handlers.GetAllGames)
	router.GET("/players", handlers.GetAllPlayers)

	return router
}

func TestGetAllPlayedGames(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseArray []models.PlayedGame
	if err := json.NewDecoder(w.Body).Decode(&responseArray); err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

	if len(responseArray) != 3 {
		t.Errorf("unexpected number of elements in the array: got %v want %v", len(responseArray), 3)
	}
	responseStruct := responseArray[0]

	expected := models.PlayedGame{
		ID:       1,
		Date:     "2023-12-13T12:00:00Z",
		GameID:   1,
		WinnerID: 2,
	}

	if responseStruct.ID != expected.ID {
		t.Errorf("unexpected value for ID: got %v want %v", responseStruct.ID, expected.ID)
	}

	if responseStruct.Date != expected.Date {
		t.Errorf("unexpected value for Date: got %v want %v", responseStruct.Date, expected.Date)
	}

	if responseStruct.GameID != expected.GameID {
		t.Errorf("unexpected value for Game: got %v want %v", responseStruct.GameID, expected.GameID)
	}

	if responseStruct.WinnerID != expected.WinnerID {
		t.Errorf("unexpected value for Winner: got %v want %v", responseStruct.WinnerID, expected.WinnerID)
	}
}

func TestGetAllGames(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/games", nil)
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseArray []models.Game
	if err := json.NewDecoder(w.Body).Decode(&responseArray); err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

	if len(responseArray) != 3 {
		t.Errorf("unexpected number of elements in the array: got %v want %v", len(responseArray), 3)
	}
	responseStruct := responseArray[0]

	expected := models.Game{
		ID:   1,
		Name: "Nemesis",
	}

	if responseStruct.ID != expected.ID {
		t.Errorf("unexpected value for ID: got %v want %v", responseStruct.ID, expected.ID)
	}

	if responseStruct.Name != expected.Name {
		t.Errorf("unexpected value for Name: got %v want %v", responseStruct.Name, expected.Name)
	}
}

func TestGetAllPlayers(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/players", nil)
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseArray []models.Player
	if err := json.NewDecoder(w.Body).Decode(&responseArray); err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

	if len(responseArray) != 3 {
		t.Errorf("unexpected number of elements in the array: got %v want %v", len(responseArray), 3)
	}
	responseStruct := responseArray[0]

	expected := models.Player{
		ID:   1,
		Name: "Jimmy",
	}

	if responseStruct.ID != expected.ID {
		t.Errorf("unexpected value for ID: got %v want %v", responseStruct.ID, expected.ID)
	}

	if responseStruct.Name != expected.Name {
		t.Errorf("unexpected value for Name: got %v want %v", responseStruct.Name, expected.Name)
	}
}

func TestAddPlayedGame(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	reqBody := "date=2023-12-12T12:00:00-07:00&game_id=1&winner_id=1"
	req, _ := http.NewRequest("POST", "/", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
