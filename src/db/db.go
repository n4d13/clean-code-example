package db

import (
	"github.com/go-redis/redis"
	"github.com/n4d13/clean-code-example/src/config"
)

type RedisConnection struct {
	client *redis.Client
}

func MakeRedisConnection() (*RedisConnection, error) {

	client := redis.NewClient(&redis.Options{
		Addr:         config.GetInstance().Address(),
		DialTimeout:  config.GetInstance().DialTimeout,
		ReadTimeout:  config.GetInstance().ReadTimeout,
		WriteTimeout: config.GetInstance().WriteTimeout,
		PoolSize:     config.GetInstance().PoolSize,
		PoolTimeout:  config.GetInstance().PoolTimeout,
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
