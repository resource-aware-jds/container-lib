package taskrunner

import "github.com/resource-aware-jds/container-lib/pkg/runnerpool"

type service struct {
	runnerPool runnerpool.Pool
}

type Service interface {
}

func ProvideService(runnerPool runnerpool.Pool) Service {
	return &service{
		runnerPool: runnerPool,
	}
}
