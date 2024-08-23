package server

import (
	"github.com/rs/zerolog"
	"go.uber.org/config"
)

type Config struct {
	Port int `yaml:"port"`
}

func ProvideConfig(provider *config.YAML, log *zerolog.Logger) Config {
	var config Config

	if err := provider.Get("server").Populate(&config); err != nil {
		log.Error().Msgf("failed to get config %v", err)
		return config
	}

	return config
}
