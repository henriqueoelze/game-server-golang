package models

import (
	"game-server-golang/internal/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	PublicId uuid.UUID
	Name     string
	Level    int
}

func (p Player) ToEntity() entities.Player {
	return entities.Player{
		PublicId: p.PublicId,
		Name:     p.Name,
		Level:    p.Level,
	}
}
