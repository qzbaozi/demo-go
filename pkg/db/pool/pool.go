package pool

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xormplus/xorm"
)

var engine *xorm.Engine

func Connect(context.Context) *xorm.Engine {
	var err error
	if engine == nil {
		engine, err = xorm.NewEngine("sqlite3", "E:\\db\\test.db")
	}

	if err != nil {
		println(err)
	}

	err = engine.Ping()

	return engine
}
