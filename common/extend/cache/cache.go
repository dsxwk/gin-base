package cache

import (
	"errors"
	"fmt"
	"github.com/dgraph-io/badger/v4"
	"github.com/go-redis/redis/v8"
	"github.com/patrickmn/go-cache"
	"golang.org/x/net/context"
	"log"
	"time"
)

// InterfaceCache Cache 接口
type InterfaceCache interface {
	SetCache(key string, value interface{}, expire time.Duration) error
	GetCache(key string) (interface{}, bool)
	DeleteCache(key string) error
	GetCacheExpire(key string) (interface{}, time.Time, bool, error)
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

func (m *MemoryCache) SetCache(key string, value interface{}, expire time.Duration) error {
	if expire == 0 {
		expire = cache.NoExpiration
	}
	m.cache.Set(key, value, expire)

	return nil
}

func (m *MemoryCache) GetCache(key string) (interface{}, bool) {
	return m.cache.Get(key)
}

func (m *MemoryCache) DeleteCache(key string) error {
	m.cache.Delete(key)

	return nil
}

func (m *MemoryCache) GetCacheExpire(key string) (interface{}, time.Time, bool, error) {
	value, exp, found := m.cache.GetWithExpiration(key)
	if !found {
		return nil, time.Time{}, false, errors.New("cache key not found")
	}

	return value, exp, true, nil
}

// RedisCache Redis 缓存实现
type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

// NewRedisCache redis实例
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

// Lock 获取 Redis 锁
func (r *RedisCache) Lock(lockKey string, expire time.Duration) (bool, error) {
	// 使用 SETNX 命令尝试设置锁
	result, err := r.client.SetNX(r.ctx, lockKey, 1, expire).Result()
	if err != nil {
		return false, fmt.Errorf("failed to acquire lock: %v", err)
	}
	if !result {
		// 如果返回 false，表示锁已存在
		return false, nil
	}

	// 锁获取成功
	return true, nil
}

// UnLock 释放 Redis 锁
func (r *RedisCache) UnLock(lockKey string) error {
	// 获取锁
	_, err := r.client.Get(r.ctx, lockKey).Result()
	if errors.Is(err, redis.Nil) {
		return errors.New("lock does not exist")
	} else if err != nil {
		return fmt.Errorf("failed to get lock: %v", err)
	}

	// 删除该锁
	err = r.client.Del(r.ctx, lockKey).Err()
	if err != nil {
		return fmt.Errorf("failed to release lock: %v", err)
	}

	return nil
}

func (r *RedisCache) SetCache(key string, value interface{}, expire time.Duration) error {
	err := r.client.Set(r.ctx, key, value, expire).Err()
	if err != nil {
		return fmt.Errorf("error setting Redis cache: %v", err)
	}

	return nil
}

func (r *RedisCache) GetCache(key string) (interface{}, bool) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		return nil, false
	}

	return val, true
}

func (r *RedisCache) DeleteCache(key string) error {
	err := r.client.Del(r.ctx, key).Err()
	if err != nil {
		return fmt.Errorf("error deleting Redis cache: %v", err)
	}

	return nil
}

func (r *RedisCache) GetCacheExpire(key string) (interface{}, time.Time, bool, error) {
	// Redis不支持使用相同的API获取到期时间,因此必须使用TTL
	ttl, err := r.client.TTL(r.ctx, key).Result()
	if err != nil {
		return nil, time.Time{}, false, fmt.Errorf("error getting TTL for key %v: %v", key, err)
	}

	val, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		return nil, time.Time{}, false, fmt.Errorf("error getting value for key %v: %v", key, err)
	}

	expireTime := time.Now().Add(ttl)

	return val, expireTime, true, nil
}

// DiskCache badger 磁盘缓存
type DiskCache struct {
	db *badger.DB
}

func NewDiskCache(dir string) (*DiskCache, error) {
	opts := badger.DefaultOptions(dir)
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(fmt.Sprintf("init disk cache failed: %s", err.Error()))
		return nil, err
	}
	return &DiskCache{db: db}, nil
}

func (b *DiskCache) SetCache(key string, value interface{}, expire time.Duration) error {
	valBytes, ok := value.([]byte)
	if !ok {
		valBytes = []byte(fmt.Sprintf("%v", value))
	}
	err := b.db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(key), valBytes)
		if expire > 0 {
			e = e.WithTTL(expire)
		}
		return txn.SetEntry(e)
	})
	return err
}

func (b *DiskCache) GetCache(key string) (interface{}, bool) {
	var val []byte
	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		val, err = item.ValueCopy(nil)
		return err
	})
	if err != nil {
		return nil, false
	}
	return val, true
}

func (b *DiskCache) DeleteCache(key string) error {
	return b.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}

// GetCacheExpire Badger 不直接支持获取剩余 TTL，只能判断是否存在
func (b *DiskCache) GetCacheExpire(key string) (interface{}, time.Time, bool, error) {
	var val []byte
	var expireTime time.Time
	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		val, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}
		ttl := item.ExpiresAt()
		if ttl > 0 {
			expireTime = time.Unix(int64(ttl), 0)
		}
		return nil
	})
	if err != nil {
		return nil, time.Time{}, false, errors.New("cache key not found")
	}
	return val, expireTime, true, nil
}
