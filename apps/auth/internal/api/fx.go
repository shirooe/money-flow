package api

import (
	desc "github.com/shirooe/money-flow/apps/auth/pkg/grpc"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("api",
		fx.Provide(New),
		fx.Invoke(
			desc.RegisterAuthUnprotectedServer,
		),
	)
}
