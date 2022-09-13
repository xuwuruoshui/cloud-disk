package models

import (
	"cloud-disk/core/internal/config"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)


func Init(c config.Config) *xorm.Engine{
	var err error
	engine, err := xorm.NewEngine("mysql", c.Mysql.Datasource)
	engine.ShowSQL(true)
	if err!=nil{
		fmt.Println("xorm connect error: ",err)
		return nil
	}
	return engine
}


func InitRedis(c config.Config) *redis.Client{
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}