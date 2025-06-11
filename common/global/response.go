package global

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

// Response 公共响应
type Response struct {
	TraceID string      `json:"traceId"`
	Code    int64       `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// ApiResponse api统一返回
func ApiResponse(ctx *gin.Context, errorCode ErrorCode, optionalParams ...interface{}) {
	var response Response

	response.TraceID = ctx.GetString("traceId")

	response.Code = errorCode.Code

	// 默认消息
	response.Message = errorCode.Message

	// 默认数据
	response.Data = []string{}

	fmt.Printf("%v\n", optionalParams)
	// 解析可选参数,当第一个参数为字符串时默认为message否则为data
	if len(optionalParams) > 0 {
		if msg, ok := optionalParams[0].(string); ok {
			response.Message = msg
		} else {
			if optionalParams[0] == nil || isEmptySlice(optionalParams[0]) {
				response.Data = []string{}
			} else {
				response.Data = optionalParams[0]
			}
		}
	}

	if len(optionalParams) > 1 {
		if msg, ok := optionalParams[0].(string); ok {
			response.Message = msg
		} else {
			panic("参数错误,第一个参数message必须是字符串")
		}
		if optionalParams[1] == nil || isEmptySlice(optionalParams[1]) {
			response.Data = []string{}
		} else {
			response.Data = optionalParams[1]
		}
	}

	ctx.JSON(http.StatusOK, response)
	ctx.Abort()
}

func isEmptySlice(data interface{}) bool {
	v := reflect.ValueOf(data)
	return v.Kind() == reflect.Slice && v.Len() == 0
}
