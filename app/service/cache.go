package service

import (
	"gin-base/common"
	"gin-base/common/global"
	"time"
)

type CacheService struct {
	common.BaseService
}

// @function: SetCache
// @description: 设置缓存
// @param: key string, value interface{}, expire time.Duration
// @return: interface{}
func (s *CacheService) SetCache(key string, value interface{}, expire time.Duration) interface{} {
	global.Cache.SetCache(key, value, expire)

	return true
}

// @function: GetCache
// @description: 获取缓存
// @param: key string
// @return: interface{}
func (s *CacheService) GetCache(key string) (interface{}, bool) {
	res, ok := global.Cache.GetCache(key)
	if ok {
		return res, ok
	}

	return false, ok
}

// @function: DeleteCache
// @description: 删除缓存
// @param: key string
// @return: interface{}
func (s *CacheService) DeleteCache(key string) interface{} {
	global.Cache.DeleteCache(key)

	return true
}
