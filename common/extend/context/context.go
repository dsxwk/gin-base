package context

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	ginContextStore = make(map[string]interface{})
	mu              sync.RWMutex
)

// SetContext 设置 Context
func SetContext(key string, value *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	ginContextStore[key] = value
}

// GetContext 获取 Context
func GetContext(key string) *gin.Context {
	mu.RLock()
	defer mu.RUnlock()
	if ctx, ok := ginContextStore[key]; ok {
		return ctx.(*gin.Context)
	}
	return nil
}

// ClearContext 清理 Context
func ClearContext(key string) {
	mu.Lock()
	defer mu.Unlock()
	delete(ginContextStore, key)
}
