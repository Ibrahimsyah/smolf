package http

import (
	"fmt"
	"net/http"
	"smolf-auth/internal/handler/http/health"
	"time"
)

type Handler struct {
	Health *health.Handler
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRouters() {
	h.handleFunc("/health-check", h.Health.HealthCheck)
}

func (h *Handler) handleFunc(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		handler(w, r)
		elapsedTime := time.Since(startTime)
		fmt.Printf("[%s] %d ms\n", path, elapsedTime.Milliseconds())
	})
}
