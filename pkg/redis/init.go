package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Redis struct {
	Rdb *redis.Client
}

func NewRedis(config *Config) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
		DB:       config.DB,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &Redis{Rdb: client}, nil
}

func (c *Redis) Set(key string, value any, ttl time.Duration) error {
	return c.Rdb.Set(ctx, key, value, ttl).Err()
}

func (c *Redis) Get(key string) (any, error) {
	return c.Rdb.Get(ctx, key).Result()
}
