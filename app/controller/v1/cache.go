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
		cacheService service.CacheService
		req          validate.CacheValidate
	)

	err := c.ShouldBind(&req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.GetCacheValidate(req, "setCache")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	res, err := cacheService.SetCache(req.Key, req.Value, (req.Expire)*time.Second)
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
		cacheService service.CacheService
		req          validate.CacheValidate
	)

	err := c.ShouldBindQuery(&req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.GetCacheValidate(req, "getCache")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	res, _ := cacheService.GetCache(req.Key)

	s.ApiResponse(c, global.Success, res)
}

// DeleteCache 删除缓存
// @Router /api/v1/cache [delete]
func (s *CacheController) DeleteCache(c *gin.Context) {
	var (
		cacheService service.CacheService
		req          validate.CacheValidate
	)

	err := c.ShouldBindQuery(&req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.GetCacheValidate(req, "deleteCache")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	res, err := cacheService.DeleteCache(req.Key)
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
