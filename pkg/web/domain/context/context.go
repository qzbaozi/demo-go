package context

import (
	"context"
	"demo-go/pkg/web/infras/db"
	"github.com/kataras/golog"
	"github.com/xormplus/xorm"
	"sync"
)

type DemoContext interface {
	GetDbOps() xorm.Interface

	GetCtx() context.Context
}

type baseDemoContext struct {
	context           context.Context
	xormSession       *xorm.Session
	xormSessionLocker sync.Mutex
}

func NewDemoContext(ctx DemoContext) DemoContext {
	if ctx == nil {
		ctx = new(baseDemoContext)
	}
	return ctx
}

func NewWebContext(ctx context.Context) DemoContext {
	return &baseDemoContext{
		context: ctx,
	}
}

func (b *baseDemoContext) GetDbOps() xorm.Interface {
	b.xormSessionLocker.Lock()
	defer b.xormSessionLocker.Unlock()
	if b.xormSession != nil {
		return b.xormSession
	}
	return db.GetEngine()
}

func (b *baseDemoContext) GetCtx() context.Context {
	if b.context == nil {
		golog.Error("base.context is nil")
		return context.TODO()
	}
	return b.context
}
