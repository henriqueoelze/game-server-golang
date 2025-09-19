package core

import (
	"game-server-golang/internal/config"
	"game-server-golang/internal/gateway"
	"game-server-golang/internal/gateway/logger"
	"sync"
)

var (
	baseLogger gateway.Logger
	once       sync.Once
)

// InitLogger initializes the base logger with configuration
// This should be called once during application startup
func InitLogger(config config.LoggingConfig) gateway.Logger {
	once.Do(func() {
		baseLogger = logger.NewZapLogger(config)
	})

	return baseLogger
}
