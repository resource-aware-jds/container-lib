package di

import (
	"github.com/resource-aware-jds/container-lib/handler"
	"github.com/resource-aware-jds/container-lib/pkg/grpc"
	"github.com/resource-aware-jds/container-lib/service/taskrunnersvc"
)

type App struct {
	GRPCServer    grpc.SocketServer
	GRPCHandler   handler.GRPCHandler
	TaskRunnerSvc taskrunnersvc.Service
}

func ProvideApp(grpcSocketServer grpc.SocketServer, grpcHandler handler.GRPCHandler, taskRunnerSvc taskrunnersvc.Service) App {
	return App{
		GRPCServer:    grpcSocketServer,
		GRPCHandler:   grpcHandler,
		TaskRunnerSvc: taskRunnerSvc,
	}
}
