package repository

import (
	"github.com/shirooe/money-flow/apps/auth/internal/repository/user"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("repository",
		fx.Provide(user.New),
	)
}
