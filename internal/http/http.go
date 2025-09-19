package http

import (
	"context"
	"game-server-golang/internal/constant"
	"game-server-golang/internal/core"
	"game-server-golang/internal/gateway"
	"game-server-golang/internal/http/middleware"
	"game-server-golang/internal/usecase"
	"net/http"

	"github.com/google/uuid"
)

type MetagameApi struct {
	securityUsecase usecase.SecurityUsecase
	playerUsecase   usecase.PlayerUsecase

	baseLogger gateway.Logger

	core.ContextActions
}

func NewMetagameApi(
	securityUsecase usecase.SecurityUsecase,
	playerUsecase usecase.PlayerUsecase,

	baseLogger gateway.Logger,
) *MetagameApi {
	return &MetagameApi{
		securityUsecase: securityUsecase,
		playerUsecase:   playerUsecase,
		baseLogger:      baseLogger,
	}
}

func (api *MetagameApi) Start(serverAddress string) {
	openRouter := http.NewServeMux()
	openRouter.HandleFunc(HealthPath, api.Health)
	openRouter.HandleFunc(CreatePlayerPath, api.CreatePlayer)

	playerRouter := http.NewServeMux()
	playerRouter.HandleFunc(GetPlayerPath, api.GetPlayer)

	authenticationMiddleware := middleware.NewAuthenticationMiddleware(api.securityUsecase, api.baseLogger)

	openRouter.Handle("/", authenticationMiddleware.Authenticate(playerRouter))

	err := http.ListenAndServe(serverAddress, openRouter)
	if err != nil {
		panic(err)
	}
}

func (api *MetagameApi) GetPlayerIdFromCtx(ctx context.Context) uuid.UUID {
	playerId := ctx.Value(constant.ContextKeyPlayerID).(uuid.UUID)
	return playerId
}
