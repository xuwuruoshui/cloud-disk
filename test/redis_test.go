package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)
var ctx = context.Background()
func TestRedis(t *testing.T){
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.0.132:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "name", "123456", time.Second*5).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
