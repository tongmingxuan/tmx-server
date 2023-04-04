// Package tmxServer /*
package tmxServer

import (
	"github.com/go-redis/redis"
	"sync"
	"unsafe"
)

type RedisPool struct {
	DbConfig map[string]map[string]string
	Pool     sync.Map
}

func (redisPool *RedisPool) Key() string {
	return "redis"
}

func (redisPool *RedisPool) Handle() *BaseFrame {
	frameConfig := (*Config)(unsafe.Pointer(frameContainer.Get(new(Config).Key())))

	redisConfig, ok := frameConfig.BaseConfigMap["redis"]

	if ok != true {
		panic("redis配置异常")
	}

	redisPool.DbConfig = redisConfig

	for connectionName, connectionConfig := range redisConfig {
		redisPool.CreateConnection(connectionName, connectionConfig)
	}

	return (*BaseFrame)(unsafe.Pointer(redisPool))
}

func (redisPool *RedisPool) CreateConnection(connectionName string, connectionInfo map[string]string) {
	db := StringToInt(connectionInfo["redis_db"], "初始redisDB异常")
	poolSize := StringToInt(connectionInfo["redis_max_connection"], "初始redis_max_connection异常")
	minIdleCon := StringToInt(connectionInfo["redis_min_open_connection"], "初始redis_min_open_connection异常")

	client := redis.NewClient(&redis.Options{
		Addr:         connectionInfo["redis_host"] + ":" + connectionInfo["redis_port"],
		Password:     connectionInfo["redis_auth"],
		DB:           db,
		PoolSize:     poolSize,
		MinIdleConns: minIdleCon,
	})

	res := client.Ping()

	_, err := res.Result()

	if err != nil {
		panic(connectionName + ":链接redis异常:" + err.Error())
	}

	redisPool.Pool.Store(connectionName, client)
}
