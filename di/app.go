package di

import (
	"github.com/resource-aware-jds/container-lib/handler"
	"github.com/resource-aware-jds/container-lib/pkg/grpc"
)

type App struct {
	GRPCServer  grpc.SocketServer
	GRPCHandler handler.GRPCHandler
}

func ProvideApp(grpcSocketServer grpc.SocketServer, grpcHandler handler.GRPCHandler) App {
	return App{
		GRPCServer:  grpcSocketServer,
		GRPCHandler: grpcHandler,
	}
}
