package main

import (
	workerlib "github.com/resource-aware-jds/worker-lib"
	"github.com/resource-aware-jds/worker-lib/pkg/workerlibcontext"
	"time"
)

func main() {
	workerlib.Run(func(ctx workerlibcontext.Context) error {
		// Simulate the log running job.
		time.Sleep(10 * time.Second)
		return nil
	})
}
