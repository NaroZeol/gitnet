package config

import (
	"encoding/json"
	"log"
	"os"

	"gitnet/internal/pkg/database"
)

type Config struct {
	Database database.Database `json:"database"`
}

func LoadConfig(path string) Config {
	config := Config{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening config file:", err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Error decoding config file:", err)
	}
	file.Close()

	return config
}
