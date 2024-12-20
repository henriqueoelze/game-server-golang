package gateways

import (
	"game-server-golang/internal/entities"

	"github.com/google/uuid"
)

type PlayerRepository interface {
	CreatePlayer(player entities.Player) error
	GetPlayer(publicId uuid.UUID) (entities.Player, error)
}
