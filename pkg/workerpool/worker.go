package workerpool

import "github.com/resource-aware-jds/worker-lib/model"

type worker struct {
	id     string
	status string
}

type Worker interface {
	GetID() string
	Run(handlerFunc model.WorkerHandlerFunc) error
}

func ProvideWorker(id string) Worker {
	return &worker{
		id: id,
	}
}

func (w worker) GetID() string {
	return w.id
}

func (w worker) Run(handlerFunc model.WorkerHandlerFunc) error {
	//TODO implement me
	panic("implement me")
}
