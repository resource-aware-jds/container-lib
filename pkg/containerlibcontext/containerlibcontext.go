package containerlibcontext

import (
	stdcontext "context"
	"time"
)

type context struct {
	ctx            stdcontext.Context
	isSuccess      bool
	cancelFunction func()
	results        [][]byte
}

type Context interface {
	stdcontext.Context
	Success()
	Cancel()
	RecordResult(result []byte)
	GetResults() [][]byte
	GetSuccessResult() bool
}

func ProvideContext(ctx stdcontext.Context) Context {
	newCtx, cancelFunc := stdcontext.WithCancel(ctx)

	return &context{
		isSuccess:      false,
		ctx:            newCtx,
		cancelFunction: cancelFunc,
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

func (c *context) GetResults() [][]byte {
	return c.results
}

func (c *context) GetSuccessResult() bool {
	return c.isSuccess
}
