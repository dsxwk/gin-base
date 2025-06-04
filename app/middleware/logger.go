package middleware

import (
	"bytes"
	"gin-base/common"
	"gin-base/common/extend/context"
	"gin-base/common/global"
	"gin-base/config"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Logger struct {
	common.BaseMiddleware
}

// LoggerMiddleware 全局日志中间件
func (s Logger) LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			params string
		)
		if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut || c.Request.Method == http.MethodPatch {
			bodyBytes, err := c.GetRawData()
			if err == nil {
				params = string(bodyBytes)
				// 恢复 body，供后续 ShouldBind 使用
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		} else if c.Request.Method == http.MethodGet || c.Request.Method == http.MethodDelete {
			params = c.Request.URL.RawQuery
		}
		// 保存 params 到 gin.Context
		c.Set("logParams", params)

		context.SetGinContext("GinLogger", c)

		c.Next()

		// 是否记录请求日志
		if global.Config.Log.Access {
			logFields := config.GetLogFields(nil)
			global.Log.Info("Access Log", logFields...)
		}

		context.ClearGinContext("GinLogger")
	}
}
