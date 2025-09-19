package logger

import (
	"game-server-golang/internal/gateway"

	"go.uber.org/zap"
)

type ZapLogger struct {
	zap *zap.Logger
}

func NewZapLogger() gateway.Logger {
	logger, _ := zap.NewProduction()

	return &ZapLogger{
		zap: logger,
	}
}

func (l *ZapLogger) Info(msg string) {
	l.zap.Info(msg)
}

func (l *ZapLogger) WithField(key string, value interface{}) gateway.Logger {
	l.zap = l.zap.With(zap.Any(key, value))
	return l
}
