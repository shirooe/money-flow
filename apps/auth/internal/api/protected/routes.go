package protected

import (
	"context"

	"github.com/shirooe/money-flow/apps/auth/pkg/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *api) Logout(context.Context, *grpc.LogoutRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (a *api) AccessToken(context.Context, *grpc.AccessTokenRequest) (*grpc.AccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccessToken not implemented")
}
func (a *api) RefreshToken(context.Context, *grpc.RefreshTokenRequest) (*grpc.RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
