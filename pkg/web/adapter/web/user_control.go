package web

import (
	"demo-go/pkg/web/app/service"
	"demo-go/pkg/web/domain/model"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12/mvc"
)

type UsersController struct {
	BaseControl
}

//BeforeActivation func
func (u *UsersController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/get", "GetUserById")
}

func (u *UsersController) GetUserById() *model.User {
	id, err := u.Ctx.URLParamInt64("id")
	if err != nil {
		golog.Error(err)
		return nil
	}

	return service.UserService(u).GetUserById(id)
}
