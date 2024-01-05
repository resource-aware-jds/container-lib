package taskrunner

import (
	"context"
	"github.com/resource-aware-jds/container-lib/facade"
	"github.com/resource-aware-jds/container-lib/model"
	"github.com/resource-aware-jds/container-lib/pkg/containerlibcontext"
	"github.com/sirupsen/logrus"
)

type runner struct {
	id     string
	logger *logrus.Entry
}

type Runner interface {
	GetID() string
	Run(ctx context.Context, handlerFunc facade.ContainerHandlerFunction, task model.Task) (containerlibcontext.Context, error)
}

func ProvideRunner(id string) Runner {
	logrus.Debugf("[TaskRunner] Runner %s has been initialized", id)
	return &runner{
		id: id,
		logger: logrus.WithFields(logrus.Fields{
			"runnerID": id,
		}),
	}
}

func (w runner) GetID() string {
	return w.id
}

func (w runner) Run(ctx context.Context, handlerFunc facade.ContainerHandlerFunction, task model.Task) (containerlibcontext.Context, error) {
	w.logger.Info("[TaskRunner] Starting to work on task ", task.ID.GetRawTaskID())
	internalCtx := containerlibcontext.ProvideContext(ctx)

	err := handlerFunc(internalCtx, task)
	if err != nil {
		w.logger.Error(err)
	}
	return internalCtx, err
}
