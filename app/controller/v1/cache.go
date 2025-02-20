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

// SetCache
// @Description 设置缓存
// @Router /v1/cache [post]
func (s *CacheController) SetCache(c *gin.Context) {
	var (
		cacheService service.CacheService
		req          validate.CacheValidate
	)

	err := c.ShouldBind(&req)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	// 验证
	err = validate.GetCacheValidate(req, "setCache")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error(), nil)
		return
	}

	res := cacheService.SetCache(req.Key, req.Value, (req.Expire)*time.Second)

	s.ApiResponse(c, global.Success, "success", res)
}

// GetCache
// @Description 获取缓存
// @Router /v1/cache [get]
func (s *CacheController) GetCache(c *gin.Context) {
	var (
		cacheService service.CacheService
		req          validate.CacheValidate
	)

	err := c.ShouldBindQuery(&req)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	// 验证
	err = validate.GetCacheValidate(req, "getCache")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error(), nil)
		return
	}

	res, _ := cacheService.GetCache(req.Key)

	s.ApiResponse(c, global.Success, "success", res)
}

// DeleteCache
// @Description 删除缓存
// @Router /v1/cache [delete]
func (s *CacheController) DeleteCache(c *gin.Context) {
	var (
		cacheService service.CacheService
		req          validate.CacheValidate
	)

	err := c.ShouldBindQuery(&req)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	// 验证
	err = validate.GetCacheValidate(req, "deleteCache")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error(), nil)
		return
	}

	res := cacheService.DeleteCache(req.Key)

	s.ApiResponse(c, global.Success, "success", res)
}

// Send
// @Description http请求测试
// @Router /v1/send [post]
func (s *CacheController) Send(c *gin.Context) {
	var (
		cacheService service.CacheService
	)

	res, err := cacheService.Send()
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error(), nil)
		return
	}

	s.ApiResponse(c, global.Success, "success", res)
}
