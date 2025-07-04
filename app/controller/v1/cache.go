package v1

import (
	"gin-base/app/service"
	"gin-base/app/validate"
	"gin-base/common/base"
	"gin-base/common/global"
	"github.com/gin-gonic/gin"
	"time"
)

type CacheController struct {
	base.BaseController
}

// SetCache 设置缓存
// @Tags 缓存管理
// @Summary 设置缓存
// @Description 设置缓存
// @Accept json
// @Produce json
// @Param token header string true "认证Token"
// @Param data body object true "创建参数" SchemaExample({"key":"缓存键","value":"缓存值","expire":60})
// @Router /api/v1/cache [post]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
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
// @Tags 缓存管理
// @Summary 获取缓存
// @Description 获取缓存
// @Param token header string true "认证Token"
// @Param key query string true "缓存键"
// @Router /api/v1/cache [get]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
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
// @Tags 缓存管理
// @Summary 删除缓存
// @Description 删除缓存
// @Param token header string true "认证Token"
// @Param key query string true "缓存键"
// @Router /api/v1/cache [delete]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
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

// Send http请求日志记录测试
// @Tags 缓存管理
// @Summary http请求日志记录测试
// @Description http请求日志记录测试
// @Param token header string true "认证Token"
// @Router /api/v1/send [post]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
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
