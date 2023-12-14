package models

import "time"

type GameLog struct {
	ID     int
	Date   time.Time
	Game   string
	Winner string
}
