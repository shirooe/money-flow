package interceptor

import (
	"context"
	"strings"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type authInterceptor struct {
	logger *zerolog.Logger
}

func NewAuth(logger *zerolog.Logger) authInterceptor {
	return authInterceptor{
		logger: logger,
	}
}

func (a *authInterceptor) UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if strings.Contains(info.FullMethod, "/auth.AuthUnprotected") {
			a.logger.Info().Msg("auth unprotected method")
			return handler(ctx, req)
		}

		return handler(ctx, req)
	}
}
