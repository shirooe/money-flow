package server

import (
	"github.com/shirooe/gomf/libs/logger"
	"go.uber.org/config"
)

type Config struct {
	Port int `yaml:"port"`
}

func ProvideConfig(provider *config.YAML, log *logger.Logger) Config {
	var config Config

	if err := provider.Get("server").Populate(&config); err != nil {
		log.Errorf("failed to get config %v", err)
		return config
	}

	return config
}
