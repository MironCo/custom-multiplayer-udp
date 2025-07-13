package game

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type NetworkClient struct {
	conn      *websocket.Conn
	connected bool
}

func NewNetworkClient() *NetworkClient {
	return &NetworkClient{
		connected: false,
	}
}

func (nc *NetworkClient) ConnectToServer(serverURL string) error {
	u, err := url.Parse(serverURL)
	if err != nil {
		return fmt.Errorf("invalid URL: %v", err)
	}

	log.Printf("Connecting to server at %s", u.String())

	// TODO: Implement actual websocket connection
	// conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	// if err != nil {
	//     return fmt.Errorf("failed to connect: %v", err)
	// }

	// nc.conn = conn
	nc.connected = true
	log.Println("Successfully connected to server (simulated)")

	return nil
}

func (nc *NetworkClient) Disconnect() {
	if nc.conn != nil {
		nc.conn.Close()
		nc.conn = nil
	}
	nc.connected = false
}

func (nc *NetworkClient) IsConnected() bool {
	return nc.connected
}

func (nc *NetworkClient) SendMessage(message []byte) error {
	if !nc.connected || nc.conn == nil {
		return fmt.Errorf("not connected to server")
	}

	// TODO: Implement actual message sending
	// return nc.conn.WriteMessage(websocket.TextMessage, message)
	return nil
}
