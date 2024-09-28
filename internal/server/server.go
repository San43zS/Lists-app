package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func New(port string, handler http.Handler) *Server {
	var httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
		// Size of the HTTP request header allowed.
		MaxHeaderBytes: 1 << 20, // 1mb
		// Timeouts when server cannot read/write data from client for 10s
		// Then connection will be closed
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return &Server{httpServer: httpServer}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

// ctx - is a contex of the running server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
