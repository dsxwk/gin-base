package global

import (
	"gin-base/config"
)

// 全局变量
var (
	DB     = config.InitDB()
	Log    = config.InitLogger()
	Config = config.InitConfig()
	Cache  = config.InitCache()
	Redis  = config.InitRedis()
)

// Response 公共响应
type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

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

// 错误码
const (
	// Success 成功
	Success = 0
	// ArgsError 参数错误
	ArgsError = 400
	// Unauthorized 请求未授权
	Unauthorized = 401
	// SystemError 系统错误
	SystemError = 500
)
