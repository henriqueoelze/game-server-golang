package usecase

import (
	"game-server-golang/internal/domain"

	"github.com/google/uuid"
)

type PlayerUsecase interface {
	CreatePlayer() (newPlayer domain.Player, err error)
	GetPlayer(publicId uuid.UUID) (player *domain.Player, err error)
}
