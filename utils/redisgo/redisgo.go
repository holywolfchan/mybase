package redisgo

import (
	"github.com/holywolfchan/yuncang/utils/logs"

	"github.com/go-redis/redis"
)

var (
	RedisEngine *redis.Client
)

func init() {
	var err error
	var pong string
	RedisEngine = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "63284200c",
		DB:       0,
	})
	pong, err = RedisEngine.Ping().Result()
	if pong == "PONG" {
		logs.Infof("redis server connected")
	} else {
		logs.Infof("redis connect failed:%v", err)
	}

}
