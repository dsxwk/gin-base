package global

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 公共响应
type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// ApiResponse api统一返回
func ApiResponse(ctx *gin.Context, errorCode ErrorCode, optionalParams ...interface{}) {
	var response Response

	response.Code = errorCode.Code

	// 默认消息
	response.Message = errorCode.Message

	// 默认数据
	response.Data = []string{}

	// 解析可选参数,当第一个参数为字符串时默认为message否则为data
	if len(optionalParams) > 0 {
		if msg, ok := optionalParams[0].(string); ok {
			response.Message = msg
		} else {
			response.Data = optionalParams[0]
		}
	}

	if len(optionalParams) > 1 {
		if msg, ok := optionalParams[0].(string); ok {
			response.Message = msg
		} else {
			panic("参数错误,第一个参数message必须是字符串")
		}
		response.Data = optionalParams[1]
	}

	ctx.JSON(http.StatusOK, response)
	ctx.Abort()
}
