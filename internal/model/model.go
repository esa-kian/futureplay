package model

type Player struct {
	ID      string `json:"id"`
	Level   int    `json:"level"`
	Country string `json:"country"`
}

type Competition struct {
	ID       string   `json:"id"`
	Players  []Player `json:"players"`
	IsActive bool     `json:"isActive"`
}
