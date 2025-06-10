package config

import (
	"gin-base/common/extend/cache"
	"sync"
)

type Redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

var (
	redisOnce  sync.Once
	redisCache *cache.RedisCache
)

// InitRedis 初始化redis
func InitRedis(config *Config) *cache.RedisCache {
	redisOnce.Do(func() {
		redisCache = cache.NewRedisCache(
			config.Cache.Redis.Address,
			config.Cache.Redis.Password,
			config.Cache.Redis.DB,
		)
	})

	return redisCache
}
