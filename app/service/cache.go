package service

import (
	"gin-base/common"
	"gin-base/common/global"
	"gin-base/helper/utils"
	"time"
)

type CacheService struct {
	common.BaseService
}

// SetCache 设置缓存
// @param: key string, value interface{}, expire time.Duration
// @return: interface{}
func (s *CacheService) SetCache(key string, value interface{}, expire time.Duration) interface{} {
	global.Cache.SetCache(key, value, expire)

	return true
}

// GetCache 获取缓存
// @param: key string
// @return: interface{}
func (s *CacheService) GetCache(key string) (interface{}, bool) {
	res, ok := global.Cache.GetCache(key)
	if ok {
		return res, ok
	}

	return false, ok
}

// DeleteCache 删除缓存
// @param: key string
// @return: interface{}
func (s *CacheService) DeleteCache(key string) interface{} {
	global.Cache.DeleteCache(key)

	return true
}

// Send http请求测试
// @param: none
// @return: interface{}
func (s *CacheService) Send() (interface{}, error) {
	res, err := utils.HttpRequest("put", "http://127.0.0.1:8080/api/v1/article/1", map[string]string{
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDAwNTAwNzAsImlkIjoxfQ.rsULvvOvFb2YO8D5hOHY6eCU9NxfNbYrNmG1VZQx0aw",
	}, map[string]interface{}{
		"title": "title1",
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
