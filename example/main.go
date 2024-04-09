package main

import (
	"errors"
	containerlib "github.com/resource-aware-jds/container-lib"
	"github.com/resource-aware-jds/container-lib/model"
	"github.com/resource-aware-jds/container-lib/pkg/containerlibcontext"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	containerlib.Run(func(ctx containerlibcontext.Context, task model.Task) error {
		count := 0
		for {
			select {
			case <-ctx.Done():
				logrus.Warn("Context has been canceled")
				return nil
			default:
				// Simulate the log running job.
				time.Sleep(1 * time.Second)
			}

			if count == 20 {
				ctx.RecordResult([]byte("Success!!"))
				ctx.Success()
				return nil
			} else if count > 20 {
				ctx.RecordResult([]byte("Error"))
				return errors.New("example error")
			}
			count++
		}
	})
}
