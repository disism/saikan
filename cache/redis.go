package cache

import (
	"github.com/disism/saikan/conf"
	"github.com/redis/go-redis/v9"
)

func NewRdbClient(db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.GetRedisAddr(),
		Password: conf.GetRedisPassword(),
		DB:       db,
	})
	return rdb
}
