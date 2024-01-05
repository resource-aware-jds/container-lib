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
	"google.golang.org/protobuf/types/known/emptypb"
	"sync"
	"time"
)

type service struct {
	runnerPool                  taskrunner.Pool
	workerNodeGRPCClient        proto.WorkerNodeContainerReceiverClient
	handlerFunc                 facade.ContainerHandlerFunction
	ctx                         context.Context
	cancelFunc                  func()
	gracefullyShutdownWaitGroup sync.WaitGroup
}

type Service interface {
	Run()
	GracefullyShutdown()
	PollTaskFromWorkerNode(ctx context.Context) (*model.Task, error)
}

func ProvideService(runnerPool taskrunner.Pool, grpcClient grpc.SocketClient, handlerFunc facade.ContainerHandlerFunction) (Service, func()) {
	ctx := context.Background()
	ctxWithCancel, cancelFunc := context.WithCancel(ctx)
	client := proto.NewWorkerNodeContainerReceiverClient(grpcClient.GetConnection())
	result := service{
		runnerPool:           runnerPool,
		workerNodeGRPCClient: client,
		handlerFunc:          handlerFunc,
		ctx:                  ctxWithCancel,
		cancelFunc:           cancelFunc,
	}
	runnerPool.Subscribe(&result)

	cleanup := func() {
		result.GracefullyShutdown()
	}

	return &result, cleanup
}

func (s *service) Run() {
	logrus.Info("[TaskRunner Manager] Starting the TaskRunner manager loop")
	s.gracefullyShutdownWaitGroup.Add(1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				s.gracefullyShutdownWaitGroup.Done()
				return
			default:
				s.loopRoutine(ctx)
			}
		}
	}(s.ctx)
}

func (s *service) GracefullyShutdown() {
	logrus.Info("[TaskRunner Manager] Gracefully Shutting down signal received")
	s.cancelFunc()
	s.gracefullyShutdownWaitGroup.Wait()
	logrus.Info("[TaskRunner Manager] Gracefully Shutdown success.")
}

func (s *service) loopRoutine(ctx context.Context) {
	// Check if pool still has some worker left
	if !s.runnerPool.IsAvailableRunner() {
		logrus.Warnf("[TaskRunner Manager] No TaskRunner available in the pool, Skipping this loop.")
		time.Sleep(10 * time.Second)
		return
	}

	// Pull the task
	task, err := s.PollTaskFromWorkerNode(ctx)
	if err != nil {
		logrus.Warnf("[TaskRunner Manager] Failed to poll task from WorkerNode with error %s", err.Error())
		time.Sleep(10 * time.Second)
		return
	}

	// Get Runner from the pool
	runner, err := s.runnerPool.RequestRunner()
	if err != nil {
		logrus.Warnf("[TaskRunner Manager] Failed to get runner from pool with error %s", err.Error())
		_, err = s.workerNodeGRPCClient.ReportTaskFailure(ctx, &proto.ReportTaskFailureRequest{
			ID:          task.ID.GetRawTaskID(),
			ErrorDetail: fmt.Sprintf("runner pool error: %s", err.Error()),
		})
		if err != nil {
			// TODO: Create a loop to retry?
			logrus.Warnf("[TaskRunner Manager] Failed to report task failure back to Worker Node with error %s", err.Error())
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
	logrus.Debugf("[TaskRunner Manager] TaskRunner GoRoutine started")
}

func (s *service) OnEvent(e taskrunner.PoolEvent) {
	s.loopRoutine(s.ctx)
}

func (s *service) PollTaskFromWorkerNode(ctx context.Context) (*model.Task, error) {
	result, err := s.workerNodeGRPCClient.GetTaskFromQueue(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	return mapper.ConvertTaskProtoToModel(result)
}
