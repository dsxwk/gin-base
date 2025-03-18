package global

// ErrorCode 错误码和提示信息
type ErrorCode struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`
}

// 错误码
var (
	Success      = ErrorCode{Code: 0, Message: "Success"}
	Redirect     = ErrorCode{Code: 301, Message: "数据已被重定向"}
	ArgsError    = ErrorCode{Code: 400, Message: "参数错误"}
	Unauthorized = ErrorCode{Code: 401, Message: "请求未授权"}
	NotFound     = ErrorCode{Code: 404, Message: "资源未找到"}
	SystemError  = ErrorCode{Code: 500, Message: "系统错误"}
)
