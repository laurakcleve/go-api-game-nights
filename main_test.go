package main

import (
	"encoding/json"
	"go-api/database"
	"go-api/handlers"
	"go-api/models"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func init() {
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
}

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handlers.HandlerFunction(w, req)

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
		Date:     time.Date(2023, 12, 13, 12, 0, 0, 0, time.UTC),
		GameID:   1,
		WinnerID: 2,
	}

	if responseStruct.ID != expected.ID {
		t.Errorf("unexpected value for ID: got %v want %v", responseStruct.ID, expected.ID)
	}

	if !responseStruct.Date.Equal(expected.Date) {
		t.Errorf("unexpected value for Date: got %v want %v", responseStruct.Date, expected.Date)
	}

	if responseStruct.GameID != expected.GameID {
		t.Errorf("unexpected value for Game: got %v want %v", responseStruct.GameID, expected.GameID)
	}

	if responseStruct.WinnerID != expected.WinnerID {
		t.Errorf("unexpected value for Winner: got %v want %v", responseStruct.WinnerID, expected.WinnerID)
	}

}
