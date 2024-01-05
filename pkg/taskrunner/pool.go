package taskrunner

import (
	"fmt"
	"github.com/resource-aware-jds/container-lib/pkg/datastructure"
	"github.com/resource-aware-jds/container-lib/pkg/observer"
	"github.com/sirupsen/logrus"
	"strconv"
)

type PoolEvent any

type pool struct {
	runnerQueue datastructure.Queue[Runner]
	subscribers []observer.Subscriber[PoolEvent]
}

type Pool interface {
	observer.Publisher[PoolEvent]

	IsAvailableRunner() bool
	RequestRunner() (Runner, error)
	ReturnRunner(Runner)
}

type PoolConfig struct {
	NumberOfWorker int
}

func ProvideTaskRunnerPool(config PoolConfig) (Pool, error) {
	runnerQueue := datastructure.ProvideQueue[Runner](config.NumberOfWorker)

	for i := 0; i < config.NumberOfWorker; i++ {
		createdWorker := ProvideRunner(strconv.Itoa(i))
		runnerQueue.Push(createdWorker)
	}

	logrus.Infof("[TaskRunnerPool] Initialized with %d runner(s)", config.NumberOfWorker)
	return &pool{
		runnerQueue: runnerQueue,
	}, nil
}

func (p *pool) IsAvailableRunner() bool {
	return !p.runnerQueue.Empty()
}

func (p *pool) RequestRunner() (Runner, error) {
	if !p.IsAvailableRunner() {
		logrus.Error("[TaskRunnerPool] No available runner in the pool")
		return nil, fmt.Errorf("no available runner")
	}

	poppedRunner, ok := p.runnerQueue.Pop()
	if !ok {
		logrus.Error("[TaskRunnerPool] Failed to get runner")
		return nil, fmt.Errorf("failed to get runner")
	}

	dePointerPoppedRunner := *poppedRunner
	logrus.Infof("[TaskRunnerPool] Runner %s has been sent out from the pool", dePointerPoppedRunner.GetID())
	return dePointerPoppedRunner, nil
}

func (p *pool) ReturnRunner(w Runner) {
	logrus.Infof("[TaskRunnerPool] Runner %s has been returned to the pool", w.GetID())
	p.runnerQueue.Push(w)

	// Notify the subscriber about the available runner
	p.NotifySubscribers(nil)
}

func (p *pool) Subscribe(subscriber observer.Subscriber[PoolEvent]) {
	p.subscribers = append(p.subscribers, subscriber)
}

func (p *pool) Unsubscribe(subscriber observer.Subscriber[PoolEvent]) {
	//TODO Let see if we have this use case, then I'll implement this later
	panic("implement me")
}

func (p *pool) NotifySubscribers(e PoolEvent) {
	logrus.Debug("[TaskRunnerPool] Notifying the Subscriber about the available runner")
	for _, subscriber := range p.subscribers {
		subscriber.OnEvent(e)
	}
}
