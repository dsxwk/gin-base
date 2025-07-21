package cache

import "time"

// InterfaceCache 缓存接口
type InterfaceCache interface {
	SetCache(key string, value interface{}, expire time.Duration) error
	GetCache(key string) (interface{}, bool)
	DeleteCache(key string) error
	GetCacheExpire(key string) (interface{}, time.Time, bool, error)
}
