package postgres

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"go.uber.org/config"
)

type Config struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	SSLMode string `yaml:"sslmode"`
}

func ProvideConfig(provider *config.YAML, log *zerolog.Logger) Config {
	var config Config

	if err := provider.Get("database").Populate(&config); err != nil {
		log.Error().Msgf("failed to load config %v", err)
	}

	return config
}

func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s port=%d sslmode=%s dbname=%s user=%s password=%s",
		c.Host, c.Port, c.SSLMode,
		os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))
}
