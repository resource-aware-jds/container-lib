package di

import (
	"github.com/google/wire"
	"github.com/resource-aware-jds/container-lib/pkg/grpc"
	"github.com/resource-aware-jds/container-lib/pkg/taskrunner"
)

var (
	PKGWireSet = wire.NewSet(
		grpc.ProvideGRPCServer,
		grpc.ProvideClient,
		taskrunner.ProvideTaskRunnerPool,
	)
)
