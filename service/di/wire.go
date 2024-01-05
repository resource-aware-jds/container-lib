package di

import (
	"github.com/google/wire"
	"github.com/resource-aware-jds/container-lib/service/taskrunnersvc"
)

var (
	ServiceWireSet = wire.NewSet(taskrunnersvc.ProvideService)
)
