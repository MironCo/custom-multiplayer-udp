package types

import "encoding/json"

const (
	MESSAGE_TYPE_JOIN = "join_lobby"
)

type WebsocketMessage struct {
	MessageType string          `json:"message_type"`
	MessageData json.RawMessage `json:"message_data"`
}

type JoinLobbyMessage struct {
	UDPAddress string `json:"udp_address"`
}

type JoinLobbyResponse struct {
	UUID       string `json:"uuid"`
	UDPAddress string `json:"udp_address"`
}
