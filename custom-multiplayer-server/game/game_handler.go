package game

import (
	"math/rand/v2"
	"net"
	"sync"
)

type GameHandler struct {
	mutex sync.RWMutex
	games map[string]*Game
}

func CreateNewGameHandler() *GameHandler {
	return &GameHandler{
		mutex: sync.RWMutex{},
		games: map[string]*Game{},
	}
}

func (h *GameHandler) CreateNewGameIntoHandler() {
	newGame := CreateNewGame()

	h.mutex.Lock()
	h.games[newGame.RoomID] = newGame
	h.mutex.Unlock()
}

func (h *GameHandler) AddPlayerToRandomGame(address *net.UDPAddr) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	if len(h.games) == 0 {
		// Create a new game if none exist
		newGame := CreateNewGame()
		h.games[newGame.RoomID] = newGame
		newGame.AddPlayer(address)
		return
	}

	randomKey := randomKeyFromMap(h.games)
	h.games[randomKey].AddPlayer(address)
}

// helper function
func randomKeyFromMap[K comparable, V any](m map[K]V) K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	randomKey := keys[rand.IntN(len(keys))]
	return randomKey
}
