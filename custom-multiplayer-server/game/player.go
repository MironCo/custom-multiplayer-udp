package game

import (
	"net"

	"github.com/google/uuid"
)

type Player struct {
	PlayerUUID string      `json:"player_uuid"`
	Address    net.UDPAddr `json:"udp_address"`
	X          float32     `json:"x_position"`
	Y          float32     `json:"y_position"`
}

func CreateNewPlayer(address *net.UDPAddr) *Player {
	player := &Player{
		PlayerUUID: uuid.NewString(),
		Address:    *address,
		X:          0,
		Y:          0,
	}
	return player
}
