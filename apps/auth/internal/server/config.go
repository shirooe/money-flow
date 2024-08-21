package server

import (
	"log/slog"

	"go.uber.org/config"
)

type Config struct {
	Port int `yaml:"port"`
}

func ProvideConfig(provider *config.YAML, log *slog.Logger) Config {
	var config Config

	if err := provider.Get("server").Populate(&config); err != nil {
		log.Error("failed to load config:", slog.Any("error", err))
		return config
	}

	return config
}
