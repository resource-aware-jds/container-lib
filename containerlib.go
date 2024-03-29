package containerlib

import (
	"github.com/resource-aware-jds/container-lib/di"
	"github.com/resource-aware-jds/container-lib/facade"
	"github.com/sirupsen/logrus"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

func Run(workerHandlerFunc facade.ContainerHandlerFunction) {
	app, cleanup, err := di.InitializeApplication(workerHandlerFunc)
	if err != nil {
		cleanup()
		panic(err)
	}

	if app.Config.Debug {
		// Start the pprof debug server
		go func() {
			err := http.ListenAndServe(app.Config.ProfilingToolsListeningURL, nil)
			if err != nil {
				logrus.Warn(err)
			}
		}()
	}

	app.GRPCServer.Serve()
	app.TaskRunnerSvc.Run()

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
