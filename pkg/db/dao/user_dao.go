package dao

import (
	"demo-go/pkg/db/context"
	"demo-go/pkg/db/entity"
)

type UserDao interface {
	Get(id int64) (*entity.User, error)
}

func NewUserDao() *userDao {
	return &userDao{baseDao{context.NewBaseContext(nil)}}
}

type userDao struct {
	baseDao
}

func (u userDao) Get(id int64) (*entity.User, error) {
	user := &entity.User{}
	_, err := u.baseDao.Get(id, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
