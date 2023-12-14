package models

import "time"

type PlayedGame struct {
	ID       int       `json:"id"`
	Date     time.Time `json:"date"`
	GameID   int       `json:"game_id"`
	WinnerID int       `json:"winner_id"`
}

type Game struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Player struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PlayedGamePlayer struct {
	PlayedGameID int `json:"played_game_id"`
	PlayerID     int `json:"player_id"`
}
