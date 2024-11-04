package server

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Option struct {
	fx.Out

	Options []grpc.ServerOption
}

func ProvideOption(interceptors []grpc.UnaryServerInterceptor) Option {
	options := []grpc.ServerOption{
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(interceptors...),
	}

	return Option{
		Options: options,
	}
}
