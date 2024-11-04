package unprotected

import (
	desc "github.com/shirooe/money-flow/apps/auth/pkg/grpc"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("unprotected",
		fx.Provide(New),
		fx.Invoke(desc.RegisterAuthUnprotectedServer),
	)
}
