package context

import (
	"context"
	"github.com/gin-gonic/gin"
)

var (
	Ctx context.Context
)

// SetGinContext 设置 Gin 上下文
func SetGinContext(c *gin.Context, key string) {
	Ctx = context.WithValue(context.Background(), key, c)
}

// GetGinContext 获取 Gin 上下文
func GetGinContext(key string) *gin.Context {
	if Ctx != nil {
		if ginCtx, ok := Ctx.Value(key).(*gin.Context); ok {
			return ginCtx
		}
	}
	return nil
}
