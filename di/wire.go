//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	configDI "github.com/resource-aware-jds/container-lib/config/di"
	"github.com/resource-aware-jds/container-lib/facade"
	handlerDI "github.com/resource-aware-jds/container-lib/handler/di"
	pkgDI "github.com/resource-aware-jds/container-lib/pkg/di"
	serviceDI "github.com/resource-aware-jds/container-lib/service/di"
)

//go:generate wire

func InitializeApplication(containerHandlerFunction facade.ContainerHandlerFunction) (App, func(), error) {
	panic(
		wire.Build(
			ProvideApp,
			configDI.ConfigWireSet,
			pkgDI.PKGWireSet,
			handlerDI.HandlerWireSet,
			serviceDI.ServiceWireSet,
		),
	)
}
