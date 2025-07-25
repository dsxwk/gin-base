package global

// ErrorCode 错误码和提示信息
type ErrorCode struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`
}

// 错误码
var (
	Success      = ErrorCode{Code: 0, Message: "errcode.success"}
	Redirect     = ErrorCode{Code: 301, Message: "errcode.redirect"}
	ArgsError    = ErrorCode{Code: 400, Message: "errcode.argsError"}
	Unauthorized = ErrorCode{Code: 401, Message: "errcode.unauthorized"}
	NotFound     = ErrorCode{Code: 404, Message: "errcode.notFound"}
	SystemError  = ErrorCode{Code: 500, Message: "errcode.systemError"}
)
