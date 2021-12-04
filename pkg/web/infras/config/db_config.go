package config

import (
	"github.com/xormplus/xorm/log"
)

type dbConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Name string `yaml:"name"` //root账户
	Pwd  string `yaml:"pwd"`  //root密码

	//不纳入yaml的配置
	ShowSql  bool `yaml:"showSql"`
	LogLevel log.LogLevel
	MaxIdl   int
	MaxOpen  int
}

func NewDB() dbConfig {
	return dbConfig{
		LogLevel: log.LOG_INFO,
		MaxIdl:   10,
		MaxOpen:  100,
	}
}
