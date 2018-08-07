package db

import (
	"github.com/go-redis/redis"
	"time"
	"github.com/n4d13/clean-code-example/src/config"
	"strconv"
)

type RedisConnection struct {
	client *redis.Client
}

func MakeRedisConnection() (*RedisConnection, error) {

	address := config.GetInstance().Host+":"+config.GetInstance().Port

	dialTimeout, err := strconv.Atoi(config.GetInstance().DialTimeout)
	if err != nil{
		return nil, err
	}

	readTimeout, err := strconv.Atoi(config.GetInstance().ReadTimeout)
	if err != nil{
		return nil, err
	}

	writeTimeout, err := strconv.Atoi(config.GetInstance().WriteTimeout)
	if err != nil{
		return nil, err
	}

	poolSize, err := strconv.Atoi(config.GetInstance().PoolSize)
	if err != nil{
		return nil, err
	}

	poolTimeout, err := strconv.Atoi(config.GetInstance().PoolTimeout)
	if err != nil{
		return nil, err
	}

	client := redis.NewClient(&redis.Options{
		Addr:         address,
		DialTimeout:  time.Duration(dialTimeout) * time.Second,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
		PoolSize:     poolSize,
		PoolTimeout:  time.Duration(poolTimeout) * time.Second,
	})
	_, connErr := client.Ping().Result()

	if connErr != nil{
		return nil, connErr
	}

	return &RedisConnection{
		client: client,
	}, nil
}

func (r *RedisConnection) Count(keys ...string) int64{
	 return r.client.Exists(keys...).Val()
}