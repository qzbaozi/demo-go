package service

import (
	"demo-go/pkg/web/domain/model"
	"demo-go/pkg/web/infras/gatewayimpl"
)

type userServiceImpl struct {
	baseService
}

func (u userServiceImpl) GetUserById(id int64) *model.User {
	return gatewayimpl.UserGateWay(u).Get(id)
}
