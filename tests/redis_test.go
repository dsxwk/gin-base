package tests

import (
	"gin-base/common/global"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// 设置缓存
func TestRedis_SetCache(t *testing.T) {
	key := "test_key"
	value := "test_value"
	err := global.Redis.SetCache(key, value, time.Second*10)
	assert.NoError(t, err)
}

// 获取缓存
func TestRedis_GetCache(t *testing.T) {
	key := "test_key"
	value := "test_value"
	result, found := global.Redis.GetCache(key)
	assert.True(t, found)
	assert.Equal(t, value, result)
}

// 获取缓存过期时间
func TestRedis_GetCacheExpire(t *testing.T) {
	key := "redis_key"
	value := "redis_value"
	e := global.Redis.SetCache(key, value, 5*time.Second)
	assert.NoError(t, e)
	val, expireAt, found, err := global.Redis.GetCacheExpire(key)

	assert.NoError(t, err)
	assert.True(t, found)
	assert.Equal(t, value, val)
	assert.Greater(t, expireAt.Unix(), time.Now().Unix())
}

// 删除缓存
func TestRedis_DeleteCache(t *testing.T) {
	key := "test_key"
	err := global.Redis.DeleteCache(key)
	assert.NoError(t, err)
}

// 测试加锁与解锁
func TestRedis_LockUnlock(t *testing.T) {
	key := "test:lock:key"
	value := uuid.NewString()

	// 正常加锁
	err := global.Redis.Lock(key, value, 5*time.Second)
	assert.NoError(t, err)

	// 再次加锁应失败
	err2 := global.Redis.Lock(key, "other-value", 5*time.Second)
	assert.Error(t, err2)

	// 正常解锁
	err3 := global.Redis.UnLock(key, value)
	assert.NoError(t, err3)

	// 解锁已释放的锁，应报错
	err4 := global.Redis.UnLock(key, value)
	assert.Error(t, err4)
}

func TestRedis_PublishSubscribe(t *testing.T) {
	channel := "test:pubsub"
	message := "hello world"
	received := make(chan string, 1)

	// 订阅频道
	err := global.Redis.Subscribe(channel, func(ch string, payload string) {
		if ch == channel {
			received <- payload
		}
	})
	assert.NoError(t, err)

	// 发布消息
	err = global.Redis.Publish(channel, message)
	assert.NoError(t, err)

	select {
	case msg := <-received:
		assert.Equal(t, message, msg)
	case <-time.After(2 * time.Second):
		t.Error("did not receive message in time")
	}

	// 取消订阅
	err = global.Redis.Unsubscribe(channel)
	assert.NoError(t, err)
}
