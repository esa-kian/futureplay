package service

import (
	"futureplay/internal/config"
	"futureplay/internal/model"
	"futureplay/internal/storage"
	"log"
	"strconv"
	"testing"
	"time"
)

// Helper function to initialize the matchmaker and storage
func setup() (*storage.InMemoryStore, *Matchmaker, *config.Config) {
	cfg, err := config.LoadConfig("../../config.yaml")
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	store := storage.NewInMemoryStore()
	matchmaker := NewMatchmaker(store, cfg.Matchmaking.CompetitionSize, cfg.Matchmaking.WaitTimeSeconds, cfg.Matchmaking.LevelRange)
	return store, matchmaker, cfg
}

func TestJoinMatchmaking(t *testing.T) {
	store, matchmaker, _ := setup()

	player := model.Player{ID: "1", Level: 5, Country: "US"}
	matchmaker.JoinMatchmaking(player)

	if len(store.GetPendingPlayers("US")) != 1 {
		t.Errorf("Expected 1 pending player for US, got %d", len(store.GetPendingPlayers("US")))
	}

	// Allow time for matchmaking to process
	time.Sleep(1 * time.Second)
}

func TestProcessMatchmaking_InsufficientPlayers(t *testing.T) {
	store, matchmaker, _ := setup()

	player := model.Player{ID: "1", Level: 5, Country: "US"}
	matchmaker.JoinMatchmaking(player)

	// Sleep for 31 seconds to ensure the function clears pending players
	time.Sleep(31 * time.Second)

	if len(store.GetPendingPlayers("US")) != 0 {
		t.Errorf("Expected 0 pending players for US after timeout, got %d", len(store.GetPendingPlayers("US")))
	}
}

func TestProcessMatchmaking_SufficientPlayers(t *testing.T) {
	store, matchmaker, _ := setup()

	for i := 1; i <= matchmaker.competitionSize; i++ {
		player := model.Player{ID: strconv.Itoa(i), Level: 5, Country: "US"}
		matchmaker.JoinMatchmaking(player)
	}

	// Allow time for matchmaking to process
	time.Sleep(1 * time.Second)

	if len(store.GetPendingPlayers("US")) != 0 {
		t.Errorf("Expected 0 pending players for US after forming a competition, got %d", len(store.GetPendingPlayers("US")))
	}
}
