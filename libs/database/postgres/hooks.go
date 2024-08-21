package postgres

import (
	"context"

	"go.uber.org/fx"
)

func InvokeClient(lc fx.Lifecycle, client *Client) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return client.Ping(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})
}
