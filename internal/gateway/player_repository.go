package gateway

import (
	entities "game-server-golang/internal/domain"

	"github.com/google/uuid"
)

type PlayerRepository interface {
	CreatePlayer(player entities.Player) error
	GetPlayer(publicId uuid.UUID) (*entities.Player, error)
}
