package routers

import (
	v1 "gin-base/app/controller/v1"
	"github.com/gin-gonic/gin"
)

type DictRouter struct{}

// RegisterRoutes 实现 Router 接口
func (r DictRouter) RegisterRoutes(routerGroup *gin.RouterGroup) {
	var (
		controller v1.DictController
	)

	// 列表
	routerGroup.GET("/dict", controller.List)

	// 创建
	routerGroup.POST("/dict", controller.Create)

	// 更新
	routerGroup.PUT("/dict/:id", controller.Update)

	// 详情
	routerGroup.GET("/dict/:id", controller.Detail)

	// 删除
	routerGroup.DELETE("/dict/:id", controller.Delete)
}
