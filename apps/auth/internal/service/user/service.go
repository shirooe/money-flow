package user

import (
	"context"

	"github.com/shirooe/money-flow/apps/auth/internal/dto"
)

type Service interface {
	Create(ctx context.Context, data dto.CreateUser) (string, error)
	Get(ctx context.Context, id string) (dto.User, error)
	Update(ctx context.Context, data dto.UpdateUser) (dto.User, error)
}

type service struct {
}

func New() Service {
	return &service{}
}
