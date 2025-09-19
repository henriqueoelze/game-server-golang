package middleware

import (
	"context"
	"game-server-golang/internal/constant"
	"game-server-golang/internal/core"
	"game-server-golang/internal/gateway"
	"game-server-golang/internal/gateway/logger"
	"net/http"

	"github.com/google/uuid"
)

type LoggerMiddleware struct {
	core.BaseLogger
}

func NewLoggerMiddleware() *LoggerMiddleware {
	return &LoggerMiddleware{}
}

func (middleware LoggerMiddleware) SetupPlayerLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		playerId := r.Context().Value(constant.ContextKeyPlayerID).(uuid.UUID)

		playerContext, _ := initLogger(r.Context(), playerId)
		next.ServeHTTP(w, r.WithContext(playerContext))
	})
}

func initLogger(ctx context.Context, playerId uuid.UUID) (context.Context, gateway.Logger) {
	ctxLog := logger.NewZapLogger()
	ctxLog = ctxLog.WithField(constant.LoggerPlayerIdField, playerId)

	ctx = context.WithValue(ctx, constant.ContextKeyLogger, ctxLog)
	return ctx, ctxLog
}
