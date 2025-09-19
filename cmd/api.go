package cmd

import (
	"fmt"
	"game-server-golang/internal/config"
	"game-server-golang/internal/core"
	"game-server-golang/internal/gateway/sql_lite"
	"game-server-golang/internal/http"
	"game-server-golang/internal/usecase/player"
	"game-server-golang/internal/usecase/security"
)

func ExecuteApi(config *config.Config) error {
	// Initialize the base logger
	baseLogger := core.InitLogger(config.Logging)
	baseLogger.Info("initializing application")

	// Initialize database with config
	playerRepository, err := sql_lite.NewPlayerRepositoryImpl(config.Database.Name)
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	playerUsecase := player.NewPlayerUsecase(playerRepository)
	securityUsecase := security.NewSecurityUsecase()

	metagameApi := http.NewMetagameApi(
		securityUsecase,
		playerUsecase,
		baseLogger,
	)

	serverAddr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	baseLogger.Info(fmt.Sprintf("starting server %s", serverAddr))
	metagameApi.Start(serverAddr)

	return nil
}
