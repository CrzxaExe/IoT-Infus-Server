package models

// Model for Infusion Data
type InfusData struct {
	// Tetes Per Menit or Drop Per Minute(DPM)
	// How many drop in a minutes
	TPM float32 `json:"tpm" bson:"tpm"`

	// Infusion liquid type
	InfusType string `json:"infusType" bson:"infusType"`

	// How many liquid left of the infusion
	// Scale 0-100%
	Percent float32 `json:"percent" bson:"percent"`

	// The angle of the servo that clamp the infusion hose
	// It controll how fast the DPM
	// The more higher the value or more squeezing the hose, DPM will drop
	// Range: 0-180
	ServoAngle int32 `json:"servoAngle" bson:"servoAngle"`

	// Total weight of the infusion
	Volume float32 `json:"volume" bson:"volume"`
}

// Model for infusion payload
// From device into MQTT
type InfusPayload struct {
	// Infusion data
	Data InfusData `json:"data" bson:"data"`

	// Device Identifier
	// Device that send the data
	DeviceID string `json:"deviceId" bson:"deviceId"`

	// Patient Identifier
	// Patient the infusion set
	PatientID string `json:"patientId" bson:"patientId"`

	// Room Identifier
	RoomID string `json:"roomId" bson:"roomId"`

	// Time of the payload submited
	// Server will adding this automaticly
	// Format: YYYY-MM-DD SS:MS:NS GMT MST
	Timestamp string `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}
