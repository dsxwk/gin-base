package middleware

import (
	"gin-base/common/extend/context"
	"github.com/gin-gonic/gin"
)

// LoggerMiddleware 全局日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		context.SetGinContext(c, "ginLogContext")

		c.Next()
	}
}
