package taskrunnersvc

import (
	"context"
	"errors"
	"fmt"
	"github.com/resource-aware-jds/container-lib/config"
	"github.com/resource-aware-jds/container-lib/facade"
	"github.com/resource-aware-jds/container-lib/generated/proto/github.com/resource-aware-jds/container-lib/generated/proto"
	"github.com/resource-aware-jds/container-lib/model"
	"github.com/resource-aware-jds/container-lib/pkg/containerlibcontext"
	"github.com/resource-aware-jds/container-lib/pkg/grpc"
	"github.com/resource-aware-jds/container-lib/pkg/taskrunner"
	"github.com/resource-aware-jds/container-lib/pkg/timeutil"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var (
	ErrNoAvailableRunner = errors.New("no available runner")
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
	RunTask(ctx context.Context, task model.Task) error
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

func (s *service) RunTask(ctx context.Context, task model.Task) error {
	s.gracefullyShutdownWaitGroup.Add(1)
	// Check if pool still has some worker left
	if !s.runnerPool.IsAvailableRunner() {
		logrus.Warnf("[TaskRunner Manager] No TaskRunner available in the pool, Skipping this loop.")
		timeutil.SleepWithContext(ctx, 10*time.Second)
		s.gracefullyShutdownWaitGroup.Done()
		return ErrNoAvailableRunner
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
		return err
	}

	go s.runTask(ctx, runner, task)
	return nil
}

func (s *service) GracefullyShutdown() {
	logrus.Info("[TaskRunner Manager] Gracefully Shutting down signal received")
	s.cancelFunc()
	s.gracefullyShutdownWaitGroup.Wait()
	logrus.Info("[TaskRunner Manager] Gracefully Shutdown success.")
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

	// TODO: Notify the worker node about the runner availability.
}
