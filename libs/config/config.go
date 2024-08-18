package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"go.uber.org/config"
)

func New() *config.YAML {
	pwd, _ := os.Getwd()
	root := filepath.Join(pwd, "../..")

	if err := godotenv.Load(filepath.Join(root, ".env")); err != nil {
		log.Fatal(err)
	}

	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		log.Fatal("Config path is not set or dotenv file not found")
	}

	provider, err := config.NewYAML(config.File(filepath.Join(root, path)))
	if err != nil {
		log.Fatal(err)
	}

	return provider
}
