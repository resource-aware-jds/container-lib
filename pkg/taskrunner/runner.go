package taskrunner

import (
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
	Run(ctx containerlibcontext.Context, handlerFunc facade.ContainerHandlerFunction, task model.Task) error
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

func (w runner) Run(ctx containerlibcontext.Context, handlerFunc facade.ContainerHandlerFunction, task model.Task) error {
	w.logger.Info("[TaskRunner] Starting to work on task ", task.ID.GetRawTaskID())

	err := handlerFunc(ctx, task)
	if err != nil {
		w.logger.Error(err)
	}
	return err
}
