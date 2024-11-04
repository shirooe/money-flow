package api

import (
	"github.com/shirooe/money-flow/apps/auth/internal/api/protected"
	"github.com/shirooe/money-flow/apps/auth/internal/api/unprotected"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("api",
		fx.Options(protected.Module(), unprotected.Module()),
	)
}
