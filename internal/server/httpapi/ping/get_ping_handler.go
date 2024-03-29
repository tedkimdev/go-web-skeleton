package ping

import (
	"net/http"

	"github.com/tedkimdev/go-web-skeleton/internal/server/httpapi"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) Ping(w http.ResponseWriter, r *http.Request) {
	httpapi.RespondWithJSON(w, http.StatusOK, "pong")
}

func GetPing(w http.ResponseWriter, r *http.Request) {
	httpapi.RespondWithJSON(w, http.StatusOK, "pong")
}
