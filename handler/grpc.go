package handler

import (
	"fmt"
	"github.com/resource-aware-jds/worker-lib/facade"
	"github.com/resource-aware-jds/worker-lib/generated/proto/github.com/resource-aware-jds/worker-lib/generated/proto"
	"github.com/resource-aware-jds/worker-lib/pkg/grpc"
)

type GRPCHandler struct {
	proto.UnimplementedWorkerContainerServer
	containerHandlerFunction facade.ContainerHandlerFunction
}

func ProvideGRPCHandler(grpcServer grpc.SocketServer, containerHandlerFunction facade.ContainerHandlerFunction) (GRPCHandler, error) {
	handler := GRPCHandler{
		containerHandlerFunction: containerHandlerFunction,
	}

	if containerHandlerFunction == nil {
		return handler, fmt.Errorf("container handler function can't be nil")
	}
	proto.RegisterWorkerContainerServer(grpcServer.GetGRPCServer(), handler)
	return handler, nil
}

func Task() {

}
