package taskrunner

import (
	"fmt"
	"github.com/resource-aware-jds/container-lib/pkg/datastructure"
	"github.com/resource-aware-jds/container-lib/pkg/observer"
	"strconv"
)

type PoolEvent any

type pool struct {
	runnerQueue datastructure.Queue[*Runner]
	subscribers []observer.Subscriber[PoolEvent]
}

type Pool interface {
	observer.Publisher[PoolEvent]

	IsAvailableRunner() bool
	RequestRunner() (*Runner, error)
	ReturnRunner(*Runner)
}

type PoolConfig struct {
	NumberOfWorker int
}

func ProvideTaskRunnerPool(config PoolConfig) (Pool, error) {
	runnerQueue := datastructure.ProvideQueue[*Runner](config.NumberOfWorker)

	for i := 0; i < config.NumberOfWorker; i++ {
		createdWorker := ProvideRunner(strconv.Itoa(i))
		runnerQueue.Push(&createdWorker)
	}

	return &pool{
		runnerQueue: runnerQueue,
	}, nil
}

func (p *pool) IsAvailableRunner() bool {
	return !p.runnerQueue.Empty()
}

func (p *pool) RequestRunner() (*Runner, error) {
	if !p.IsAvailableRunner() {
		return nil, fmt.Errorf("no available worker")
	}

	poppedWorker, ok := p.runnerQueue.Pop()
	if !ok {
		return nil, fmt.Errorf("failed to get worker")
	}

	return *poppedWorker, nil
}

func (p *pool) ReturnRunner(w *Runner) {
	p.runnerQueue.Push(w)
}

func (p *pool) Subscribe(subscriber observer.Subscriber[PoolEvent]) {
	p.subscribers = append(p.subscribers, subscriber)
}

func (p *pool) Unsubscribe(subscriber observer.Subscriber[PoolEvent]) {
	//TODO Let see if we have this use case, then I'll implement this later
	panic("implement me")
}

func (p *pool) NotifySubscribers(e PoolEvent) {
	for _, subscriber := range p.subscribers {
		subscriber.OnEvent(e)
	}
}
