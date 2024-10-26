package storage

import (
	"futureplay/internal/model"
	"sync"
)

type InMemoryStore struct {
	players        map[string]model.Player
	pendingPlayers map[string][]model.Player
	competitions   map[string]model.Competition
	mu             sync.Mutex
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		players:        make(map[string]model.Player),
		pendingPlayers: make(map[string][]model.Player),
		competitions:   make(map[string]model.Competition),
	}
}

func (s *InMemoryStore) AddPlayer(player model.Player) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.players[player.ID] = player
	s.pendingPlayers[player.Country] = append(s.pendingPlayers[player.Country], player)
}

func (s *InMemoryStore) GetPendingPlayers(country string) []model.Player {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.pendingPlayers[country]
}

func (s *InMemoryStore) GetCompetitions() map[string]model.Competition {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.competitions
}

func (s *InMemoryStore) RemovePendingPlayers(country string, count int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.pendingPlayers[country]) >= count {
		s.pendingPlayers[country] = s.pendingPlayers[country][count:]
	}
}

func (s *InMemoryStore) CreateCompetition(id string, players []model.Player) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.competitions[id] = model.Competition{ID: id, Players: players, IsActive: true}
}
