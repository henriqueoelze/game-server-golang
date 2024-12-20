package base

import (
	"context"
	"game-server-golang/internal/constants"

	"github.com/google/uuid"
)

type BaseSession struct{}

func (base *BaseSession) GetPlayerIdFromCtx(ctx context.Context) uuid.UUID {
	playerId := ctx.Value(constants.ContextKeyPlayerID).(uuid.UUID)
	return playerId
}
