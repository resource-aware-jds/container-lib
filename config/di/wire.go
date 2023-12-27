package di

import (
	"github.com/google/wire"
	"github.com/resource-aware-jds/worker-lib/config"
)

var (
	ConfigWireSet = wire.NewSet(
		config.ProvideConfig,
		config.ProvideGRPCSocketServerConfig,
	)
)
