//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	configDI "github.com/resource-aware-jds/worker-lib/config/di"
	pkgDI "github.com/resource-aware-jds/worker-lib/pkg/di"
)

//go:generate wire

func InitializeApplication() (App, func(), error) {
	panic(
		wire.Build(
			ProvideApp,
			configDI.ConfigWireSet,
			pkgDI.PKGWireSet,
		),
	)
}
