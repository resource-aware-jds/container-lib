package di

import (
	"github.com/google/wire"
	"github.com/resource-aware-jds/worker-lib/pkg/grpc"
)

var (
	PKGWireSet = wire.NewSet(
		grpc.ProvideGRPCSocketServer,
	)
)
