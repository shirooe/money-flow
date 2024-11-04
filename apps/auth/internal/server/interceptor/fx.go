package interceptor

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("interceptor",
		fx.Provide(NewAuth, New),
	)
}
