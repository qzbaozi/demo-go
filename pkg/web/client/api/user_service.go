package service

import "demo-go/pkg/web/domain/model"

type UserService interface {
	GetUserById(id int64) *model.User
}
