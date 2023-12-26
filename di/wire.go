//go:build wireinject
// +build wireinject

package di

import "github.com/google/wire"

//go:generate wire

func InitializeApplication() (App, func(), error) {
	panic(
		wire.Build(
			ProvideApp,
		),
	)
}
