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

	log.Println("Starting Chess Tournament API on :8085")
	if err := http.ListenAndServe(":8085", nil); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
