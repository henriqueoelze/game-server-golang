package http

import (
	"encoding/json"
	"net/http"
)

var GetPlayerPath = "GET /v1/player"

func (handler MetagameApi) GetPlayer(w http.ResponseWriter, r *http.Request) {
	logger := handler.GetLogger(r.Context())
	playerId := handler.GetPlayerIdFromCtx(r.Context())

	player, err := handler.playerUsecase.GetPlayer(playerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	returnJson, err := json.Marshal(player)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logger.Info("player retrieved")
	w.Write([]byte(returnJson))
}
