package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Config struct {
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
	BasePath  string `json:"base_path_astrobin"`
}

var Data Config

func LoadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	fileContent, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	err = json.Unmarshal(fileContent, &Data)
	if err != nil {
		log.Fatalf("failed to unmarshal file content: %v", err)
	}

	if Data.ApiKey == "" {
		log.Fatal("missing API Key")
	}

	if Data.ApiSecret == "" {
		log.Fatal("missing API Secret")
	}
	if Data.BasePath == "" {
		log.Fatal("missing Base Path")
	}
}
