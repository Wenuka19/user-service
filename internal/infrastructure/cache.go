package infrastructure

import (
	"github.com/Wenuka19/user-service/internal/config"
	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	cfg := config.AppConfig

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPwd,
		DB:       0,
	})

	return rdb
}
