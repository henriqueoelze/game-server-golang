package player

import (
	"game-server-golang/internal/gateway"
	"game-server-golang/internal/usecase"
)

var _ = usecase.PlayerUsecase(&PlayerUsecaseImpl{})

type PlayerUsecaseImpl struct {
	playerRepository gateway.PlayerRepository
}

func NewPlayerUsecase(
	playerRepository gateway.PlayerRepository,
) usecase.PlayerUsecase {
	return &PlayerUsecaseImpl{
		playerRepository: playerRepository,
	}
}
