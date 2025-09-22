package player

import (
	"fmt"
	entities "game-server-golang/internal/domain"

	"github.com/google/uuid"
)

func (usecase *PlayerUsecaseImpl) CreatePlayer() (entities.Player, error) {
	newPlayer := entities.Player{
		PublicID: uuid.New(),
		Name:     "New Player",
		Level:    1,
	}

	err := usecase.playerRepository.CreatePlayer(newPlayer)
	if err != nil {
		return entities.Player{}, fmt.Errorf("error creating player: %w", err)
	}

	return newPlayer, nil
}
