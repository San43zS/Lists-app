package http

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func New(handler http.Handler) *Server {
	var httpServer = &http.Server{
		Addr:    "localhost" + ":" + viper.GetString("http_port"),
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

func (s Server) Run() error {
	fmt.Printf("Сервер запущен на адресе: http://%s\n", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

// ctx - is a contex of the running server
func (s Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
