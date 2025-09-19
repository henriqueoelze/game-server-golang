package player

import (
	base "game-server-golang/internal/core"
	entities "game-server-golang/internal/domain"
	"game-server-golang/internal/gateways"
	"game-server-golang/internal/usecases"

	"github.com/google/uuid"
)

var _ = usecases.PlayerUsecase(&PlayerUsecaseImpl{})

type PlayerUsecaseImpl struct {
	base.BaseLogger
	playerRepository gateways.PlayerRepository
}

func NewPlayerUsecase(
	playerRepository gateways.PlayerRepository,
) usecases.PlayerUsecase {
	return &PlayerUsecaseImpl{
		playerRepository: playerRepository,
	}
}

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

func (usecase *PlayerUsecaseImpl) GetPlayer(publicId uuid.UUID) (entities.Player, error) {
	player, err := usecase.playerRepository.GetPlayer(publicId)
	if err != nil {
		return entities.Player{}, err
	}

	return player, nil
}
