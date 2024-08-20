package postgres

import (
	"context"

	"go.uber.org/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func Module(service string) fx.Option {
	return fx.Module(
		"database",
		fx.Provide(func(provider *config.YAML, logger *zap.Logger) Config {
			return NewConfig(provider, logger, service)
		}, NewClient),
		fx.Invoke(
			func(lc fx.Lifecycle, client *Client) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						return client.Ping(ctx)
					},
					OnStop: func(ctx context.Context) error {
						return client.Close()
					},
				})
			},
		),
	)
}
