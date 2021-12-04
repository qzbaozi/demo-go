package gatewayimpl

import (
	"demo-go/pkg/web/domain/model"
	"demo-go/pkg/web/infras/mapper"
	"github.com/kataras/golog"
)

type userGatewayImpl struct {
	baseGateWay
}

func (u userGatewayImpl) Get(id int64) *model.User {
	get := new(model.User)
	has, err := mapper.UserDao(u).Get(id, get)
	if err != nil {
		golog.Error(err)
	}
	if has {
		return get
	}
	return nil
}

func (u userGatewayImpl) Del(id int64) bool {
	del := new(model.User)
	del.Id = id
	count, err := mapper.UserDao(u).Delete(del)
	if err != nil {
		golog.Error(err)
	}
	return count > 0
}
