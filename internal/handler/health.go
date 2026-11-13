package handler

import (
	"log/slog"
	"net/http"
	"time"
)

const healthTimeout = 2 * time.Second

type HealthHandler struct {
	logger *slog.Logger
}

func NewHealthHandler(logger *slog.Logger) *HealthHandler {
	return &HealthHandler{logger: logger}
}

func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Health check", "path", r.URL.Path)
	w.WriteHeader(http.StatusOK)
}
