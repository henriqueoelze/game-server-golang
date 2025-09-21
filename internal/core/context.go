package core

import (
	"context"
	"game-server-golang/internal/constant"
	"game-server-golang/internal/gateway"
)

type ContextActions struct{}

func (c *ContextActions) GetLoggerFromContext(ctx context.Context) gateway.Logger {
	return ctx.Value(constant.ContextKeyLogger).(gateway.Logger)
}
