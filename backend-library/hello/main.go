package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jhonnye0/backend-library/hello/api"
	"github.com/jhonnye0/backend-library/hello/server"
)

func main() {
	ctx := context.Background()
	logger := log.New(os.Stdout, "server: ", log.LstdFlags|log.Lshortfile)
	mux := http.NewServeMux()

	h := api.NewHandler(logger)

	mux.HandleFunc("/", h.LoggerMiddleware(h.Hello))

	s := server.New(mux, ctx, logger)

	logger.Println("Running server...")

	err := s.ListenAndServe()
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
