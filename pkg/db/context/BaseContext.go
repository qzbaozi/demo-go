package context

import (
	"context"
	"demo-go/pkg/db/pool"
	"github.com/xormplus/xorm"
	"sync"
)

type BaseContext interface {
	GetDbOps() xorm.Interface

	GetContext() context.Context
}

type baseContext struct {
	context context.Context

	xormSession       *xorm.Session
	xormSessionLocker sync.Mutex
}

func NewBaseContext(ctx BaseContext) BaseContext {
	if ctx == nil {
		ctx = new(baseContext)
	}
	return ctx
}

func (b baseContext) GetDbOps() xorm.Interface {
	b.xormSessionLocker.Lock()
	defer b.xormSessionLocker.Unlock()

	if b.xormSession != nil {
		return b.xormSession
	}

	return pool.Connect(b.GetContext())
}

func (b baseContext) GetContext() context.Context {
	if b.context == nil {
		b.context = context.TODO()
	}
	return b.context
}
