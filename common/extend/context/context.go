package context

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	ginContextStore = make(map[string]interface{})
	mu              sync.RWMutex
)

// SetGinContext 设置 gin.Context
func SetGinContext(key string, value interface{}) {
	mu.Lock()
	defer mu.Unlock()
	ginContextStore[key] = value
}

// GetGinContext 获取 gin.Context
func GetGinContext(key string) *gin.Context {
	mu.RLock()
	defer mu.RUnlock()
	if ctx, ok := ginContextStore[key]; ok {
		return ctx.(*gin.Context)
	}
	return nil
}

// ClearGinContext 清理 gin.Context
func ClearGinContext(key string) {
	mu.Lock()
	defer mu.Unlock()
	delete(ginContextStore, key)
}
