package http

import "net/http"

const HealthPath = "GET /v1/health"

func (handler MetagameApi) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
