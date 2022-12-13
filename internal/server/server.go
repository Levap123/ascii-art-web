package server

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           "0.0.0.0:" + port,
		Handler:        handler,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	
	return s.httpServer.ListenAndServe()
}
