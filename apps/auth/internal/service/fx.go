package service

import (
	"github.com/shirooe/money-flow/apps/auth/internal/service/auth"
	"github.com/shirooe/money-flow/apps/auth/internal/service/user"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("service", fx.Provide(user.New, auth.New))
}
