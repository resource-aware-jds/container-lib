package di

import (
	"github.com/resource-aware-jds/container-lib/handler"
	"github.com/resource-aware-jds/container-lib/pkg/grpc"
	"github.com/resource-aware-jds/container-lib/service/taskrunnersvc"
)

type App struct {
	GRPCServer    grpc.Server
	GRPCHandler   handler.GRPCHandler
	TaskRunnerSvc taskrunnersvc.Service
}

func ProvideApp(grpcServer grpc.Server, grpcHandler handler.GRPCHandler, taskRunnerSvc taskrunnersvc.Service) App {
	return App{
		GRPCServer:    grpcServer,
		GRPCHandler:   grpcHandler,
		TaskRunnerSvc: taskRunnerSvc,
	}
}
