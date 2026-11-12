package models

type InfusData struct {
	TPM        float32 `json:"tpm" bson:"tpm"`
	InfusType  string  `json:"infusType" bson:"infusType"`
	Percent    float32 `json:"percent" bson:"percent"`
	ServoAngle int32   `json:"servoAngle" bson:"servoAngle"`
	Volume     float32 `json:"volume" bson:"volume"`
}

type InfusPayload struct {
	Data      InfusData `json:"data" bson:"data"`
	DeviceID  string    `json:"deviceId" bson:"deviceId"`
	PatientID string    `json:"patientId" bson:"patientId"`
	RoomID    string    `json:"roomId" bson:"roomId"`

	Timestamp string `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}
