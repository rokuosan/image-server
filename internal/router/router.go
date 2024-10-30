package router

import (
	"net/http"

	"github.com/rokuosan/image-server/internal/handler"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	// Handlers
	healthHandler := new(handler.HealthHandler)

	// Routes
	mux.HandleFunc("/health", healthHandler.Get)

	return mux
}
