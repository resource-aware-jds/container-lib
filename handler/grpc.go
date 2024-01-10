package handler

import (
	"fmt"
	"github.com/resource-aware-jds/container-lib/facade"
	"github.com/resource-aware-jds/container-lib/generated/proto/github.com/resource-aware-jds/container-lib/generated/proto"
	"github.com/resource-aware-jds/container-lib/pkg/grpc"
)

type GRPCHandler struct {
	proto.UnimplementedContainerReceiverServer
	containerHandlerFunction facade.ContainerHandlerFunction
}

func ProvideGRPCHandler(grpcServer grpc.Server, containerHandlerFunction facade.ContainerHandlerFunction) (GRPCHandler, error) {
	handler := GRPCHandler{
		containerHandlerFunction: containerHandlerFunction,
	}

	if containerHandlerFunction == nil {
		return handler, fmt.Errorf("container handler function can't be nil")
	}
	proto.RegisterContainerReceiverServer(grpcServer.GetGRPCServer(), handler)
	return handler, nil
}

func Task() {

}
