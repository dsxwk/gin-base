package helper

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/patrickmn/go-cache"
	"golang.org/x/net/context"
	"time"
)

// CacheInterface Cache 接口,所有缓存类型都实现这个接口
type CacheInterface interface {
	SetCache(key string, value interface{}, expire time.Duration)
	GetCache(key string) (interface{}, bool)
	DeleteCache(key string)
	GetCacheExpire(key string) (interface{}, time.Time, bool)
}

// MemoryCache 内存缓存实现
type MemoryCache struct {
	cache *cache.Cache
}

func NewMemoryCache(defaultExpiration, cleanupInterval time.Duration) *MemoryCache {
	return &MemoryCache{
		cache: cache.New(defaultExpiration, cleanupInterval),
	}
}

func (m *MemoryCache) SetCache(key string, value interface{}, expire time.Duration) {
	if expire == 0 {
		expire = cache.NoExpiration
	}
	m.cache.Set(key, value, expire)
}

func (m *MemoryCache) GetCache(key string) (interface{}, bool) {
	return m.cache.Get(key)
}

func (m *MemoryCache) DeleteCache(key string) {
	m.cache.Delete(key)
}

func (m *MemoryCache) GetCacheExpire(key string) (interface{}, time.Time, bool) {
	return m.cache.GetWithExpiration(key)
}

// RedisCache Redis 缓存实现
type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisCache(address, password string, db int) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	return &RedisCache{
		client: client,
		ctx:    context.Background(),
	}
}

func (r *RedisCache) SetCache(key string, value interface{}, expire time.Duration) {
	err := r.client.Set(r.ctx, key, value, expire).Err()
	if err != nil {
		fmt.Println("Error setting Redis cache:", err)
	}
}

func (r *RedisCache) GetCache(key string) (interface{}, bool) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		return nil, false
	}

	return val, true
}

func (r *RedisCache) DeleteCache(key string) {
	err := r.client.Del(r.ctx, key).Err()
	if err != nil {
		fmt.Println("Error deleting Redis cache:", err)
	}
}

func (r *RedisCache) GetCacheExpire(key string) (interface{}, time.Time, bool) {
	// Redis不支持使用相同的API获取到期时间,因此我们必须使用TTL
	ttl, err := r.client.TTL(r.ctx, key).Result()
	if err != nil {
		return nil, time.Time{}, false
	}

	val, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		return nil, time.Time{}, false
	}

	expireTime := time.Now().Add(ttl)

	return val, expireTime, true
}
