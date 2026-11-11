package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	BrokerUri  string
	ClientName string
	Port       string
}

func LoadConfig() Configuration {
	godotenv.Load()
	return Configuration{
		BrokerUri:  GetEnv("BROKER_URI", "localhost"),
		ClientName: GetEnv("CLIENT_NAME", "Infus Server"),
		Port:       GetEnv("BROKER_PORT", "1883"),
	}
}

func GetEnv(key string, defaultValue string) string {
	result := os.Getenv(key)
	if result == "" {
		return defaultValue
	}
	return result
}
