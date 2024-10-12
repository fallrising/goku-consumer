package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	MQTTBroker  string
	ClientID    string
	Topic       string
	BatchSize   int
	APIEndpoint string
}

func Load() (*Config, error) {
	batchSize, err := strconv.Atoi(getEnv("BATCH_SIZE", "10"))
	if err != nil {
		return nil, fmt.Errorf("invalid BATCH_SIZE: %v", err)
	}

	return &Config{
		MQTTBroker:  getEnv("MQTT_BROKER", "tcp://localhost:1883"),
		ClientID:    getEnv("CLIENT_ID", "goku-consumer"),
		Topic:       getEnv("TOPIC", "urls"),
		BatchSize:   batchSize,
		APIEndpoint: getEnv("API_ENDPOINT", "http://api.example.com/upload"),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
