package db

import (
	"github.com/go-redis/redis"
	"time"
)

type RedisConnection struct {
	client *redis.Client
}

func MakeRedisConnection(address string, dialTimeout time.Duration, readTimeout time.Duration,
	writeTimeout time.Duration, poolSize int, poolTimeout time.Duration) (*RedisConnection, error) {

	client := redis.NewClient(&redis.Options{
		Addr:         address,
		DialTimeout:  dialTimeout,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		PoolSize:     poolSize,
		PoolTimeout:  poolTimeout,
	})
	_, connErr := client.Ping().Result()

	if connErr != nil {
		return nil, connErr
	}

	return &RedisConnection{
		client: client,
	}, nil
}

func (r *RedisConnection) Count(keys ...string) int64 {
	return r.client.Exists(keys...).Val()
}
