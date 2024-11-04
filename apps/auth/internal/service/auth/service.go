package auth

import (
	"context"

	"github.com/shirooe/money-flow/apps/auth/internal/dto"
)

type Service interface {
	Login(ctx context.Context, data dto.LoginDto) (string, error)
	Register(ctx context.Context, data dto.RegisterDto) (string, error)
}

type service struct {
}

func New() Service {
	return &service{}
}
