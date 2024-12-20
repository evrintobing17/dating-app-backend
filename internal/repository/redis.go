package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(addr, password string, db int) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisClient{Client: rdb}
}

func (r *RedisClient) Close() error {
	return r.Client.Close()
}

func (r *RedisClient) AddToSet(key string, value interface{}, expiration time.Duration) error {
	if err := r.Client.SAdd(ctx, key, value).Err(); err != nil {
		return err
	}
	return r.Client.Expire(ctx, key, expiration).Err()
}

func (r *RedisClient) IsMemberOfSet(key string, value interface{}) (bool, error) {
	return r.Client.SIsMember(ctx, key, value).Result()
}
