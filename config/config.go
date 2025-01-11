package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	ServerPort string `json:"SERVER_PORT"`
	DBDriver   string `json:"DB_DRIVER"`
	DBSource   string `json:"DB_SOURCE"`
}

var AppConfig Config

func LoadConfig() {

	file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}
}
