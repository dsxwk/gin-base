package config

import (
	"fmt"
	"gin-base/common/extend/cache"
	"log"
	"time"
)

// Cache 缓存
type Cache struct {
	Driver string `yaml:"driver"`
	Redis  Redis
}

// InitCache 初始化缓存
func InitCache(config *Config) cache.InterfaceCache {
	var (
		c cache.InterfaceCache
	)

	// 没有配置缓存类型,默认使用内存缓存
	if config.Cache.Driver == "memory" || config.Cache.Driver == "" {
		c = cache.NewMemoryCache(5*time.Minute, 10*time.Minute)
	} else if config.Cache.Driver == "disk" {
		c, _ = cache.NewDiskCache("storage/cache")
	} else if config.Cache.Driver == "redis" {
		// 使用 Redis 缓存
		c = InitRedis(config)
	} else {
		log.Fatalf("Unsupported cache driver: %s", config.Cache.Driver)
	}

	fmt.Println("Cache initialized with driver:", config.Cache.Driver)

	return c
}
