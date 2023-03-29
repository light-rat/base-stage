package redis

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis"
)

var Rdb *redis.Client

func Init() {
	rdb := redis.NewClient(&redis.Options{
		// 需要修改成你的配置，本地无需修改
		Addr:     "175.178.131.222:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("连接成功")
	// 成功连接将其赋值给全局变量
	Rdb = rdb
}

func RedisSet(key, value string) error {
	result := Rdb.Set(key, value, 0)
	if result.Err() != nil {
		return errors.New("redis set fail")
	} else {
		return nil
	}
}
