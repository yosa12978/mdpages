package data

import (
	"sync"

	"github.com/go-redis/redis"
	"github.com/yosa12978/mdpages/config"
)

var (
	redisOnce sync.Once
	rdb       *redis.Client
)

func Redis() *redis.Client {
	redisOnce.Do(func() {
		cfg := config.Get()
		rdb = redis.NewClient(&redis.Options{
			Addr:     cfg.Redis.Addr,
			Password: cfg.Redis.Password,
			DB:       cfg.Redis.Db,
		})
		if err := rdb.Ping().Err(); err != nil {
			panic(err)
		}
	})
	return rdb
}
