package main

import (
	workerlib "github.com/resource-aware-jds/worker-lib"
	"time"
)

func main() {
	workerlib.Run(func(ctx containerlibcontext.Context) error {
		// Simulate the log running job.
		time.Sleep(10 * time.Second)
		return nil
	})
}
