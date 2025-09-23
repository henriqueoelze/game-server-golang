package gateway

import (
	"game-server-golang/internal/domain"

	"github.com/google/uuid"
)

type PlayerRepository interface {
	CreatePlayer(player domain.Player) error
	GetPlayer(publicId uuid.UUID) (*domain.Player, error)
}
