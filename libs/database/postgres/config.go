package postgres

import (
	"fmt"
	"os"

	"go.uber.org/config"
	"go.uber.org/zap"
)

type Config struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	SSLMode string `yaml:"sslmode"`
}

func NewConfig(provider *config.YAML, log *zap.Logger) Config {
	var config Config

	if err := provider.Get("database").Populate(&config); err != nil {
		log.Error(err.Error())
	}

	return config
}

func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s port=%d sslmode=%s dbname=%s user=%s password=%s",
		c.Host, c.Port, c.SSLMode,
		os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))
}
