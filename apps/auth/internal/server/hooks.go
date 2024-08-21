package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func InvokeServer(lc fx.Lifecycle, server *grpc.Server, config Config, log *slog.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
			if err != nil {
				log.Error("failed to listen", slog.Any("error", err))
				return err
			}

			go func() {
				log.Info("Starting server on port", slog.Int("port", config.Port))
				if err := server.Serve(lis); err != nil {
					log.Error("failed to serve", slog.Any("error", err))
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.GracefulStop()
			return nil
		},
	})
}
