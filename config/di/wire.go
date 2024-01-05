package di

import (
	"github.com/google/wire"
	"github.com/resource-aware-jds/container-lib/config"
)

var (
	ConfigWireSet = wire.NewSet(
		config.ProvideConfig,
		config.ProvideGRPCSocketServerConfig,
		config.ProvideTaskRunnerPoolConfig,
		config.ProvideGRPCSocketClientConfig,
	)
)
