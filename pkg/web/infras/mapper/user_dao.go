package mapper

import (
	"demo-go/pkg/web/domain/context"
	"demo-go/pkg/web/domain/model"
	"reflect"
)

type UserDaoIn interface {
	GetById(id int8) (*model.User, error)
}

func UserDao(ctx context.DemoContext) *userDao {
	of := reflect.TypeOf(ctx)
	of.Name()

	return &userDao{baseDao{context.NewDemoContext(ctx)}}
}

type userDao struct {
	baseDao
}

func (u userDao) GetById(id int64) (*model.User, error) {
	user := &model.User{}
	_, err := u.baseDao.Get(id, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
