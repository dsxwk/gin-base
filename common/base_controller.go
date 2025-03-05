package common

import (
	"gin-base/common/global"
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

// GetUserId 获取当前登录用户id
func (s *BaseController) GetUserId(ctx *gin.Context) int64 {
	id, _ := ctx.Get("user.id")
	uid := id.(float64)
	return int64(uid)
}

// ApiResponse api统一返回
func (s *BaseController) ApiResponse(ctx *gin.Context, errorCode global.ErrorCode, optionalParams ...interface{}) {
	global.ApiResponse(ctx, errorCode, optionalParams...)
}
