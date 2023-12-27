package workerpool

import (
	"fmt"
	"github.com/resource-aware-jds/worker-lib/pkg/datastructure"
	"github.com/resource-aware-jds/worker-lib/pkg/observer"
	"strconv"
)

type PoolEvent any

type pool struct {
	workerQueue datastructure.Queue[*Worker]
	subscribers []observer.Subscriber[PoolEvent]
}

type Pool interface {
	observer.Publisher[PoolEvent]

	IsAvailableWorker() bool
	RequesterWorker() (*Worker, error)
	ReturnWorker(*Worker)
}

func ProvideWorkerPool(numberOfWorker int) (Pool, error) {
	workerQueue := datastructure.ProvideQueue[*Worker](numberOfWorker)

	for i := 0; i < numberOfWorker; i++ {
		createdWorker := ProvideWorker(strconv.Itoa(i))
		workerQueue.Push(&createdWorker)
	}

	return &pool{
		workerQueue: workerQueue,
	}, nil
}

func (p *pool) IsAvailableWorker() bool {
	return !p.workerQueue.Empty()
}

func (p *pool) RequesterWorker() (*Worker, error) {
	if !p.IsAvailableWorker() {
		return nil, fmt.Errorf("no available worker")
	}

	poppedWorker, ok := p.workerQueue.Pop()
	if !ok {
		return nil, fmt.Errorf("failed to get worker")
	}

	return *poppedWorker, nil
}

func (p *pool) ReturnWorker(w *Worker) {
	p.workerQueue.Push(w)
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
