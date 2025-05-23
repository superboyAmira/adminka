package main

import (
	"adminka/internal/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/players", api.GetPlayers)
	http.HandleFunc("/players/create", api.CreatePlayer)

	http.HandleFunc("/matches", api.GetMatches)
	http.HandleFunc("/matches/create", api.CreateMatch)

	http.HandleFunc("/health", api.HealthCheck)

	log.Println("Starting Chess Tournament API on :5000")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
