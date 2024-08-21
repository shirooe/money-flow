package config

import (
	"fmt"
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

	provider, err := config.NewYAML(config.File(filepath.Join(root, fmt.Sprintf("config/%s.yml", filepath.Base(pwd)))))
	if err != nil {
		log.Fatal(err)
	}

	return provider
}
