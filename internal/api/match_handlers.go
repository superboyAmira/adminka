package api

import (
	"adminka/internal/model"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

var (
	matchStorage = struct {
		sync.RWMutex
		matches map[uuid.UUID]model.Match
	}{matches: make(map[uuid.UUID]model.Match)}
)

func CreateMatch(w http.ResponseWriter, r *http.Request) {
	var match model.Match
	if err := json.NewDecoder(r.Body).Decode(&match); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	match.ID = uuid.New()

	matchStorage.Lock()
	matchStorage.matches[match.ID] = match
	matchStorage.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(match)
}

func GetMatches(w http.ResponseWriter, r *http.Request) {
	matchStorage.RLock()
	defer matchStorage.RUnlock()

	matches := make([]model.Match, 0, len(matchStorage.matches))
	for _, m := range matchStorage.matches {
		matches = append(matches, m)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matches)
}
