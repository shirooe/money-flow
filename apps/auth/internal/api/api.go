package api

import (
	"github.com/shirooe/money-flow/apps/auth/pkg/grpc"
)

type api struct {
	grpc.UnimplementedAuthUnprotectedServer
}

func New() grpc.AuthUnprotectedServer {
	return &api{}
}
