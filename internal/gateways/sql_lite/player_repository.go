package sql_lite

import (
	"game-server-golang/internal/entities"
	"game-server-golang/internal/gateways"
	"game-server-golang/internal/models"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type PlayerRepositoryImpl struct {
	db *gorm.DB
}

var _ gateways.PlayerRepository = &PlayerRepositoryImpl{}

func NewPlayerRepositoryImpl() (*PlayerRepositoryImpl, error) {
	db, err := gorm.Open(sqlite.Open("in_memory_db.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &PlayerRepositoryImpl{
		db: db,
	}, nil
}

func (dal *PlayerRepositoryImpl) CreatePlayer(player entities.Player) error {
	playerModel := models.Player{
		PublicId: player.PublicId,
		Name:     player.Name,
		Level:    player.Level,
	}

	result := dal.db.Create(&playerModel)
	if result.Error != nil {
		return result.Error
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
