package tests

import (
	"gin-base/common/global"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// 获取缓存驱动
func TestCache_GetDriver(t *testing.T) {
	driver := global.Config.Cache.Driver
	assert.NotEmpty(t, driver)
}

// 设置缓存
func TestCache_SetCache(t *testing.T) {
	key := "test_key"
	value := "test_value"
	err := global.Cache.SetCache(key, value, time.Second*10)
	assert.NoError(t, err)
}

// 获取缓存
func TestCache_GetCache(t *testing.T) {
	key := "test_key"
	value := "test_value"
	result, found := global.Cache.GetCache(key)
	assert.True(t, found)
	assert.Equal(t, value, result)
}

// 获取缓存过期时间
func TestCache_GetCacheExpire(t *testing.T) {
	key := "test_key"
	value := "test_value"
	e := global.Cache.SetCache(key, value, 5*time.Second)
	assert.NoError(t, e)
	val, expireAt, found, err := global.Cache.GetCacheExpire(key)

	assert.NoError(t, err)
	assert.True(t, found)
	assert.Equal(t, value, val)
	assert.Greater(t, expireAt.Unix(), time.Now().Unix())
}

// 删除缓存
func TestCache_DeleteCache(t *testing.T) {
	key := "test_key"
	err := global.Cache.DeleteCache(key)
	assert.NoError(t, err)
}
