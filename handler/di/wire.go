package di

import (
	"github.com/google/wire"
	"github.com/resource-aware-jds/worker-lib/handler"
)

var (
	HandlerWireSet = wire.NewSet(
		handler.ProvideGRPCHandler,
	)
)
