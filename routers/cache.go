package routers

import (
	"gin-base/app/controller/v1"
	"github.com/gin-gonic/gin"
)

// CacheRouter 缓存路由
type CacheRouter struct{}

// RegisterRoutes 实现 Router 接口
func (r CacheRouter) RegisterRoutes(routerGroup *gin.RouterGroup) {
	var (
		controller v1.CacheController
	)

	// 设置缓存
	routerGroup.GET("/cache", controller.SetCache)
	// 获取缓存
	routerGroup.GET("/cache/:key", controller.GetCache)
	// 删除缓存
	routerGroup.DELETE("/cache", controller.DeleteCache)
	// http请求测试
	routerGroup.GET("/send", controller.Send)
}
