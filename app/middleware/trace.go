package middleware

import (
	"gin-base/common/base"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Trace struct {
	base.BaseMiddleware
}

// TraceMiddleware trace
func (s Trace) TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := uuid.New().String()
		c.Set("traceId", traceId)
		// 响应头返回
		c.Writer.Header().Set("X-Trace-Id", traceId)
		c.Next()
	}
}
