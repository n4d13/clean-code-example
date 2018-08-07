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

	//Obtenemos el address de configuración
	address := config.GetInstance().Host+":"+config.GetInstance().Port

	//Obtenemos el Dial timeout
	intDtmt, err := strconv.Atoi(config.GetInstance().DialTimeout)
	if err != nil{
		return nil, err
	}

	//Obtenemos el Read timeout
	intRtmt, err := strconv.Atoi(config.GetInstance().ReadTimeout)
	if err != nil{
		return nil, err
	}

	//Obtenemos el Write timeout
	intWtmt, err := strconv.Atoi(config.GetInstance().WriteTimeout)
	if err != nil{
		return nil, err
	}

	//Obtenemos el tamaño del pool
	intPSize, err := strconv.Atoi(config.GetInstance().PoolSize)
	if err != nil{
		return nil, err
	}

	//Obtenemos el pool timeout
	intPtmt, err := strconv.Atoi(config.GetInstance().PoolTimeout)
	if err != nil{
		return nil, err
	}

	client := redis.NewClient(&redis.Options{
		Addr:         address,
		DialTimeout:  time.Duration(intDtmt) * time.Second,
		ReadTimeout:  time.Duration(intRtmt) * time.Second,
		WriteTimeout: time.Duration(intWtmt) * time.Second,
		PoolSize:     intPSize,
		PoolTimeout:  time.Duration(intPtmt) * time.Second,
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