package handler

import (
	"fmt"
	"github.com/resource-aware-jds/worker-lib/generated/proto/github.com/resource-aware-jds/worker-lib/generated/proto"
	"github.com/resource-aware-jds/worker-lib/model"
	"github.com/resource-aware-jds/worker-lib/pkg/grpc"
)

type GRPCHandler struct {
	proto.UnimplementedWorkerContainerServer
	workerHandlerFunc model.WorkerHandlerFunc
}

func ProvideGRPCHandler(grpcServer grpc.SocketServer, workerHandlerFunc model.WorkerHandlerFunc) (GRPCHandler, error) {
	handler := GRPCHandler{
		workerHandlerFunc: workerHandlerFunc,
	}

	if workerHandlerFunc == nil {
		return handler, fmt.Errorf("WorkerHandlerFunc can't be nil")
	}
	proto.RegisterWorkerContainerServer(grpcServer.GetGRPCServer(), handler)
	return handler, nil
}
