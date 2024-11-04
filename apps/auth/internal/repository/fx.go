package repository

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("repository", fx.Provide())
}
