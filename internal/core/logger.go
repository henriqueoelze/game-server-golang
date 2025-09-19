package core

import (
	"context"
	"game-server-golang/internal/constant"
	"game-server-golang/internal/gateway"
)

type BaseLogger struct{}

func (base *BaseLogger) GetLogger(ctx context.Context) gateway.Logger {
	ctxLog := ctx.Value(constant.ContextKeyLogger).(gateway.Logger)
	return ctxLog
}
