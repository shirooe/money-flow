package server

import (
	"go.uber.org/config"
	"go.uber.org/zap"
)

type Config struct {
	Port int `yaml:"port"`
}

func ProvideConfig(provider *config.YAML, log *zap.Logger) Config {
	var config Config

	if err := provider.Get("server").Populate(&config); err != nil {
		log.Sugar().Errorf("failed to get config: %v", err)
		return config
	}

	return config
}
