package middleware

import (
	"gin-base/common/base"
	"gin-base/common/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Cors struct {
	base.BaseMiddleware
}

// Handle 跨域请求
func (s Cors) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", global.Config.Cors.AllowOrigin)
		c.Header("Access-Control-Allow-Headers", global.Config.Cors.AllowHeaders)
		c.Header("Access-Control-Expose-Headers", global.Config.Cors.ExposeHeaders)
		c.Header("Access-Control-Allow-Methods", global.Config.Cors.AllowMethods)
		c.Header("Access-Control-Allow-Credentials", global.Config.Cors.AllowCredentials)

		// 放行所有OPTIONS方法
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}
