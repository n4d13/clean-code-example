package config

import (
	"os"
	"sync"
	"strings"
	"strconv"
	"log"
	"time"
)

type ConfigurationManager struct {
	Host         string
	Port         int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolSize     int
	PoolTimeout  time.Duration
}

var instance *ConfigurationManager
var once sync.Once

func GetInstance() *ConfigurationManager {
	once.Do(func() {
		instance = &ConfigurationManager{
			Host:         getValueOrDefault("REDIS_HOST", ""),
			Port:         toInt(getValueOrFail("REDIS_PORT")),
			DialTimeout:  toDuration(getValueOrDefault("REDIS_DIAL_TIMEOUT", "2")),
			ReadTimeout:  toDuration(getValueOrDefault("REDIS_READ_TIMEOUT", "2")),
			WriteTimeout: toDuration(getValueOrDefault("REDIS_WRITE_TIMEOUT", "1")),
			PoolSize:     toInt(getValueOrDefault("REDIS_POOL_SIZE", "5")),
			PoolTimeout:  toDuration(getValueOrDefault("REDIS_POOL_TIMEOUT", "30")),
		}
	})
	return instance
}

func (c *ConfigurationManager) Address() string {
	return c.Host + ":" + strconv.Itoa(c.Port)
}

func getValueOrFail(name string) string {
	val := os.Getenv(name)
	if strings.TrimSpace(val) == "" {
		log.Fatal("Required value is missing: " + name)
	}
	return val
}

func getValueOrDefault(name string, defaultValue string) string {
	val := os.Getenv(name)
	if strings.TrimSpace(val) == "" {
		return defaultValue
	}
	return val
}

func toInt(val string) int {
	result, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal("Error converting to int. Value: " + val)
	}
	return result
}

func toDuration(val string) time.Duration {
	return time.Duration(toInt(val)) * time.Second
}
