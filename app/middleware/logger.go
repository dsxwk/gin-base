package middleware

import (
	"gin-base/common/extend/context"
	"gin-base/common/global"
	"gin-base/config"
	"github.com/gin-gonic/gin"
)

// LoggerMiddleware 全局日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		context.SetGinContext("GinLogger", c)
		// 是否记录请求日志
		if global.Config.Log.Access {
			logFields := config.GetLogFields(nil)
			global.Log.Info("Access Log", logFields...)
		}

		c.Next()

		context.ClearGinContext("GinLogger")
	}
}
