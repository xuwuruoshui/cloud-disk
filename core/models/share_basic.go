package models

import "time"

type ShareBasic struct {
	Id int
	Identity string
	UserIdentity string
	RepositoryIdentity string
	UserRepositoryIdentity string
	ExpireTime int
	ClickNum int
	CreateAt time.Time `xorm:"created"`
	UpdateAt time.Time `xorm:"updated"`
	DeleteAt time.Time `xorm:"deleted"`
}

func (table ShareBasic)TableName()string{
	return "share_basic"
}
