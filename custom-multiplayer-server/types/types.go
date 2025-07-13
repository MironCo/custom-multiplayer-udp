package types

import "encoding/json"

type WebsocketMessage struct {
	MessageType string          `json:"message_type"`
	MessageData json.RawMessage `json:"message_data"`
}
