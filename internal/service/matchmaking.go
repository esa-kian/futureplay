package service

import (
	"fmt"
	"futureplay/internal/model"
	"futureplay/internal/storage"
	"log"
	"math"
	"time"
)

// Matchmaker handles the matchmaking process.
type Matchmaker struct {
	store           *storage.InMemoryStore
	competitionSize int
	waitTimeSeconds int
	levelRange      int // Maximum allowed difference in player levels
}

// NewMatchmaker creates a new Matchmaker.
func NewMatchmaker(store *storage.InMemoryStore, competitionSize, waitTimeSeconds, levelRange int) *Matchmaker {
	return &Matchmaker{
		store:           store,
		competitionSize: competitionSize,
		waitTimeSeconds: waitTimeSeconds,
		levelRange:      levelRange,
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

	// Filter players by level
	matchedPlayers := m.filterPlayersByLevel(pendingPlayers)

	if len(matchedPlayers) < m.competitionSize {
		time.Sleep(time.Duration(m.waitTimeSeconds) * time.Second)
		pendingPlayers = m.store.GetPendingPlayers(country)
		matchedPlayers = m.filterPlayersByLevel(pendingPlayers)

		if len(matchedPlayers) < m.competitionSize {
			m.clearPendingPlayers(country)
			log.Printf("Not enough players to form a competition in country %s. Clearing pending players.\n", country)
			return
		}
	}

	m.createCompetition(country, matchedPlayers[:m.competitionSize])
}

// filterPlayersByLevel filters players within the allowed level range.
func (m *Matchmaker) filterPlayersByLevel(players []model.Player) []model.Player {
	if len(players) == 0 {
		return players
	}

	var filtered []model.Player
	// Calculate the average level of players to find a reference point
	var totalLevel int
	for _, player := range players {
		totalLevel += player.Level
	}
	averageLevel := totalLevel / len(players)

	// Filter players within the specified level range
	for _, player := range players {
		if math.Abs(float64(player.Level-averageLevel)) <= float64(m.levelRange) {
			filtered = append(filtered, player)
		}
	}
	log.Printf("ready players: %+v\n", filtered)

	return filtered
}

// createCompetition creates a competition with the selected players.
func (m *Matchmaker) createCompetition(country string, players []model.Player) {
	competitionID := generateCompetitionID() // UUID or similar
	m.store.CreateCompetition(competitionID, players)
	m.store.RemovePendingPlayers(country, len(players))
	log.Printf("Competition %s created with players: %+v\n", competitionID, players)
}

// clearPendingPlayers clears the pending players for the specified country.
func (m *Matchmaker) clearPendingPlayers(country string) {
	m.store.RemovePendingPlayers(country, len(m.store.GetPendingPlayers(country)))
}

func generateCompetitionID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
