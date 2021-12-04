package web

import (
	"demo-go/pkg/web/domain/context"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	irisContext "github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
)

type BaseControl struct {
	Ctx iris.Context
	context.DemoContext
}

var Controllers = make(map[string]mvc.BaseController)

func init() {
	Controllers["user"] = new(UsersController)
}

func (b *BaseControl) BeginRequest(ctx irisContext.Context) {
	b.DemoContext = context.NewWebContext(ctx.Request().Context())
}

func (b *BaseControl) EndRequest(ctx irisContext.Context) {
	golog.Debugf("EndRequest: %v", ctx.Request().RequestURI)
}
