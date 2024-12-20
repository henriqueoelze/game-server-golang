package middlewares

import (
	"context"
	"game-server-golang/internal/base"
	"game-server-golang/internal/constants"
	"game-server-golang/internal/gateways"
	"game-server-golang/internal/gateways/logger"
	"net/http"

	"github.com/google/uuid"
)

type LoggerMiddleware struct {
	base.BaseLogger
}

func NewLoggerMiddleware() *LoggerMiddleware {
	return &LoggerMiddleware{}
}

func (middleware LoggerMiddleware) SetupPlayerLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		playerId := r.Context().Value(constants.ContextKeyPlayerID).(uuid.UUID)

		playerContext, _ := initLogger(r.Context(), playerId)
		next.ServeHTTP(w, r.WithContext(playerContext))
	})
}

func initLogger(ctx context.Context, playerId uuid.UUID) (context.Context, gateways.Logger) {
	ctxLog := logger.NewZapLogger()
	ctxLog = ctxLog.WithField(logger.PlayerIdField, playerId)

	ctx = context.WithValue(ctx, constants.ContextKeyLogger, ctxLog)
	return ctx, ctxLog
}
