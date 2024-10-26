package storage

import (
	"futureplay/internal/model"
	"testing"
)

func TestAddPlayer(t *testing.T) {
	store := NewInMemoryStore()
	player := model.Player{ID: "1", Level: 5, Country: "US"}

	store.AddPlayer(player)

	if len(store.players) != 1 {
		t.Errorf("Expected 1 player, got %d", len(store.players))
	}

	if len(store.pendingPlayers["US"]) != 1 {
		t.Errorf("Expected 1 pending player for US, got %d", len(store.pendingPlayers["US"]))
	}
}

func TestGetPendingPlayers(t *testing.T) {
	store := NewInMemoryStore()
	player1 := model.Player{ID: "1", Level: 5, Country: "US"}
	player2 := model.Player{ID: "2", Level: 6, Country: "US"}
	store.AddPlayer(player1)
	store.AddPlayer(player2)

	pendingPlayers := store.GetPendingPlayers("US")

	if len(pendingPlayers) != 2 {
		t.Errorf("Expected 2 pending players for US, got %d", len(pendingPlayers))
	}
}

func TestRemovePendingPlayers(t *testing.T) {
	store := NewInMemoryStore()
	player := model.Player{ID: "1", Level: 5, Country: "US"}
	store.AddPlayer(player)

	store.RemovePendingPlayers("US", 1)

	if len(store.GetPendingPlayers("US")) != 0 {
		t.Errorf("Expected 0 pending players for US, got %d", len(store.GetPendingPlayers("US")))
	}
}

func TestCreateCompetition(t *testing.T) {
	store := NewInMemoryStore()
	player1 := model.Player{ID: "1", Level: 5, Country: "US"}
	player2 := model.Player{ID: "2", Level: 5, Country: "US"}
	player3 := model.Player{ID: "3", Level: 5, Country: "US"}
	player4 := model.Player{ID: "4", Level: 5, Country: "US"}
	player5 := model.Player{ID: "5", Level: 5, Country: "US"}
	player6 := model.Player{ID: "6", Level: 5, Country: "US"}
	player7 := model.Player{ID: "7", Level: 5, Country: "US"}
	player8 := model.Player{ID: "8", Level: 5, Country: "US"}
	player9 := model.Player{ID: "9", Level: 5, Country: "US"}
	player10 := model.Player{ID: "10", Level: 5, Country: "US"}

	store.AddPlayer(player1)
	store.AddPlayer(player2)
	store.AddPlayer(player3)
	store.AddPlayer(player4)
	store.AddPlayer(player5)
	store.AddPlayer(player6)
	store.AddPlayer(player7)
	store.AddPlayer(player8)
	store.AddPlayer(player9)
	store.AddPlayer(player10)

	store.CreateCompetition("competition1", []model.Player{player1, player2, player3, player4, player5, player6, player7, player8, player9, player10})

	if len(store.competitions) != 1 {
		t.Errorf("Expected 1 competition, got %d", len(store.competitions))
	}
}
