package http

import (
	"fmt"
	"net/http"
)

const HealthPath = "GET /v1/health"

func (handler MetagameApi) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte("OK"))
	if err != nil {
		logger := handler.GetLoggerFromContext(r.Context())
		logger.Error(fmt.Sprintf("error writing response: %v", err))
		panic(err)
	}
}
