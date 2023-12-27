package workerlib

import (
	"github.com/resource-aware-jds/worker-lib/di"
	"github.com/resource-aware-jds/worker-lib/pkg/workerlibcontext"
)

type WorkerHandlerFunc func(ctx workerlibcontext.Context)

func Run(workerHandlerFunc WorkerHandlerFunc) {
	app, cleanup, err := di.InitializeApplication()
	if err != nil {
		cleanup()
		panic(err)
	}

	panic(app)
}
