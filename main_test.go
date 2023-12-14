package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	var err error
	db, err = GetDB()
	if err != nil {
		fmt.Println("Failed to initialize the database:", err)
		return
	}

	handlerFunction(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseArray []GameLog
	if err := json.NewDecoder(rr.Body).Decode(&responseArray); err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

	if len(responseArray) != 1 {
		t.Errorf("unexpected number of elements in the array: got %v want %v", len(responseArray), 1)
	}
	responseStruct := responseArray[0]

	expected := GameLog{
		ID:      1,
		Date:    time.Date(2023, 12, 13, 12, 0, 0, 0, time.UTC),
		Game:    "Nemesis",
		Players: "",
		Winner:  "Laura",
	}

	if responseStruct.ID != expected.ID {
		t.Errorf("unexpected value for ID: got %v want %v", responseStruct.ID, expected.ID)
	}

	if !responseStruct.Date.Equal(expected.Date) {
		t.Errorf("unexpected value for Date: got %v want %v", responseStruct.Date, expected.Date)
	}

	if responseStruct.Game != expected.Game {
		t.Errorf("unexpected value for Game: got %v want %v", responseStruct.Game, expected.Game)
	}

	if responseStruct.Players != expected.Players {
		t.Errorf("unexpected value for Players: got %v want %v", responseStruct.Players, expected.Players)
	}

	if responseStruct.Winner != expected.Winner {
		t.Errorf("unexpected value for Winner: got %v want %v", responseStruct.Winner, expected.Winner)
	}

}
