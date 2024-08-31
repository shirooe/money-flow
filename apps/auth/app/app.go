package app

import (
	"context"

	"go.uber.org/fx"

	"github.com/shirooe/money-flow/apps/auth/internal/server"
	"github.com/shirooe/money-flow/libs/config"
	"github.com/shirooe/money-flow/libs/database/postgres"
	"github.com/shirooe/money-flow/libs/logger"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(func() context.Context { return context.Background() }),
		fx.Provide(config.New, logger.New),
		fx.Options(postgres.Module(), server.Module()),
	)
}
