package app

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/shirooe/gomf/apps/auth/internal/server"
	"github.com/shirooe/gomf/libs/config"
	"github.com/shirooe/gomf/libs/database/postgres"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(func() context.Context {
			return context.Background()
		}),
		fx.Provide(config.New, zap.NewDevelopment),
		fx.Options(postgres.Module("auth"), server.Module()),
	)
}
