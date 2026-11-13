package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/ilushka-off/GoAuthService/internal/handler"
)

func main() {
	if err := run(); err != nil {
		slog.Error("service stopped", "err", err)
		os.Exit(1)
	}
}

func run() error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	healthHandler := handler.NewHealthHandler(logger)

	router := handler.NewRouter(healthHandler, logger)

	logger.Info("service start", "port", ":8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		return fmt.Errorf("service start: %w", err)
	}

	return nil
}
