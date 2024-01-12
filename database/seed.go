package database

import (
	"go-api/models"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	games := []models.Game{
		{Name: "Nemesis"},
		{Name: "Kingdom Death"},
		{Name: "Wingspan"},
	}

	for _, game := range games {
		var count int64
		db.Model(&models.Game{}).Where("name = ?", game.Name).Count(&count)
		if count == 0 {
			if err := db.Create(&game).Error; err != nil {
				return err
			}
		}
	}

	players := []models.Player{
		{Name: "Jimmy"},
		{Name: "Laura"},
		{Name: "Ylish"},
	}

	for _, player := range players {
		var count int64
		db.Model(&models.Player{}).Where("name = ?", player.Name).Count(&count)
		if count == 0 {
			if err := db.Create(&player).Error; err != nil {
				return err
			}
		}
	}

	playedGames := []models.PlayedGame{
		{Date: "2023-12-13 12:00:00", GameID: 1, WinnerID: 2},
		{Date: "2023-12-14 15:30:00", GameID: 2, WinnerID: 3},
		{Date: "2023-12-15 18:45:00", GameID: 3, WinnerID: 1},
	}

	for _, playedGame := range playedGames {
		var count int64
		db.Model(&models.PlayedGame{}).Where("date = ?", playedGame.Date).Count(&count)
		if count == 0 {
			if err := db.Create(&playedGame).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
