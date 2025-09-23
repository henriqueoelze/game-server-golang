package models

import (
	"game-server-golang/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model

	PublicId uuid.UUID
	Name     string
	Level    int
}

func (p Player) ToDomain() *domain.Player {
	return &domain.Player{
		PublicID: p.PublicId,
		Name:     p.Name,
		Level:    p.Level,
	}
}
