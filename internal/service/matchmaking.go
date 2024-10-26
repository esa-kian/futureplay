package service

import (
	"fmt"
	"futureplay/internal/model"
	"futureplay/internal/storage"
	"log"
	"time"
)

// Matchmaker handles the matchmaking process.
type Matchmaker struct {
	store           *storage.InMemoryStore
	competitionSize int
	waitTimeSeconds int
}

// NewMatchmaker creates a new Matchmaker.
func NewMatchmaker(store *storage.InMemoryStore, competitionSize, waitTimeSeconds int) *Matchmaker {
	return &Matchmaker{
		store:           store,
		competitionSize: competitionSize,
		waitTimeSeconds: waitTimeSeconds,
	}
}

// JoinMatchmaking adds a player to the matchmaking pool.
func (m *Matchmaker) JoinMatchmaking(player model.Player) {
	m.store.AddPlayer(player)
	go m.processMatchmaking(player.Country)
}

// processMatchmaking checks if enough players are available for a competition.
func (m *Matchmaker) processMatchmaking(country string) {
	pendingPlayers := m.store.GetPendingPlayers(country)

	if len(pendingPlayers) < m.competitionSize {
		time.Sleep(time.Duration(m.waitTimeSeconds) * time.Second)
		pendingPlayers = m.store.GetPendingPlayers(country)

		if len(pendingPlayers) < m.competitionSize {
			m.clearPendingPlayers(country)
			log.Printf("Not enough players to form a competition in country %s. Clearing pending players.\n", country)
			return
		}
	}

	m.createCompetition(country, pendingPlayers[:m.competitionSize])

}

// createCompetition creates a competition with the selected players.
func (m *Matchmaker) createCompetition(country string, players []model.Player) {
	competitionID := generateCompetitionID() // UUID or similar
	m.store.CreateCompetition(competitionID, players)
	m.store.RemovePendingPlayers(country, 10)
	log.Printf("Competition %s created with players: %+v\n", competitionID, players)
}

// clearPendingPlayers clears the pending players for the specified country.
func (m *Matchmaker) clearPendingPlayers(country string) {
	m.store.RemovePendingPlayers(country, len(m.store.GetPendingPlayers(country)))
}

func generateCompetitionID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
