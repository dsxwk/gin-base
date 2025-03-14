package global

import "gin-base/config"

// 全局变量
var (
	FormatDate = "2006-01-02 15:04:05" // 全局日期格式
	DB         = config.InitMysql()    // 数据库
	Log        = config.InitLogger()   // 日志
	Config     = config.InitConfig()   // 配置
	Cache      = config.InitCache()    // 缓存
	Redis      = config.InitRedis()    // Redis
	Event      = config.InitEvent()    // 事件
)
