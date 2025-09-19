package constant

type contextKey string

const (
	// ContextKeyPlayerID is the key used to store the player ID in the context
	ContextKeyPlayerID contextKey = "context.playerId"
	ContextKeyLogger   contextKey = "context.logger"
)
