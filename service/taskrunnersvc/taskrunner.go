package taskrunnersvc

import (
	"context"
	"fmt"
	"github.com/resource-aware-jds/container-lib/facade"
	"github.com/resource-aware-jds/container-lib/generated/proto/github.com/resource-aware-jds/container-lib/generated/proto"
	"github.com/resource-aware-jds/container-lib/model"
	"github.com/resource-aware-jds/container-lib/pkg/grpc"
	"github.com/resource-aware-jds/container-lib/pkg/mapper"
	"github.com/resource-aware-jds/container-lib/pkg/taskrunner"
	"github.com/sirupsen/logrus"
	"time"
)

type service struct {
	runnerPool           taskrunner.Pool
	workerNodeGRPCClient proto.WorkerNodeContainerReceiverClient
	handlerFunc          facade.ContainerHandlerFunction
}

type Service interface {
	Loop(ctx context.Context)
	PollTaskFromWorkerNode(ctx context.Context) (*model.Task, error)
}

func ProvideService(runnerPool taskrunner.Pool, grpcClient grpc.SocketClient, handlerFunc facade.ContainerHandlerFunction) Service {
	client := proto.NewWorkerNodeContainerReceiverClient(grpcClient.GetConnection())
	result := service{
		runnerPool:           runnerPool,
		workerNodeGRPCClient: client,
		handlerFunc:          handlerFunc,
	}
	runnerPool.Subscribe(&result)

	return &result
}

func (s *service) Loop(ctx context.Context) {
	go func(innerCtx context.Context) {
		for {
			select {
			case <-ctx.Done():
			// Graceful shutdown
			default:
				s.loopRoutine(ctx)
			}
		}
	}(ctx)
}

func (s *service) loopRoutine(ctx context.Context) {
	// Check if pool still has some worker left
	if !s.runnerPool.IsAvailableRunner() {
		time.Sleep(10 * time.Second)
		return
	}

	task, err := s.PollTaskFromWorkerNode(ctx)
	if err != nil {
		time.Sleep(10 * time.Second)
		return
	}

	// Pull the task
	runner, err := s.runnerPool.RequestRunner()
	if err != nil {
		_, err = s.workerNodeGRPCClient.ReportTaskFailure(ctx, &proto.ReportTaskFailureRequest{
			ID:          task.ID.GetRawTaskID(),
			ErrorDetail: fmt.Sprintf("runner pool error: %s", err.Error()),
		})
		if err != nil {
			// TODO: Create a loop to retry?
			logrus.Error("ReportTaskFailure failed: ", err)
		}
		time.Sleep(10 * time.Second)
		return
	}

	go func(innerCtx context.Context, innerRunner taskrunner.Runner, handlerFunc facade.ContainerHandlerFunction, innerTask model.Task) {
		internalContext, innerErr := innerRunner.Run(innerCtx, handlerFunc, innerTask)

		// Always return runner to the pool
		defer s.runnerPool.ReturnRunner(innerRunner)

		if innerErr == nil {
			// Report Success
			_, err := s.workerNodeGRPCClient.SubmitSuccessTask(innerCtx, &proto.SubmitSuccessTaskRequest{
				ID:      innerTask.ID.GetRawTaskID(),
				Results: internalContext.GetResults(),
			})
			if err != nil {
				// TODO: Create a retry?
				return
			}
			return
		} else {
			// Report Success
			_, err := s.workerNodeGRPCClient.ReportTaskFailure(innerCtx, &proto.ReportTaskFailureRequest{
				ID:          innerTask.ID.GetRawTaskID(),
				ErrorDetail: innerErr.Error(),
			})
			if err != nil {
				// TODO: Create a retry?
				return
			}
			return
		}

	}(ctx, runner, s.handlerFunc, *task)

}

func (s *service) OnEvent(e taskrunner.PoolEvent) {
	//TODO implement me
	panic("implement me")
}

func (s *service) PollTaskFromWorkerNode(ctx context.Context) (*model.Task, error) {
	result, err := s.workerNodeGRPCClient.GetTaskFromQueue(ctx, nil)
	if err != nil {
		return nil, err
	}

	return mapper.ConvertTaskProtoToModel(result)
}
