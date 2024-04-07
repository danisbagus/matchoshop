package server

import (
	"github.com/danisbagus/matchoshop/cmd/api"
)

type Server struct {
	Http *httpServer
}

type httpServer struct {
}

func NewServer() *Server {
	return &Server{
		Http: &httpServer{},
	}
}

func (http *httpServer) Start() {
	api.Set()
}
