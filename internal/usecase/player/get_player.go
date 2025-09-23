package player

import (
	"fmt"
	"game-server-golang/internal/domain"

	"github.com/google/uuid"
)

func (usecase *PlayerUsecaseImpl) GetPlayer(publicId uuid.UUID) (*domain.Player, error) {
	player, err := usecase.playerRepository.GetPlayer(publicId)
	if err != nil {
		return player, fmt.Errorf("error getting player: %w", err)
	}

	return player, nil
}
