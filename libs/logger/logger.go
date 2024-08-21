package logger

import (
	"log/slog"
	"os"
	"path/filepath"
)

func New() *slog.Logger {
	var level slog.Level

	if env := os.Getenv("ENVIRONMENT"); env == "production" {
		level = slog.LevelInfo
	} else {
		level = slog.LevelDebug
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     level,
		AddSource: false,
	}))

	pwd, err := os.Getwd()
	if err != nil {
		return logger
	}

	logger.With(slog.String("service", filepath.Base(pwd)))
	slog.SetDefault(logger)

	return logger
}
