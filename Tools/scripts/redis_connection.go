package scripts

import (
	"github.com/go-redis/redis"
)

func GetRedis(DB int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "47.107.112.81:6379",
		Password: "Wy20001976",
		DB:       DB, // use default DB
	})

	return rdb
}
