package danta

import "context"

type Context struct {
	Context    context.Context
	CancelFunc context.CancelFunc
}

func NewContext() (Context, func()) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	return Context{ctx, cancelFunc}, cancelFunc
}

func NewContextSide(self Context) context.Context {
	return self.Context
}

func NewCancelFuncSide(self Context) context.CancelFunc {
	return self.CancelFunc
}
