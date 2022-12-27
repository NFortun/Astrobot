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
	User      string `json:"user"`
	Passwd    string `json:"password"`
	Port      int    `json:"port"`
	Addr      string `json:"addr"`
	DBName    string `json:"dbname"`
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

	if Data.Addr == "" {
		log.Fatal("missing database address")
	}

	if Data.DBName == "" {
		log.Fatal("missing database name")
	}

	if Data.Port <= 0 {
		log.Fatal("missing port")
	}

	if Data.User == "" {
		log.Fatal("missing database user")
	}

	if Data.Passwd == "" {
		log.Fatal("missing database password")
	}
}
