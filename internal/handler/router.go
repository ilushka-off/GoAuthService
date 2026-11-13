package handler

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
)

func NewRouter(healthHandler *HealthHandler, logger *slog.Logger) http.Handler {
	r := chi.NewRouter()

	r.Get("/health", healthHandler.Check)

	return r
}
