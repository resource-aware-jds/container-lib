package taskrunnersvc

import (
	"context"
	"fmt"
	"github.com/resource-aware-jds/container-lib/config"
	"github.com/resource-aware-jds/container-lib/facade"
	"github.com/resource-aware-jds/container-lib/generated/proto/github.com/resource-aware-jds/container-lib/generated/proto"
	"github.com/resource-aware-jds/container-lib/model"
	"github.com/resource-aware-jds/container-lib/pkg/containerlibcontext"
	"github.com/resource-aware-jds/container-lib/pkg/grpc"
	"github.com/resource-aware-jds/container-lib/pkg/mapper"
	"github.com/resource-aware-jds/container-lib/pkg/taskrunner"
	"github.com/resource-aware-jds/container-lib/pkg/timeutil"
	"github.com/sirupsen/logrus"
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
	config                      *config.Config
}

type Service interface {
	Run()
	GracefullyShutdown()
	PollTaskFromWorkerNode(ctx context.Context) (*model.Task, error)
}

func ProvideService(config *config.Config, runnerPool taskrunner.Pool, grpcClient grpc.Client, handlerFunc facade.ContainerHandlerFunction) (Service, func()) {
	ctx := context.Background()
	ctxWithCancel, cancelFunc := context.WithCancel(ctx)
	client := proto.NewWorkerNodeContainerReceiverClient(grpcClient.GetConnection())
	result := service{
		runnerPool:           runnerPool,
		workerNodeGRPCClient: client,
		handlerFunc:          handlerFunc,
		ctx:                  ctxWithCancel,
		cancelFunc:           cancelFunc,
		config:               config,
	}
	runnerPool.Subscribe(&result)

	cleanup := func() {
		result.GracefullyShutdown()
	}

	return &result, cleanup
}

func (s *service) Run() {
	logrus.Info("[TaskRunner Manager] Starting the TaskRunner manager loop")
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
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
	s.gracefullyShutdownWaitGroup.Add(1)
	// Check if pool still has some worker left
	if !s.runnerPool.IsAvailableRunner() {
		logrus.Warnf("[TaskRunner Manager] No TaskRunner available in the pool, Skipping this loop.")
		timeutil.SleepWithContext(ctx, 10*time.Second)
		s.gracefullyShutdownWaitGroup.Done()
		return
	}

	// Pull the task
	task, err := s.PollTaskFromWorkerNode(ctx)
	if err != nil {
		logrus.Warnf("[TaskRunner Manager] Failed to poll task from WorkerNode with error %s", err.Error())
		timeutil.SleepWithContext(ctx, 10*time.Second)
		s.gracefullyShutdownWaitGroup.Done()
		return
	}

	// Get Runner from the pool
	runner, err := s.runnerPool.RequestRunner()
	if err != nil {
		logrus.Warnf("[TaskRunner Manager] Failed to get runner from pool with error %s", err.Error())
		_, err = s.workerNodeGRPCClient.ReportTaskFailure(ctx, &proto.ReportTaskFailureRequest{
			ID:          task.ID,
			ErrorDetail: fmt.Sprintf("runner pool error: %s", err.Error()),
		})
		if err != nil {
			// TODO: Create a loop to retry?
			logrus.Warnf("[TaskRunner Manager] Failed to report task failure back to Worker Node with error %s", err.Error())
		}
		timeutil.SleepWithContext(ctx, 10*time.Second)
		s.gracefullyShutdownWaitGroup.Done()
		return
	}

	go s.runTask(ctx, runner, *task)
}

func (s *service) runTask(ctx context.Context, runner taskrunner.Runner, task model.Task) {
	defer s.gracefullyShutdownWaitGroup.Done()
	logger := logrus.WithFields(logrus.Fields{
		"taskID":   task.ID,
		"jobID":    task.JobID,
		"runnerID": runner.GetID(),
	})

	logger.Debugf("[TaskRunner Manager] TaskRunner GoRoutine started")
	internalContext := containerlibcontext.ProvideContext(ctx)
	innerErr := runner.Run(internalContext, s.handlerFunc, task)

	// Always return runner to the pool
	defer s.runnerPool.ReturnRunner(runner)

	if innerErr == nil && internalContext.GetSuccessResult() {
		// Report Success
		_, err := s.workerNodeGRPCClient.SubmitSuccessTask(context.Background(), &proto.SubmitSuccessTaskRequest{
			ID:      task.ID,
			Results: internalContext.GetResults(),
		})
		if err != nil {
			logger.Errorf("[TaskRunner Manager] Fail to submit the success task with error (%s)", err.Error())
			// TODO: Create a retry?
			return
		}
		return
	}

	errMsg := ""
	if innerErr != nil {
		errMsg = innerErr.Error()
	}

	logger.Warnf("[TaskRunner Manager] Handler Function didn't call the containerlibcontext.Success() method and report this error (%s)", innerErr)
	// Report Failure
	_, err := s.workerNodeGRPCClient.ReportTaskFailure(context.Background(), &proto.ReportTaskFailureRequest{
		ID:          task.ID,
		ErrorDetail: errMsg,
	})
	if err != nil {
		logger.Errorf("[TaskRunner Manager] Fail to submit the failure task with error (%s)", err.Error())
		// TODO: Create a retry?
		return
	}
}

func (s *service) OnEvent(e taskrunner.PoolEvent) {
	if s.ctx.Err() != nil {
		logrus.Debugf("[TaskRunner Manager] Ignoring the TaskRunnerPool event due to context error (%s)", s.ctx.Err())
		return
	}
	s.loopRoutine(s.ctx)
}

func (s *service) PollTaskFromWorkerNode(ctx context.Context) (*model.Task, error) {
	result, err := s.workerNodeGRPCClient.GetTaskFromQueue(ctx, &proto.GetTaskPayload{
		ImageUrl:    s.config.ImageURL,
		ContainerId: s.config.ContainerId,
	})
	if err != nil {
		return nil, err
	}

	return mapper.ConvertTaskProtoToModel(result)
}
