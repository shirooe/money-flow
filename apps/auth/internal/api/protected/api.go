package protected

import (
	"github.com/shirooe/money-flow/apps/auth/internal/service/auth"
	"github.com/shirooe/money-flow/apps/auth/internal/service/user"
	"github.com/shirooe/money-flow/apps/auth/pkg/grpc"
)

type api struct {
	grpc.UnimplementedAuthProtectedServer

	authService auth.Service
	userService user.Service
}

func New(authService auth.Service, userService user.Service) grpc.AuthProtectedServer {
	return &api{
		authService: authService,
		userService: userService,
	}
}
