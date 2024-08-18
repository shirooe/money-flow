package server

import (
	"go.uber.org/config"
	"go.uber.org/zap"
)

type Config struct {
	ServerPort int `yaml:"server_port"`
}

func NewConfig(provider *config.YAML, log *zap.Logger) Config {
	var config Config

	if err := provider.Get("auth").Populate(&config); err != nil {
		log.Sugar().Errorf("failed to get config: %v", err)
		return config
	}

	return config
}
