package main

import (
	"futureplay/internal/api"
	"futureplay/internal/config"
	"futureplay/internal/service"
	"futureplay/internal/storage"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	store := storage.NewInMemoryStore()
	matchmaker := service.NewMatchmaker(store, cfg.Matchmaking.CompetitionSize, cfg.Matchmaking.WaitTimeSeconds, cfg.Matchmaking.LevelRange)

	handler := api.NewHandler(matchmaker)

	r := mux.NewRouter()
	r.HandleFunc("/matchmaking/join", handler.JoinMatchmaking).Methods("POST")

	port := ":" + strconv.Itoa(cfg.Server.Port)
	log.Printf("Starting server on http://localhost%s...", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Server failed: %s\n", err)
	}
}
