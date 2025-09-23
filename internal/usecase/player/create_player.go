package player

import (
	"fmt"
	"game-server-golang/internal/domain"

	"github.com/google/uuid"
)

func (usecase *PlayerUsecaseImpl) CreatePlayer() (domain.Player, error) {
	newPlayer := domain.Player{
		PublicID: uuid.New(),
		Name:     "New Player",
		Level:    1,
	}

	err := usecase.playerRepository.CreatePlayer(newPlayer)
	if err != nil {
		return domain.Player{}, fmt.Errorf("error creating player: %w", err)
	}

	return newPlayer, nil
}
