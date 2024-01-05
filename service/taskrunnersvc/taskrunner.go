package taskrunnersvc

import (
	"context"
	"github.com/resource-aware-jds/container-lib/generated/proto/github.com/resource-aware-jds/container-lib/generated/proto"
	"github.com/resource-aware-jds/container-lib/model"
	"github.com/resource-aware-jds/container-lib/pkg/grpc"
	"github.com/resource-aware-jds/container-lib/pkg/mapper"
	"github.com/resource-aware-jds/container-lib/pkg/taskrunner"
)

type service struct {
	runnerPool           taskrunner.Pool
	workerNodeGRPCClient proto.WorkerNodeContainerReceiverClient
}

type Service interface {
	Loop(ctx context.Context)
	PollTaskFromWorkerNode(ctx context.Context) (*model.Task, error)
}

func ProvideService(runnerPool taskrunner.Pool, grpcClient grpc.SocketClient) Service {
	client := proto.NewWorkerNodeContainerReceiverClient(grpcClient.GetConnection())
	result := service{
		runnerPool:           runnerPool,
		workerNodeGRPCClient: client,
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
			}
		}
	}(ctx)
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
