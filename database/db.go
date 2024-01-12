package database

import (
	"go-api/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	const file string = "gamelogs.db"

	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err := db.Migrator().DropTable(
		&models.PlayedGame{},
		&models.Game{},
		&models.Player{},
		&models.PlayedGamePlayer{}); err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}

	err = db.AutoMigrate(
		&models.PlayedGame{},
		&models.Game{},
		&models.Player{},
		&models.PlayedGamePlayer{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	err = Seed(db)
	if err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	DB = db
	return nil
}
