package unprotected

import (
	"context"

	"github.com/shirooe/money-flow/apps/auth/pkg/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *api) Login(context.Context, *grpc.LoginRequest) (*grpc.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (a *api) Register(context.Context, *grpc.RegisterRequest) (*grpc.RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
