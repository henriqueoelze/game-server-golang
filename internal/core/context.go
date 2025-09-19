package core

import (
	"context"
	"game-server-golang/internal/constant"
	"game-server-golang/internal/gateway"
)

type ContextActions struct{}

func (c *ContextActions) AddLoggerToContext(ctx context.Context, logger gateway.Logger) context.Context {
	return context.WithValue(ctx, constant.ContextKeyLogger, logger)
}

func (c *ContextActions) GetLoggerFromContext(ctx context.Context) gateway.Logger {
	return ctx.Value(constant.ContextKeyLogger).(gateway.Logger)
}

func (c *ContextActions) AddPlayerIDToContext(ctx context.Context, playerID string) context.Context {
	return context.WithValue(ctx, constant.ContextKeyPlayerID, playerID)
}

func (c *ContextActions) GetPlayerIDFromContext(ctx context.Context) string {
	return ctx.Value(constant.ContextKeyPlayerID).(string)
}
