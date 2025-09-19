package http

import (
	"context"
	"game-server-golang/internal/constant"
	"game-server-golang/internal/core"
	"game-server-golang/internal/http/middleware"
	"game-server-golang/internal/usecase"
	"net/http"

	"github.com/google/uuid"
)

type MetagameApi struct {
	core.BaseLogger
	securityUsecase usecase.SecurityUsecase
	playerUsecase   usecase.PlayerUsecase
}

func NewMetagameApi(
	securityUsecase usecase.SecurityUsecase,
	playerUsecase usecase.PlayerUsecase,
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

	authenticationMiddleware := middleware.NewAuthenticationMiddleware(api.securityUsecase)
	loggerMiddleware := middleware.NewLoggerMiddleware()

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
	playerId := ctx.Value(constant.ContextKeyPlayerID).(uuid.UUID)
	return playerId
}
