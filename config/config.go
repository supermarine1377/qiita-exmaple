package config

import (
	"encoding/json"
	"io"
	"os"
)

type ServerConfig struct {
	Port string `json:"port"`
}

func (sc *ServerConfig) GetPort() string {
	return sc.Port
}

func GetConfigFile() (*os.File, error) {
	return os.Open("config.json")
}
func GetConfig(reader io.Reader) (*ServerConfig, error) {
	var dec = json.NewDecoder(reader)
	var conf ServerConfig
	if err := dec.Decode(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
