package workerlibcontext

import (
	stdcontext "context"
	"github.com/resource-aware-jds/worker-lib/model"
	"time"
)

type context struct {
	ctx            stdcontext.Context
	isSuccess      bool
	cancelFunction func()
	results        [][]byte
	task           model.Task
	taskAttributes model.TaskAttributes
}

type Context interface {
	stdcontext.Context
	Success()
	Cancel()
	RecordResult(result []byte)
	GetTask() model.Task
}

func ProvideContext(ctx stdcontext.Context, task model.Task) Context {
	newCtx, cancelFunc := stdcontext.WithCancel(ctx)

	return &context{
		ctx:            newCtx,
		cancelFunction: cancelFunc,
		task:           task,
		results:        make([][]byte, 0),
	}
}

func (c *context) Deadline() (deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}

func (c *context) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *context) Err() error {
	return c.ctx.Err()
}

func (c *context) Value(key any) any {
	return c.ctx.Value(key)
}

func (c *context) Success() {
	c.isSuccess = true
}

func (c *context) Cancel() {
	c.cancelFunction()
}

func (c *context) RecordResult(result []byte) {
	c.results = append(c.results, result)
}

func (c *context) GetTask() model.Task {
	return c.task
}
