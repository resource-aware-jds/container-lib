package taskrunner

import (
	"context"
	"github.com/resource-aware-jds/container-lib/pkg/runnerpool"
)

type service struct {
	runnerPool runnerpool.Pool
}

type Service interface {
	Loop(ctx context.Context)
}

func ProvideService(runnerPool runnerpool.Pool) Service {
	result := service{
		runnerPool: runnerPool,
	}

	runnerPool.Subscribe(&result)

	return &result
}

func (s *service) Loop(ctx context.Context) {
	go func(innerCtx context.Context) {
		for {
			select {
			case <-ctx.Done():
			// Graceful shutdown
			default:
				s.loop()
			}
		}
	}(ctx)
}

func (s *service) OnEvent(e runnerpool.PoolEvent) {
	//TODO implement me
	panic("implement me")
}

func (s *service) loop() {

}
