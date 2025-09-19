package http

import (
	"context"
	"game-server-golang/internal/constants"
	base "game-server-golang/internal/core"
	"game-server-golang/internal/http/middlewares"
	"game-server-golang/internal/usecases"
	"net/http"

	"github.com/google/uuid"
)

type MetagameApi struct {
	base.BaseLogger
	base.BaseSession
	securityUsecase usecases.SecurityUsecase
	playerUsecase   usecases.PlayerUsecase
}

func NewMetagameApi(
	securityUsecase usecases.SecurityUsecase,
	playerUsecase usecases.PlayerUsecase,
) *MetagameApi {
	return &MetagameApi{
		securityUsecase: securityUsecase,
		playerUsecase:   playerUsecase,
	}
}

func (api *MetagameApi) Start() {
	openRouter := http.NewServeMux()
	openRouter.HandleFunc(HealthPath, api.Health)
	openRouter.HandleFunc(CreatePlayerPath, api.CreatePlayer)

	playerRouter := http.NewServeMux()
	playerRouter.HandleFunc(GetPlayerPath, api.GetPlayer)

	authenticationMiddleware := middlewares.NewAuthenticationMiddleware(api.securityUsecase)
	loggerMiddleware := middlewares.NewLoggerMiddleware()

	openRouter.Handle("/", authenticationMiddleware.CheckAuthentication(
		loggerMiddleware.SetupPlayerLog(
			playerRouter,
		)))

	err := http.ListenAndServe(":8080", openRouter)
	if err != nil {
		panic(err)
	}
}

func (api *MetagameApi) GetPlayerIdFromCtx(ctx context.Context) uuid.UUID {
	playerId := ctx.Value(constants.ContextKeyPlayerID).(uuid.UUID)
	return playerId
}
