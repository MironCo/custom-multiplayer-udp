package server

import "net"

type ServerHandler struct {
	address    *net.UDPAddr
	connection *net.UDPConn
}

func (h *ServerHandler) Start() {

}

func (h *ServerHandler) Close() {

}
