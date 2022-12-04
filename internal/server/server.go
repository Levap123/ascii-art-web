package server

import (
	"net/http"
	"time"
)

type Server struct {
	HttpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.HttpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	
	return s.HttpServer.ListenAndServe()
}
