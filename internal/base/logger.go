package base

import (
	"context"
	"game-server-golang/internal/constants"
	"game-server-golang/internal/gateways"
)

type BaseLogger struct{}

func (base *BaseLogger) GetLogger(ctx context.Context) gateways.Logger {
	ctxLog := ctx.Value(constants.ContextKeyLogger).(gateways.Logger)
	return ctxLog
}
