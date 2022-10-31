package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	ServerAddr = os.Getenv("SERVER_ADDR")
)

func New(mux *http.ServeMux, ctx context.Context, l *log.Logger) *http.Server {

	l.Printf("Creating server at address: %s", ServerAddr)
	var s *http.Server

	s = &http.Server{
		Addr:         ServerAddr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      mux,
	}

	return s
}
