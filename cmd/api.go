package cmd

import (
	"fmt"
	"game-server-golang/internal/config"
	"game-server-golang/internal/gateway/sql_lite"
	"game-server-golang/internal/http"
	"game-server-golang/internal/usecase/player"
	"game-server-golang/internal/usecase/security"
)

func ExecuteApi(config *config.Config) error {
	// Initialize database with config
	playerRepository, err := sql_lite.NewPlayerRepositoryImpl(config.Database.Name)
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	playerUsecase := player.NewPlayerUsecase(playerRepository)
	securityUsecase := security.NewSecurityUsecase()

	serverAddr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	metagameApi := http.NewMetagameApi(securityUsecase, playerUsecase)
	metagameApi.Start(serverAddr)

	return nil
}
