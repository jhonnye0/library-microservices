package api

import (
	"log"
	"net/http"
	"time"
)

type Handler struct {
	logger *log.Logger
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}

func (h *Handler) LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.logger.Println("Request received")
		defer h.logger.Printf("Request completed in %v", time.Since(start))

		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func NewHandler(l *log.Logger) *Handler {
	return &Handler{l}
}
