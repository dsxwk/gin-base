package global

import (
	"gin-base/common/extend/cache"
	"gin-base/config"
	"gorm.io/gorm"
)

// 全局变量
var (
	FormatDate string               // 全局日期格式
	Config     *config.Config       // 配置
	Log        *config.Logger       // 日志
	DB         *gorm.DB             // 数据库
	Cache      cache.InterfaceCache // 缓存
	Redis      *cache.RedisCache    // Redis
	Event      *config.Events       // 事件
)

func init() {
	FormatDate = "2006-01-02 15:04:05"
	Config = config.InitConfig()
	Log = config.InitLogger(Config)
	DB = config.InitMysql(Config, Log)
	Cache = config.InitCache(Config)
	Redis = config.InitRedis(Config)
	Event = config.InitEvent()
}
