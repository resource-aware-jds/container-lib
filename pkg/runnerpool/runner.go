package runnerpool

import (
	"context"
	"github.com/resource-aware-jds/worker-lib/model"
	"github.com/resource-aware-jds/worker-lib/pkg/workerlibcontext"
	"github.com/resource-aware-jds/worker-lib/test"
)

type runner struct {
	id         string
	status     string
	errChannel chan<- error
}

type Runner interface {
	GetID() string
	Run(ctx context.Context, handlerFunc test.WorkerHandlerFunc, task model.Task) error
}

func ProvideRunner(id string) Runner {
	return &runner{
		id: id,
	}
}

func (w runner) GetID() string {
	return w.id
}

func (w runner) Run(ctx context.Context, handlerFunc test.WorkerHandlerFunc, task model.Task) error {
	internalCtx := workerlibcontext.ProvideContext(ctx, task)

	go func(innerCtx workerlibcontext.Context) {
		err := handlerFunc(internalCtx)
		w.errChannel <- err
	}(internalCtx)

	return nil
}
