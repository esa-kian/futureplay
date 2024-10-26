package api

import (
	"bytes"
	"encoding/json"
	"futureplay/internal/config"
	"futureplay/internal/model"
	"futureplay/internal/service"
	"futureplay/internal/storage"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setup() (*storage.InMemoryStore, *service.Matchmaker, *Handler) {
	cfg, err := config.LoadConfig("../../config.yaml")
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	store := storage.NewInMemoryStore()
	matchmaker := service.NewMatchmaker(store, cfg.Matchmaking.CompetitionSize, cfg.Matchmaking.WaitTimeSeconds)
	handler := NewHandler(matchmaker)

	return store, matchmaker, handler
}

func TestJoinMatchmaking(t *testing.T) {
	store, _, handler := setup()

	player := model.Player{ID: "1", Level: 5, Country: "US"}
	body, _ := json.Marshal(player)

	req, err := http.NewRequest("POST", "/matchmaking/join", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.JoinMatchmaking(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", status)
	}

	if len(store.GetPendingPlayers("US")) != 1 {
		t.Errorf("Expected 1 pending player for US, got %d", len(store.GetPendingPlayers("US")))
	}
}

func TestJoinMatchmaking_InvalidJSON(t *testing.T) {
	_, _, handler := setup()

	req, err := http.NewRequest("POST", "/matchmaking/join", bytes.NewBuffer([]byte(`invalid json`)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.JoinMatchmaking(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", status)
	}
}
