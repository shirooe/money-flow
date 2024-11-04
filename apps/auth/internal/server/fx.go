package server

import (
	"github.com/shirooe/money-flow/apps/auth/internal/server/interceptor"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"server",
		fx.Provide(ProvideConfig, ProvideOption, ProvideServer),
		fx.Options(interceptor.Module()),
		fx.Invoke(InvokeServer),
	)
}
