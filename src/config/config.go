package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	BrokerUri     string
	BrokerPort    string
	ClientName    string
	MongoUri      string
	WebsocketPort string
}

func LoadConfig() Configuration {
	godotenv.Load()
	var uri string
	if uri = GetEnv("MONGO_URI", ""); uri == "" {
		panic("Mongo URI not found on .env")
	}

	return Configuration{
		BrokerUri:     GetEnv("BROKER_URI", "localhost"),
		BrokerPort:    GetEnv("BROKER_PORT", "1883"),
		ClientName:    GetEnv("CLIENT_NAME", "Infus Server"),
		MongoUri:      uri,
		WebsocketPort: GetEnv("WEBSOCKET_PORT", "8080"),
	}
}

func GetEnv(key string, defaultValue string) string {
	result := os.Getenv(key)
	if result == "" {
		return defaultValue
	}
	return result
}
