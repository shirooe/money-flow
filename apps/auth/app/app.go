package app

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/shirooe/gomf/apps/auth/internal/server"
	"github.com/shirooe/gomf/libs/config"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(config.New, zap.NewDevelopment),
		fx.Options(server.Module()),
	)
}
