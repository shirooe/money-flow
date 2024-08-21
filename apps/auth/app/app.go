package app

import (
	"context"

	"go.uber.org/fx"

	"github.com/shirooe/gomf/apps/auth/internal/server"
	"github.com/shirooe/gomf/libs/config"
	"github.com/shirooe/gomf/libs/database/postgres"
	"github.com/shirooe/gomf/libs/logger"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(func() context.Context { return context.Background() }),
		fx.Provide(config.New, logger.New),
		fx.Options(postgres.Module(), server.Module()),
	)
}
