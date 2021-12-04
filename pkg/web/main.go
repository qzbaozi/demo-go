package main

import (
	"demo-go/pkg/db/entity"
	"demo-go/pkg/web/adapter/web"
	"demo-go/pkg/web/domain/model"
	"demo-go/pkg/web/infras/db"
	"demo-go/pkg/web/infras/mapper"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	dbInit()
	app := iris.New()
	// You got full debug messages, useful when using MVC and you want to make
	// sure that your code is aligned with the Iris' MVC Architecture.
	app.Logger().SetLevel("debug")

	for k, v := range web.Controllers {
		myMvc := mvc.New(app.Party(k))
		myMvc.Handle(v)
	}

	app.Listen(":8080", iris.WithOptimizations)
}

func dbInit() {
	connect := db.GetEngine()
	user := entity.User{}
	exist, err := connect.IsTableExist(user)
	if err != nil {
		println(err)
	} else {
		if !exist {
			connect.CreateTables(user)
		}
	}

	userDao := mapper.UserDao(nil)

	u := model.User{Id: 3, Name: "xxx"}
	connect.Insert(u)
	get, err := userDao.GetById(u.Id)
	if get.Id == 3 {
		println("true")
	}
}
