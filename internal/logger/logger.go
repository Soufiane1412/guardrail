package logger

import (
	"log/slog"
	"os"
)

// New creates an enterprise-grade JSON logger
func New() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo, // Set to debug in dev, Info in prod
	}
	// JSON format is mandatory for log aggregators like Datadog or ELK
	handler := slog.NewJSONHandler(os.Stdout, opts)
	return slog.New(handler)
}
