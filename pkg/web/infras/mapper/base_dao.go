package mapper

import (
	"demo-go/pkg/web/domain/context"
	"github.com/pkg/errors"
)

type baseDao struct {
	context.DemoContext
}

func (b *baseDao) Get(id int64, obj interface{}) (has bool, err error) {
	get, err := b.GetDbOps().Where("id=?", id).Get(obj)
	if err != nil {
		return false, err
	}
	if !get {
		return false, errors.Errorf("Get %T,id [%d] not fount", obj, id)
	}
	return true, nil
}

func (b *baseDao) Delete(obj interface{}) (count int64, err error) {
	c, err := b.GetDbOps().Delete(obj)
	if err != nil {
		return 0, err
	}
	if c != 1 {
		return 0, errors.Errorf("delete %T,count [%d] not matched", obj, c)
	}
	return 1, nil
}

func (b *baseDao) Save(obj interface{}) error {
	_, err := b.GetDbOps().InsertOne(obj)
	return err
}
