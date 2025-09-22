package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const GetPlayerPath = "GET /v1/player"

func (handler MetagameApi) GetPlayer(w http.ResponseWriter, r *http.Request) {
	logger := handler.GetLoggerFromContext(r.Context())
	playerID := handler.GetPlayerIdFromCtx(r.Context())

	player, err := handler.playerUsecase.GetPlayer(playerID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	returnJson, err := json.Marshal(player)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(returnJson)
	if err != nil {
		logger.Error(fmt.Sprintf("error writing response: %v", err))
		panic(err)
	}
}
