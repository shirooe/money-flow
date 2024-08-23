package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	log *zerolog.Logger
}

func New() *Logger {
	logger := zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
		NoColor:    false,
		FormatLevel: func(i interface{}) string {
			return fmt.Sprintf("[%s]", i)
		},
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf("| %s |", i)
		},
		FormatCaller: func(i interface{}) string {
			return filepath.Base(fmt.Sprintf("%s", i))
		},
	}).With().Timestamp().Caller().Logger()

	if env := os.Getenv("APP_ENV"); env == "production" {
		logger = logger.Level(zerolog.InfoLevel)
		return &Logger{log: &logger}
	}

	logger.Level(zerolog.DebugLevel)
	return &Logger{log: &logger}
}

func (l *Logger) Debug(args ...interface{}) {
	l.log.Debug().Msg(fmt.Sprint(args...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.log.Debug().Msgf(format, v...)
}

func (l *Logger) Debugw(msg string, KV map[string]interface{}) {
	l.log.Debug().Fields(KV).Msg(msg)
}

func (l *Logger) Info(args ...interface{}) {
	l.log.Info().Msg(fmt.Sprint(args...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.log.Info().Msgf(format, v...)
}

func (l *Logger) Infow(msg string, KV map[string]interface{}) {
	l.log.Info().Fields(KV).Msg(msg)
}

func (l *Logger) Warn(args ...interface{}) {
	l.log.Warn().Msg(fmt.Sprint(args...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.log.Warn().Msgf(format, v...)
}

func (l *Logger) Warnw(msg string, KV map[string]interface{}) {
	l.log.Warn().Fields(KV).Msg(msg)
}

func (l *Logger) Error(args ...interface{}) {
	l.log.Error().Msg(fmt.Sprint(args...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.log.Error().Msgf(format, v...)
}

func (l *Logger) Errorw(msg string, KV map[string]interface{}) {
	l.log.Error().Fields(KV).Msg(msg)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.log.Fatal().Msg(fmt.Sprint(args...))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.log.Fatal().Msgf(format, v...)
}

func (l *Logger) Fatalw(msg string, KV map[string]interface{}) {
	l.log.Fatal().Fields(KV).Msg(msg)
}
