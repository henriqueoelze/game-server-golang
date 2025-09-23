package models

import (
	entities "game-server-golang/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model

	PublicId uuid.UUID
	Name     string
	Level    int
}

func (p Player) ToEntity() *entities.Player {
	return &entities.Player{
		PublicID: p.PublicId,
		Name:     p.Name,
		Level:    p.Level,
	}
}
