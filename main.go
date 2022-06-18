package main

import (
	"fmt"
	"net/http"
	"qiita/config"
)

type Server struct {
	conf Config
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
	var server = newServer(conf)
	server.serve()
}

func newServer(conf Config) *Server {
	return &Server{conf: conf}
}

func (s *Server) serve() {
	http.HandleFunc("/qiita", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello, world!")
	})
	http.ListenAndServe(":"+s.conf.GetPort(), nil)
}
