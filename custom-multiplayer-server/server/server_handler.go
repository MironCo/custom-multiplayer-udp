package server

import (
	"fmt"
	"log"
	"net"
)

type ServerHandler struct {
	address          *net.UDPAddr
	connection       *net.UDPConn
	WebsocketHandler *WebsocketHandler
}

func CreateServerHandler(websocketHandler *WebsocketHandler) *ServerHandler {
	return &ServerHandler{
		WebsocketHandler: websocketHandler,
	}
}

func (h *ServerHandler) Start() {
	address, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		log.Fatal("Error resolving UDP Address")
	}
	h.address = address

	h.connection, err = net.ListenUDP("udp", h.address)
	if err != nil {
		log.Fatal("Error starting UDP server:", err)
	}
	defer h.Close()

	fmt.Println("UDP Server listing on Port http://127.0.0.1:8080")

	buffer := make([]byte, 1024)
	h.ReceiveMessages(buffer)
}

func (h *ServerHandler) ReceiveMessages(buffer []byte) {
	for {
		n, clientAddress, err := h.connection.ReadFromUDP(buffer)
		if err != nil {
			log.Println("Error reading UDP message: ", err)
			continue
		}

		message := string(buffer[:n])
		fmt.Println("Received from %s: %s\n", clientAddress, message)

		response := fmt.Sprintf("Echo: %s", message)
		_, err = h.connection.WriteToUDP([]byte(response), clientAddress)
		if err != nil {
			log.Println("Error sending response: ", err)
		}
	}
}

func (h *ServerHandler) Close() {
	h.connection.Close()
}
