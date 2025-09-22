package player

import (
	entities "game-server-golang/internal/domain"

	"github.com/google/uuid"
)

func (usecase *PlayerUsecaseImpl) CreatePlayer() (entities.Player, error) {
	newPlayer := entities.Player{
		PublicId: uuid.New(),
		Name:     "New Player",
		Level:    1,
	}

	err := usecase.playerRepository.CreatePlayer(newPlayer)
	if err != nil {
		return entities.Player{}, err
	}

	return newPlayer, nil
}
