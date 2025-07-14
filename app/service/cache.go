package service

import (
	"gin-base/common/base"
	"gin-base/common/global"
	"gin-base/helper"
	"time"
)

type CacheService struct {
	base.BaseService
}

// SetCache 设置缓存
// @param key string
// @param value interface{}
// @param expire time.Duration
// @return bool, error
func (s *CacheService) SetCache(key string, value interface{}, expire time.Duration) (bool, error) {
	err := global.Cache.SetCache(key, value, expire)
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetCache 获取缓存
// @param key string
// @return interface{}, bool
func (s *CacheService) GetCache(key string) (interface{}, bool) {
	res, ok := global.Cache.GetCache(key)
	if ok {
		return res, ok
	}

	return false, ok
}

// DeleteCache 删除缓存
// @param key string
// @return bool, error
func (s *CacheService) DeleteCache(key string) (bool, error) {
	err := global.Cache.DeleteCache(key)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Send http请求测试
// @return interface{}
func (s *CacheService) Send() (interface{}, error) {
	res, err := helper.HttpRequest("put", "http://127.0.0.1:8080/api/v1/article/1", map[string]string{
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDAwNTAwNzAsImlkIjoxfQ.rsULvvOvFb2YO8D5hOHY6eCU9NxfNbYrNmG1VZQx0aw",
	}, map[string]interface{}{
		"title": "title1",
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
