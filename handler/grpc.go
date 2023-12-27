package handler

import (
	"github.com/resource-aware-jds/worker-lib/generated/proto/github.com/resource-aware-jds/worker-lib/generated/proto"
	"github.com/resource-aware-jds/worker-lib/pkg/grpc"
)

type GRPCHandler struct {
	proto.UnimplementedWorkerContainerServer
}

func ProvideGRPCHandler(grpcServer grpc.SocketServer) GRPCHandler {
	handler := GRPCHandler{}
	proto.RegisterWorkerContainerServer(grpcServer.GetGRPCServer(), handler)
	return handler
}
