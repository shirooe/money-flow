package user

import (
	"context"

	"github.com/shirooe/money-flow/apps/auth/internal/dto"
)

func (s *service) Update(ctx context.Context, data dto.UpdateUser) (dto.User, error) {
	return dto.User{}, nil
}
