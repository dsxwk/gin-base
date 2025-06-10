package config

import "gin-base/common/extend/cache"

type Redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// InitRedis 初始化redis
func InitRedis(config *Config) *cache.RedisCache {
	return cache.NewRedisCache(config.Cache.Redis.Address, config.Cache.Redis.Password, config.Cache.Redis.DB)
}
