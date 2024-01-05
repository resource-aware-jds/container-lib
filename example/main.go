package main

import (
	workerlib "github.com/resource-aware-jds/container-lib"
	"github.com/resource-aware-jds/container-lib/model"
	"github.com/resource-aware-jds/container-lib/pkg/containerlibcontext"
	"time"
)

func main() {
	workerlib.Run(func(ctx containerlibcontext.Context, task model.Task) error {
		// Simulate the log running job.
		time.Sleep(10 * time.Second)
		return nil
	})
}
