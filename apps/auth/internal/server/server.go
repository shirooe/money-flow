package server

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type Server struct {
	fx.Out

	*grpc.Server
}

func NewServer(opts []grpc.ServerOption) Server {
	server := grpc.NewServer(opts...)

	return Server{
		Server: server,
	}
}
