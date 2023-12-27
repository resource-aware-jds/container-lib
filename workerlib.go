package workerlib

import (
	"github.com/resource-aware-jds/worker-lib/di"
	"github.com/resource-aware-jds/worker-lib/model"
)

func Run(workerHandlerFunc model.WorkerHandlerFunc) {
	app, cleanup, err := di.InitializeApplication(workerHandlerFunc)
	if err != nil {
		cleanup()
		panic(err)
	}

	panic(app)
}
