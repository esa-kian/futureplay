package api

import (
	"encoding/json"
	"futureplay/internal/model"
	"futureplay/internal/service"
	"net/http"
)

type Handler struct {
	matchmaker *service.Matchmaker
}

func NewHandler(matchmaker *service.Matchmaker) *Handler {
	return &Handler{matchmaker: matchmaker}
}

func (h *Handler) JoinMatchmaking(w http.ResponseWriter, r *http.Request) {
	var player model.Player
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.matchmaker.JoinMatchmaking(player)
	w.WriteHeader(http.StatusOK)
}
