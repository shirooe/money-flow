package postgres

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"database",
		fx.Provide(ProvideConfig, ProvideClient),
		fx.Invoke(InvokeClient),
	)
}
