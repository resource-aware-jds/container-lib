package taskrunner

import (
	"context"
)

type service struct {
	runnerPool taskrunner.Pool
}

type Service interface {
	Loop(ctx context.Context)
}

func ProvideService(runnerPool taskrunner.Pool) Service {
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

func (s *service) OnEvent(e taskrunner.PoolEvent) {
	//TODO implement me
	panic("implement me")
}

func (s *service) loop() {

}
