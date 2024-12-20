package cmd

import (
	"game-server-golang/internal/gateways/sql_lite"
	"game-server-golang/internal/http"
	"game-server-golang/internal/usecases/player"
	"game-server-golang/internal/usecases/security"
)

func ExecuteApi() error {
	playerRepository, err := sql_lite.NewPlayerRepositoryImpl()
	if err != nil {
		return err
	}

	playerUsecase := player.NewPlayerUsecase(playerRepository)
	securityUsecase := security.NewSecurityUsecase()
	metagameApi := http.NewMetagameApi(securityUsecase, playerUsecase)
	metagameApi.Start()

	return nil
}
