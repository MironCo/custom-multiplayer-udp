package game

import (
	"net"
	"sync"

	"github.com/google/uuid"
)

type Game struct {
	mutex   sync.RWMutex
	RoomID  string
	Players map[string]*Player
}

func CreateNewGame() *Game {
	return &Game{
		RoomID:  uuid.NewString(),
		Players: map[string]*Player{},
	}
}

func (h *Game) AddPlayer(address *net.UDPAddr) *Player {
	h.mutex.Lock()
	newPlayer := CreateNewPlayer(address)
	h.Players[newPlayer.PlayerUUID] = newPlayer
	h.mutex.Unlock()
	return newPlayer
}
