package http

import "go.uber.org/fx"

var Module = fx.Provide(
	provideClient,
)

func provideClient() Client {
	return NewClient()
}
