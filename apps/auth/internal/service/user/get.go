package user

import (
	"context"

	"github.com/shirooe/money-flow/apps/auth/internal/dto"
)

func (s *service) Get(ctx context.Context, id string) (dto.User, error) {
	return dto.User{}, nil
}
