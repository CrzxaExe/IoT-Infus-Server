package models

// Models for commanding IoT.
// From Websocket into MQTT
type SocketCommand[T any] struct {
	// Type of commands
	// Value: set | reset
	Type string `json:"type" bson:"type"`

	// Device identifier
	// ID of the device want to give the command
	DeviceID string `json:"deviceId" bson:"deviceId"`

	// Time of the command submited
	// Server will adding this automaticly
	// Format: YYYY-MM-DD SS:MS:NS GMT MST
	Timestamp string `json:"timestamp,omitempty" bson:"timestamp"`

	// Device validation
	// Device has accept the command and the time it accepted
	IsSuccess        bool   `json:"isSuccess" bson:"isSuccess"`
	SuccessTimestamp string `json:"successTimestamp" bson:"successTimestamp"`

	// Data of the command
	// To controll the device
	Data T `json:"data" bson:"data"`
}

// Models used to wrapping data into Websocket clients
type WebsocketPayload[T any] struct {
	// Type of request
	Type string `json:"type" bson:"type"`

	// Message of request
	// (Optional)
	Message string `json:"message" bson:"message"`

	// The data
	Result T `json:"data" bson:"data"`
}
