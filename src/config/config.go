package config

import (
	"os"
	"sync"
)

type ConfigurationManager struct {
	Host         string
	Port         string
	DialTimeout  string
	ReadTimeout  string
	WriteTimeout string
	PoolSize     string
	PoolTimeout  string
}

var instance *ConfigurationManager
var once sync.Once

func GetInstance() *ConfigurationManager {
	once.Do(func(){
		instance = &ConfigurationManager{
			Host:         os.Getenv("REDIS_HOST"),
			Port:         os.Getenv("REDIS_PORT"),
			DialTimeout:  os.Getenv("REDIS_DIAL_TIMEOUT"),
			ReadTimeout:  os.Getenv("REDIS_READ_TIMEOUT"),
			WriteTimeout: os.Getenv("REDIS_WRITE_TIMEOUT"),
			PoolSize:     os.Getenv("REDIS_POOL_SIZE"),
			PoolTimeout:  os.Getenv("REDIS_POOL_TIMEOUT"),
		}
	})
	return instance
}
