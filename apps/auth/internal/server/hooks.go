package server

import (
	"context"
	"fmt"
	"net"

	"github.com/shirooe/gomf/libs/logger"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func InvokeServer(lc fx.Lifecycle, server *grpc.Server, config Config, log *logger.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
			if err != nil {
				log.Errorf("failed to listen %v", err)
				return err
			}

			go func() {
				log.Infof("Starting server on port %d", config.Port)
				if err := server.Serve(lis); err != nil {
					log.Errorf("failed to serve %v", err)
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
