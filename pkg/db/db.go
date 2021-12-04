package main

import (
	"demo-go/pkg/db/dao"
	"demo-go/pkg/db/entity"
	"demo-go/pkg/db/pool"
)

func main() {
	connect := pool.Connect(nil)
	user := entity.User{}

	exist, err := connect.IsTableExist(user)
	if err != nil {
		println(err)
	} else {
		if !exist {
			connect.CreateTables(user)
		}
	}

	userDao := dao.NewUserDao()

	u := entity.User{Id: 3, Name: "xxx"}
	userDao.GetDbOps().Insert(u)
	get, err := userDao.Get(3)
	if get.Id == 3 {
		println("true")
	}
}
