package common

import (
	"gin-base/common/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

// 获取当前登录用户id
func (this *BaseController) GetUserId(ctx *gin.Context) int64 {
	id, _ := ctx.Get("user.id")
	uid := id.(float64)
	return int64(uid)
}

// api统一返回
func (this *BaseController) ApiResponse(ctx *gin.Context, code int64, message string, data interface{}) {
	var response global.Response

	if code != 0 {
		response.Code = code
		response.Message = "Error"
	} else {
		response.Message = "Success"
	}

	if message != "" {
		response.Message = message
	}

	if data != nil {
		response.Data = data
	} else {
		response.Data = []string{}
	}

	ctx.JSON(http.StatusOK, response)
	ctx.Abort()
}
