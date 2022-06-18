package main

import (
	"qiita/config"
	"qiita/server"
)

type Server struct {
	Conf Config
}

type Config interface {
	GetPort() string
}

func main() {
	var file, err = config.GetConfigFile()
	if err != nil {
		panic(err)
	}
	conf, err := config.GetConfig(file)
	if err != nil {
		panic(err)
	}
	var server = server.NewServer(conf)
	server.Serve()
}
