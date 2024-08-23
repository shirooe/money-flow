package postgres

import (
	"context"

	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

func InvokeClient(lc fx.Lifecycle, client *Client, log *zerolog.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := client.Ping(ctx); err != nil {
				log.Error().Msgf("failed to ping %v", err)
				return err
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})
}
