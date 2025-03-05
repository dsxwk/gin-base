package global

import (
	"gin-base/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 全局变量
var (
	DB     = config.InitDB()
	Log    = config.InitLogger()
	Config = config.InitConfig()
	Cache  = config.InitCache()
	Redis  = config.InitRedis()
)

// PageData 公共分页数据返回
type PageData struct {
	//  总条数
	Total int64 `json:"total"`
	// 当前页
	Page int `json:"page"`
	// 每页条数
	PageSize int `json:"pageSize"`
	// 数据列表
	List interface{} `json:"list"`
}

// ErrorCode 错误码和提示信息
type ErrorCode struct {
	Code    int64
	Message string
}

// Response 公共响应
type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// 错误码
var (
	Success      = ErrorCode{Code: 0, Message: "Success"}
	ArgsError    = ErrorCode{Code: 400, Message: "参数错误"}
	Unauthorized = ErrorCode{Code: 401, Message: "请求未授权"}
	SystemError  = ErrorCode{Code: 500, Message: "系统错误"}
)

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
