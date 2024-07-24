package data

import (
	"os"
	"strconv"
	"sync"

	"github.com/go-redis/redis"
)

var (
	redisOnce sync.Once
	rdb       *redis.Client
)

func Redis() *redis.Client {
	redisOnce.Do(func() {
		db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
		if err != nil {
			panic(err)
		}
		rdb = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       db,
		})
		if err := rdb.Ping().Err(); err != nil {
			panic(err)
		}
	})
	return rdb
}
