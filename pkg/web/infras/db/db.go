package db

import (
	"demo-go/pkg/web/infras/config"
	"github.com/kataras/golog"
	"github.com/xormplus/xorm"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var engine *xorm.Engine
var once sync.Once

func GetEngine() *xorm.Engine {
	once.Do(func() {
		engine = getEngine()
	})
	return engine
}

func getEngine() *xorm.Engine {
	db := config.NewDB()
	engine, err := xorm.NewEngine("sqlite3", "E:\\db\\test.db")

	if err != nil {
		golog.Fatal(err)
		return nil
	}
	if db.ShowSql {
		engine.ShowSQL(db.ShowSql)
		engine.SetLogLevel(db.LogLevel)
	}

	engine.SetMaxIdleConns(db.MaxIdl)
	engine.SetMaxOpenConns(db.MaxOpen)
	engine.SetConnMaxLifetime(time.Minute * 10)

	for i := 120; ; i-- {
		if i <= 0 {
			golog.Fatal("connect db fatal")
		}

		err := engine.Ping()
		if err != nil {
			golog.Errorf("db ping err [%d] err:%s", i, err)
		} else {
			break
		}
		time.Sleep(time.Second)
	}

	engine.TZLocation = time.Local

	return engine
}
