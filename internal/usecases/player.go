package usecases

import (
	"game-server-golang/internal/entities"

	"github.com/google/uuid"
)

type PlayerUsecase interface {
	CreatePlayer() (newPlayer entities.Player, err error)
	GetPlayer(publicId uuid.UUID) (player entities.Player, err error)
}
