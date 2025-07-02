package game

import (
	"sync"
)

type GameHandler struct {
	mut   sync.RWMutex
	games []Game
}
