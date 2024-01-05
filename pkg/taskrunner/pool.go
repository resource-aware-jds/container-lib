package taskrunner

import (
	"fmt"
	"github.com/resource-aware-jds/container-lib/pkg/datastructure"
	"github.com/resource-aware-jds/container-lib/pkg/observer"
	"github.com/sirupsen/logrus"
	"strconv"
	"sync"
)

type PoolEvent any

type pool struct {
	runnerMu    sync.Mutex
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
	NumberOfInitialTaskRunner int
}

func ProvideTaskRunnerPool(config PoolConfig) (Pool, error) {
	runnerQueue := datastructure.ProvideQueue[Runner](config.NumberOfInitialTaskRunner)

	for i := 0; i < config.NumberOfInitialTaskRunner; i++ {
		createdWorker := ProvideRunner(strconv.Itoa(i))
		runnerQueue.Push(createdWorker)
	}

	logrus.Infof("[TaskRunnerPool] Initialized with %d runner(s)", config.NumberOfInitialTaskRunner)
	return &pool{
		runnerQueue: runnerQueue,
	}, nil
}

func (p *pool) IsAvailableRunner() bool {
	p.runnerMu.Lock()
	defer p.runnerMu.Unlock()

	return !p.runnerQueue.Empty()
}

func (p *pool) RequestRunner() (Runner, error) {
	if !p.IsAvailableRunner() {
		logrus.Error("[TaskRunnerPool] No available runner in the pool")
		return nil, fmt.Errorf("no available runner")
	}

	p.runnerMu.Lock()
	poppedRunner, ok := p.runnerQueue.Pop()
	p.runnerMu.Unlock()
	if !ok {
		logrus.Error("[TaskRunnerPool] Failed to get runner")
		return nil, fmt.Errorf("failed to get runner")
	}

	dePointerPoppedRunner := *poppedRunner
	logrus.WithFields(logrus.Fields{
		"runnerID": dePointerPoppedRunner.GetID(),
	}).Info("[TaskRunnerPool] A Runner has been sent out from the pool")
	return dePointerPoppedRunner, nil
}

func (p *pool) ReturnRunner(w Runner) {

	logrus.WithFields(logrus.Fields{
		"runnerID": w.GetID(),
	}).Info("[TaskRunnerPool] A Runner has been returned to the pool")
	p.runnerMu.Lock()
	p.runnerQueue.Push(w)
	p.runnerMu.Unlock()

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
