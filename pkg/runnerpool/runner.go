package runnerpool

import (
	"context"
	"github.com/resource-aware-jds/container-lib/facade"
	"github.com/resource-aware-jds/container-lib/model"
)

type runner struct {
	id         string
	status     string
	errChannel chan<- error
}

type Runner interface {
	GetID() string
	Run(ctx context.Context, handlerFunc facade.ContainerHandlerFunction, task model.Task) error
}

func ProvideRunner(id string) Runner {
	return &runner{
		id: id,
	}
}

func (w runner) GetID() string {
	return w.id
}

func (w runner) Run(ctx context.Context, handlerFunc facade.ContainerHandlerFunction, task model.Task) error {
	internalCtx := containerlibcontext.ProvideContext(ctx, task)

	go func(innerCtx containerlibcontext.Context) {
		err := handlerFunc(internalCtx)
		w.errChannel <- err
	}(internalCtx)

	return nil
}
