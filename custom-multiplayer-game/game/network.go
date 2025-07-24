package game

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/url"

	"custom-multiplayer-game/types"

	"github.com/gorilla/websocket"
)

const (
	SERVER_BASE_ADDRESS   = "127.0.0.1"
	SERVER_WEBSOCKET_PORT = "8081"
	SERVER_UDP_PORT       = "8000"
)

type NetworkClient struct {
	conn          *websocket.Conn
	udpConn       *net.UDPConn
	udpConnString string
	connected     bool
}

func NewNetworkClient() *NetworkClient {
	return &NetworkClient{
		connected: false,
	}
}

func (nc *NetworkClient) ConnectToUDPServer() error {
	if nc.udpConn != nil {
		return nil
	}

	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 0})
	if err != nil {
		return fmt.Errorf("failed to bind UDP socket: %v", err)
	}

	nc.udpConn = conn
	nc.udpConnString = conn.LocalAddr().String()
	log.Println("Successfully connected to UDP server")

	return nil
}

func (nc *NetworkClient) ConnectToWebsocketServer() error {
	u, err := url.Parse("ws://" + SERVER_BASE_ADDRESS + ":" + SERVER_WEBSOCKET_PORT + "/ws")
	if err != nil {
		return fmt.Errorf("invalid URL: %v", err)
	}

	log.Printf("Connecting to server at %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to connect: %v", err)
	}

	nc.conn = conn
	nc.connected = true
	log.Println("Successfully connected to websocket server")
	return nil
}

func (nc *NetworkClient) ConnectToServer() error {
	nc.ConnectToWebsocketServer()
	nc.ConnectToUDPServer()

	return nil
}

func (nc *NetworkClient) JoinGame() error {
	joinGameMessage := &types.JoinLobbyMessage{
		UDPAddress: nc.udpConnString,
	}

	messageData, err := json.Marshal(joinGameMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal join message: %v", err)
	}

	websocketMessage := &types.WebsocketMessage{
		MessageType: types.MESSAGE_TYPE_JOIN,
		MessageData: messageData,
	}

	return nc.SendWebsocketMessage(websocketMessage)
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

func (nc *NetworkClient) SendWebsocketMessage(msg *types.WebsocketMessage) error {
	if !nc.connected || nc.conn == nil {
		return fmt.Errorf("not connected to server")
	}

	return nc.conn.WriteJSON(msg)
}
