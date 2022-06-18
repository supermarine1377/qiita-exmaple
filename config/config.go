package config

import (
	"encoding/json"
	"os"
)

type ServerConfig struct {
	Port string `json:"port"`
}

func GetConfigFile() (*os.File, error) {
	return os.Open("config.json")
}
func GetConfig(file *os.File) (*ServerConfig, error) {
	var dec = json.NewDecoder(file)
	var conf ServerConfig
	if err := dec.Decode(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
