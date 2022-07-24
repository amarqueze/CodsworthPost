//go:build wireinject
// +build wireinject

package app

import "github.com/google/wire"

func Setup() *Components {
	wire.Build(
		PostProvider,
		PostPublisher,
		HomeController,
		APIComponent,
	)
	return new(Components)
}
