package workerlib

import (
	"github.com/resource-aware-jds/worker-lib/di"
	"github.com/resource-aware-jds/worker-lib/model"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func Run(workerHandlerFunc model.WorkerHandlerFunc) {
	app, cleanup, err := di.InitializeApplication(workerHandlerFunc)
	if err != nil {
		cleanup()
		panic(err)
	}

	app.GRPCServer.Serve()

	// Gracefully Shutdown
	// Make channel listen for signals from OS
	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	<-gracefulStop

	logrus.Info("Gracefully shutting down, cleaning up.")
	cleanup()
	logrus.Info("Clean up success. Good Bye")
}
