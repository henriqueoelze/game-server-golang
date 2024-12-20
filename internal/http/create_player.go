package http

import (
	"net/http"
)

const CreatePlayerPath = "POST /v1/player"

func (handler MetagameApi) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	newPlayer, err := handler.playerUsecase.CreatePlayer()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	playerId := newPlayer.PublicId
	encryptedPlayerId, err := handler.securityUsecase.Encrypt(playerId.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Player-Key", encryptedPlayerId)
	w.WriteHeader(http.StatusCreated)
}
