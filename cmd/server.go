package main

import (
	"futureplay/internal/api"
	"futureplay/internal/config"
	"futureplay/internal/service"
	"futureplay/internal/storage"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.LoadConfig("./config.yaml")
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	store := storage.NewInMemoryStore()
	matchmaker := service.NewMatchmaker(store, cfg.Matchmaking.CompetitionSize, cfg.Matchmaking.WaitTimeSeconds)

	handler := api.NewHandler(matchmaker)

	r := mux.NewRouter()
	r.HandleFunc("/matchmaking/join", handler.JoinMatchmaking).Methods("POST")

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed: %s\n", err)
	}
}
