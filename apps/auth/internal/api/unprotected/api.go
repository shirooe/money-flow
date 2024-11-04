package unprotected

import (
	"github.com/shirooe/money-flow/apps/auth/internal/service/auth"
	"github.com/shirooe/money-flow/apps/auth/internal/service/user"
	"github.com/shirooe/money-flow/apps/auth/pkg/grpc"
)

type api struct {
	grpc.UnimplementedAuthUnprotectedServer

	authService auth.Service
	userService user.Service
}

func New(authService auth.Service, userService user.Service) grpc.AuthUnprotectedServer {
	return &api{
		authService: authService,
		userService: userService,
	}
}
