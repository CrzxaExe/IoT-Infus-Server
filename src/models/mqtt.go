package models

import "time"

type InfusData struct {
	tpm       float32
	infusType string
	volume    float32
}

type InfusPayload struct {
	data      InfusData
	deviceID  string
	patientID string
	roomID    string
	timestamp time.Time
}
