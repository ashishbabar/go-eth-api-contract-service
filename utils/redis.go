package utils

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type IRedisUtil interface {
	GetVal(ctx context.Context, key string) string
}
type redisUtil struct {
	client *redis.Client
}

func NewRedisUtil() IRedisUtil {
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("RedisURL"),
		Username: viper.GetString("RedisUsername"),
		Password: viper.GetString("RedisPassword"), // no password set
		DB:       0,                                // use default DB
	})

	return &redisUtil{client: rdb}
}

func (rc *redisUtil) GetVal(ctx context.Context, key string) string {
	return rc.client.Incr(ctx, key).String()
}
