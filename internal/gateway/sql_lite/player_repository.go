package sql_lite

import (
	"fmt"
	entities "game-server-golang/internal/domain"
	"game-server-golang/internal/gateway"
	"game-server-golang/internal/gateway/sql_lite/models"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type PlayerRepositoryImpl struct {
	db *gorm.DB
}

var _ gateway.PlayerRepository = &PlayerRepositoryImpl{}

func NewPlayerRepositoryImpl(databaseName string) (*PlayerRepositoryImpl, error) {
	db, err := gorm.Open(sqlite.Open(databaseName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	return &PlayerRepositoryImpl{
		db: db,
	}, nil
}

func (dal *PlayerRepositoryImpl) CreatePlayer(player entities.Player) error {
	playerModel := models.Player{
		PublicId: player.PublicID,
		Name:     player.Name,
		Level:    player.Level,
	}

	result := dal.db.Create(&playerModel)
	if result.Error != nil {
		return fmt.Errorf("error creating player: %w", result.Error)
	}

	return nil
}

func (dal *PlayerRepositoryImpl) GetPlayer(publicId uuid.UUID) (entities.Player, error) {
	var playerModel models.Player
	result := dal.db.First(&playerModel, "public_id = ?", publicId)
	if result.Error != nil {
		return entities.Player{}, result.Error
	}

	return playerModel.ToEntity(), nil
}
