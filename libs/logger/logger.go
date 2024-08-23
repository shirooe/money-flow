package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func New() *zerolog.Logger {
	if env := os.Getenv("APP_ENV"); env == "production" {
		logger := zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
			FormatMessage: func(i interface{}) string {
				return fmt.Sprintf("| %s |", i)
			},
			FormatCaller: func(i interface{}) string {
				return fmt.Sprintf("%v", i)
			}}).With().Timestamp().Caller().Logger().Level(zerolog.InfoLevel)

		return &logger
	}

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger().Level(zerolog.DebugLevel)
	return &logger
}
