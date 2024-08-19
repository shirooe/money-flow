package postgres

import (
	"context"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"database",
		fx.Provide(NewConfig, NewClient),
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
