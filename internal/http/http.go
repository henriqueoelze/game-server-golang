package http

import (
	"context"
	"fmt"
	"game-server-golang/internal/config"
	"game-server-golang/internal/constant"
	"game-server-golang/internal/core"
	"game-server-golang/internal/gateway"
	"game-server-golang/internal/http/middleware"
	"game-server-golang/internal/usecase"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type MetagameApi struct {
	core.ContextActions

	securityUsecase usecase.SecurityUsecase
	playerUsecase   usecase.PlayerUsecase

	baseLogger gateway.Logger
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

func (api MetagameApi) Start(serverConfig config.ServerConfig) {
	serverAddress := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)
	api.baseLogger.Info(fmt.Sprintf("starting server %s", serverAddress))

	openRouter := http.NewServeMux()
	openRouter.HandleFunc(HealthPath, api.Health)
	openRouter.HandleFunc(CreatePlayerPath, api.CreatePlayer)

	playerRouter := http.NewServeMux()
	playerRouter.HandleFunc(GetPlayerPath, api.GetPlayer)

	authenticationMiddleware := middleware.NewAuthenticationMiddleware(api.securityUsecase, api.baseLogger)

	openRouter.Handle("/", authenticationMiddleware.Authenticate(playerRouter))

	timeoutDuration := time.Duration(serverConfig.TimeOutInSeconds) * time.Second
	server := http.Server{
		Addr:              serverAddress,
		ReadHeaderTimeout: timeoutDuration,
		WriteTimeout:      timeoutDuration,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func (api MetagameApi) GetPlayerIdFromCtx(ctx context.Context) uuid.UUID {
	playerID := ctx.Value(constant.ContextKeyPlayerID).(uuid.UUID)
	return playerID
}
