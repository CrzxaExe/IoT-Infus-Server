package models

type SocketCommand struct {
	Type     string `json:"type"`
	DeviceID string `json:"deviceId"`
	Data     any    `json:"data"`
}

type WebsocketPayload[T any] struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}
