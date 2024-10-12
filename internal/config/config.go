package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	MQTTBroker  string `json:"mqtt_broker"`
	ClientID    string `json:"client_id"`
	Topic       string `json:"topic"`
	BatchSize   int    `json:"batch_size"`
	APIEndpoint string `json:"api_endpoint"`
}

func Load() (*Config, error) {
	// Load config from JSON file
	cfg, err := loadConfigFile()
	if err != nil {
		return nil, fmt.Errorf("error loading config file: %v", err)
	}

	// Override with environment variables if they exist
	if envValue := os.Getenv("MQTT_BROKER"); envValue != "" {
		cfg.MQTTBroker = envValue
	}
	if envValue := os.Getenv("CLIENT_ID"); envValue != "" {
		cfg.ClientID = envValue
	}
	if envValue := os.Getenv("TOPIC"); envValue != "" {
		cfg.Topic = envValue
	}
	if envValue := os.Getenv("BATCH_SIZE"); envValue != "" {
		batchSize, err := strconv.Atoi(envValue)
		if err != nil {
			return nil, fmt.Errorf("invalid BATCH_SIZE in environment: %v", err)
		}
		cfg.BatchSize = batchSize
	}
	if envValue := os.Getenv("API_ENDPOINT"); envValue != "" {
		cfg.APIEndpoint = envValue
	}

	return cfg, nil
}

func loadConfigFile() (*Config, error) {
	file, err := os.Open("configs/config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
