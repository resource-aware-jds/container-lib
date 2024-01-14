package handler

import (
	"context"
	"fmt"
	"github.com/resource-aware-jds/container-lib/facade"
	"github.com/resource-aware-jds/container-lib/generated/proto/github.com/resource-aware-jds/container-lib/generated/proto"
	"github.com/resource-aware-jds/container-lib/model"
	"github.com/resource-aware-jds/container-lib/pkg/grpc"
	"github.com/resource-aware-jds/container-lib/pkg/mapper"
	"github.com/resource-aware-jds/container-lib/service/taskrunnersvc"
)

type GRPCHandler struct {
	proto.UnimplementedContainerReceiverServer
	containerHandlerFunction facade.ContainerHandlerFunction
	taskRunnerService        taskrunnersvc.Service
}

func ProvideGRPCHandler(grpcServer grpc.Server, taskRunnerService taskrunnersvc.Service, containerHandlerFunction facade.ContainerHandlerFunction) (GRPCHandler, error) {
	handler := GRPCHandler{
		containerHandlerFunction: containerHandlerFunction,
		taskRunnerService:        taskRunnerService,
	}

	if containerHandlerFunction == nil {
		return handler, fmt.Errorf("container handler function can't be nil")
	}
	proto.RegisterContainerReceiverServer(grpcServer.GetGRPCServer(), &handler)
	return handler, nil
}

func (g *GRPCHandler) SendTasks(ctx context.Context, req *proto.SendTasksRequest) (*proto.SendTasksResponse, error) {
	convertedTasks := make([]model.Task, 0, len(req.GetTasks()))

	response := &proto.SendTasksResponse{
		FailureTasks: make([]*proto.RunTaskFailureResponse, 0, len(req.GetTasks())),
	}
	for _, taskReq := range req.GetTasks() {
		taskModel, err := mapper.ConvertTaskProtoToModel(taskReq)
		if err != nil {
			response.FailureTasks = append(response.FailureTasks, &proto.RunTaskFailureResponse{
				FailureTaskId: taskModel.ID,
				Error:         err.Error(),
			})
			continue
		}
		convertedTasks = append(convertedTasks, *taskModel)
	}

	for _, convertedTask := range convertedTasks {
		err := g.taskRunnerService.RunTask(ctx, convertedTask)
		if err != nil {
			response.FailureTasks = append(response.FailureTasks, &proto.RunTaskFailureResponse{
				FailureTaskId: convertedTask.ID,
				Error:         err.Error(),
			})
			return nil, err
		}
	}

	return response, nil
}
