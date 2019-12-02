package db

import (
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/kushao1267/Facade/facade/config"
	"github.com/mgutz/ansi"
)

var RedisDB *redis.Client

// See use: https://github.com/go-redis/redis/blob/master/example_test.go
func Init() {
	// 初始化redis
	RedisDB = NewRedis(config.AllConf.Redis)
}

func NewRedis(c config.Redis) *redis.Client {
	db := redis.NewClient(&redis.Options{
		Addr:        c.Addr,
		Password:    c.Password,
		DB:          c.DB, // use default DB
		PoolSize:    10,
		PoolTimeout: 30 * time.Second,
	})

	if err := db.Ping().Err(); err != nil {
		log.Println(ansi.Color("[初始化redis失败]:", "red"), err)
	}
	return db
}
