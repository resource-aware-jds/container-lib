package di

import "github.com/resource-aware-jds/worker-lib/pkg/grpc"

type App struct {
	GRPCServer grpc.SocketServer
}

func ProvideApp(grpcSocketServer grpc.SocketServer) App {
	return App{
		GRPCServer: grpcSocketServer,
	}
}
