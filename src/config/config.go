package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Logger  LoggerConfig
}

func LoadConfig(path string) Configuration {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)
	var config Configuration
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	return config
}