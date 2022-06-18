package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	conf config
}

type config interface {
	GetPort() string
}

func NewServer(conf config) *Server {
	return &Server{conf: conf}
}

func (s *Server) Serve() {
	http.HandleFunc("/qiita", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello, world!")
	})
	http.ListenAndServe(":"+s.conf.GetPort(), nil)
}
