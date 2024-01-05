package taskrunner

import (
	"context"
	"github.com/resource-aware-jds/container-lib/facade"
	"github.com/resource-aware-jds/container-lib/model"
	"github.com/resource-aware-jds/container-lib/pkg/containerlibcontext"
)

type runner struct {
	id         string
	status     string
	errChannel chan<- error
}

type Runner interface {
	GetID() string
	Run(ctx context.Context, handlerFunc facade.ContainerHandlerFunction, task model.Task) (containerlibcontext.Context, error)
}

func ProvideRunner(id string) Runner {
	return &runner{
		id: id,
	}
}

func (w runner) GetID() string {
	return w.id
}

func (w runner) Run(ctx context.Context, handlerFunc facade.ContainerHandlerFunction, task model.Task) (containerlibcontext.Context, error) {
	internalCtx := containerlibcontext.ProvideContext(ctx)
	return internalCtx, handlerFunc(internalCtx, task)
}
