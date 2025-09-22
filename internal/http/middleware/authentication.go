package middleware

import (
	"context"
	"fmt"
	"game-server-golang/internal/constant"
	"game-server-golang/internal/gateway"
	"game-server-golang/internal/usecase"
	"net/http"

	"github.com/google/uuid"
)

type AuthenticationMiddleware struct {
	securityUsecase usecase.SecurityUsecase

	baseLogger gateway.Logger
}

func NewAuthenticationMiddleware(
	securityUsecase usecase.SecurityUsecase,
	baseLogger gateway.Logger,
) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{
		securityUsecase: securityUsecase,
		baseLogger:      baseLogger,
	}
}

func (api AuthenticationMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		header, err := api.securityUsecase.Decrypt(authHeader)
		if err != nil {
			api.baseLogger.Error(fmt.Sprintf("error decrypting auth header: %v", err))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		playerID, err := uuid.Parse(header)
		if err != nil {
			api.baseLogger.Error(fmt.Sprintf("error decrypting auth header: %v", err))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		reqLogger := api.baseLogger.WithField(constant.LoggerPlayerIdField, playerID)
		ctxWithLog := context.WithValue(r.Context(), constant.ContextKeyLogger, reqLogger)

		ctxWithLogAndPlayerId := context.WithValue(ctxWithLog, constant.ContextKeyPlayerID, playerID)

		next.ServeHTTP(w, r.WithContext(ctxWithLogAndPlayerId))
	})
}
