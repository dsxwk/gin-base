package v1

import (
	"gin-base/app/service"
	"gin-base/app/validate"
	"gin-base/common"
	"gin-base/common/global"
	"github.com/gin-gonic/gin"
	"time"
)

type CacheController struct {
	common.BaseController
}

// SetCache 设置缓存
// @Router /api/v1/cache [post]
func (s *CacheController) SetCache(c *gin.Context) {
	var (
		cacheService  service.CacheService
		cacheValidate validate.Cache
	)

	err := c.ShouldBind(&cacheValidate)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.Cache{}.GetValidate(cacheValidate, "setCache")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	res, err := cacheService.SetCache(cacheValidate.Key, cacheValidate.Value, (cacheValidate.Expire)*time.Second)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, res)
}

// GetCache 获取缓存
// @Router /api/v1/cache [get]
func (s *CacheController) GetCache(c *gin.Context) {
	var (
		cacheService  service.CacheService
		cacheValidate validate.Cache
	)

	err := c.ShouldBindQuery(&cacheValidate)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.Cache{}.GetValidate(cacheValidate, "getCache")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	res, _ := cacheService.GetCache(cacheValidate.Key)

	s.ApiResponse(c, global.Success, res)
}

// DeleteCache 删除缓存
// @Router /api/v1/cache [delete]
func (s *CacheController) DeleteCache(c *gin.Context) {
	var (
		cacheService  service.CacheService
		cacheValidate validate.Cache
	)

	err := c.ShouldBindQuery(&cacheValidate)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.Cache{}.GetValidate(cacheValidate, "deleteCache")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	res, err := cacheService.DeleteCache(cacheValidate.Key)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, res)
}

// Send http请求测试
// @Router /api/v1/send [post]
func (s *CacheController) Send(c *gin.Context) {
	var (
		cacheService service.CacheService
	)

	res, err := cacheService.Send()
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, res)
}
