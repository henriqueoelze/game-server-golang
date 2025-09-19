package main

import (
	"fmt"
	"game-server-golang/internal/config"
	"game-server-golang/internal/gateway/sql_lite/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Error loading config: %v", err))
	}

	db, err := gorm.Open(sqlite.Open(config.Database.Name), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(models.Player{})
}
