package http

import (
	"game-server-golang/internal/base"
	"game-server-golang/internal/http/middlewares"
	"game-server-golang/internal/usecases"
	"net/http"
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
