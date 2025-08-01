package cache

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"time"
)

// RedisCache Redis缓存
type RedisCache struct {
	client  *redis.Client
	ctx     context.Context
	pubsubs map[string]*redis.PubSub
}

// NewRedisCache redis实例
func NewRedisCache(address, password string, db int) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	return &RedisCache{
		client:  client,
		ctx:     context.Background(),
		pubsubs: make(map[string]*redis.PubSub),
	}
}

// Lock 获取 Redis 锁
func (r *RedisCache) Lock(key string, value string, expire time.Duration) error {
	// 使用 SETNX 命令尝试设置锁
	result, err := r.client.SetNX(r.ctx, key, value, expire).Result()
	if err != nil {
		return fmt.Errorf("failed to acquire lock: %v", err)
	}

	if !result {
		// 如果返回 false,表示锁已存在
		return fmt.Errorf("lock already exists")
	}

	return nil
}

// UnLock 释放 Redis 锁
func (r *RedisCache) UnLock(key string, value string) error {
	script := `
	if redis.call("get", KEYS[1]) == ARGV[1] then
		return redis.call("del", KEYS[1])
	else
		return 0
	end`
	// 使用 EVAL 命令执行 Lua 脚本
	status, err := r.client.Eval(r.ctx, script, []string{key}, value).Int()
	if err != nil {
		return fmt.Errorf("failed to unlock: %v", err)
	}

	if status == 0 {
		return fmt.Errorf("unlock failed: lock not owned or already released")
	}

	return nil
}

// Publish 发布
func (r *RedisCache) Publish(channel string, message interface{}) error {
	err := r.client.Publish(r.ctx, channel, message).Err()
	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}
	return nil
}

// Subscribe 订阅
func (r *RedisCache) Subscribe(channel string, handler func(channel string, payload string)) error {
	pubsub := r.client.Subscribe(r.ctx, channel)

	// 等待订阅确认
	_, err := pubsub.Receive(r.ctx)
	if err != nil {
		return fmt.Errorf("failed to subscribe to channel %s: %v", channel, err)
	}

	// 保存 pubsub 对象
	r.pubsubs[channel] = pubsub

	// 消息处理协程
	go func() {
		ch := pubsub.Channel()
		for msg := range ch {
			handler(msg.Channel, msg.Payload)
		}
	}()

	return nil
}

// Unsubscribe 取消订阅
func (r *RedisCache) Unsubscribe(channel string) error {
	pubsub, ok := r.pubsubs[channel]
	if !ok {
		return fmt.Errorf("channel %s not found in subscriptions", channel)
	}

	err := pubsub.Unsubscribe(r.ctx, channel)
	if err != nil {
		return fmt.Errorf("failed to unsubscribe from channel %s: %v", channel, err)
	}

	// 关闭并删除记录
	err = pubsub.Close()
	if err != nil {
		return fmt.Errorf("failed to close pubsub for channel %s: %v", channel, err)
	}

	delete(r.pubsubs, channel)
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
