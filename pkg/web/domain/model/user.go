package model

import "time"

type User struct {
	Id         int64     `xorm:"INTEGER pk 'id'" json:"id"`
	Name       string    `xorm:"TEXT notnull 'name'" json:name`
	Age        int       `xorm:"INTEGER notnull 'age'" json:"age"`
	Sex        string    `xorm:"TEXT notnull 'sex'" json:"sex"`
	CreateTime time.Time `xorm:"NUMERIC notnull 'create_time'" json:"createTime"`
}

func (u User) TableName() string {
	return "user"
}
