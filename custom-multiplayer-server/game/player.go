package game

import "net"

type Player struct {
	address net.UDPAddr
	X       float32
	Y       float32
}
