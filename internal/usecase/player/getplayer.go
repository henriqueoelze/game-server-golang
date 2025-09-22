package player

import (
	"fmt"
	entities "game-server-golang/internal/domain"

	"github.com/google/uuid"
)

func (usecase *PlayerUsecaseImpl) GetPlayer(publicId uuid.UUID) (entities.Player, error) {
	player, err := usecase.playerRepository.GetPlayer(publicId)
	if err != nil {
		return entities.Player{}, fmt.Errorf("error getting player: %w", err)
	}

	return player, nil
}
