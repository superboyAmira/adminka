package api

import (
	"adminka/internal/model"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

var (
	playerStorage = struct {
		sync.RWMutex
		players map[uuid.UUID]model.Player
	}{players: make(map[uuid.UUID]model.Player)}
)

// Создание нового игрока
func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var player model.Player
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	player.ID = uuid.New()

	playerStorage.Lock()
	playerStorage.players[player.ID] = player
	playerStorage.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(player)
}

// Получение всех игроков
func GetPlayers(w http.ResponseWriter, r *http.Request) {
	playerStorage.RLock()
	defer playerStorage.RUnlock()

	players := make([]model.Player, 0, len(playerStorage.players))
	for _, p := range playerStorage.players {
		players = append(players, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}
