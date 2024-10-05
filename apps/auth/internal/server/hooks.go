package server

import (
	"context"
	"fmt"
	"net"

	"github.com/rs/zerolog"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func InvokeServer(lc fx.Lifecycle, server *grpc.Server, config Config, log *zerolog.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
			if err != nil {
				log.Error().Msgf("failed to listen %v", err)
				return err
			}

			reflection.Register(server)

			go func() {
				log.Info().Msgf("Starting server on port %d", config.Port)
				if err := server.Serve(lis); err != nil {
					log.Error().Msgf("failed to serve %v", err)
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
