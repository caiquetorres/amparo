package handlers

import "net/http"

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

// @Summary Ping
// @Description check connection
// @Produce  plain
// @Success 200
// @Router /api/ping [get]
func (h *PingHandler) HandlePingGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
