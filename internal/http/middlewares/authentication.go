package middlewares

import (
	"context"
	"game-server-golang/internal/constants"
	"game-server-golang/internal/usecases"
	"net/http"

	"github.com/google/uuid"
)

type AuthenticationMiddleware struct {
	securityUsecase usecases.SecurityUsecase
}

func NewAuthenticationMiddleware(
	securityUsecase usecases.SecurityUsecase,
) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{
		securityUsecase: securityUsecase,
	}
}

func (api AuthenticationMiddleware) CheckAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		header, err := api.securityUsecase.Decrypt(authHeader)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		playerId, err := uuid.Parse(header)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		playerContext := context.WithValue(r.Context(), constants.ContextKeyPlayerID, playerId)
		next.ServeHTTP(w, r.WithContext(playerContext))
	})
}
