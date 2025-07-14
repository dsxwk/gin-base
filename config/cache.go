package config

import (
	"fmt"
	"gin-base/common/extend/cache"
	"log"
	"time"
)

// Cache 缓存
type Cache struct {
	Type  string `yaml:"type"`
	Redis Redis
}

// InitCache 初始化缓存
func InitCache(config *Config) cache.InterfaceCache {
	var (
		c cache.InterfaceCache
	)

	// 没有配置缓存类型,默认使用内存缓存
	if config.Cache.Type == "memory" || config.Cache.Type == "" {
		c = cache.NewMemoryCache(5*time.Minute, 10*time.Minute)
	} else if config.Cache.Type == "disk" {
		c, _ = cache.NewDiskCache("storage/cache")
	} else if config.Cache.Type == "redis" {
		// 使用 Redis 缓存
		c = InitRedis(config)
	} else {
		log.Fatalf("Unsupported cache type: %s", config.Cache.Type)
	}

	fmt.Println("Cache initialized with type:", config.Cache.Type)

	return c
}
