package routers

import (
	v1 "gin-base/app/controller/v1"
	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

// RegisterRoutes 实现 Router 接口
func (r RoleRouter) RegisterRoutes(routerGroup *gin.RouterGroup) {
	var (
		controller v1.RoleController
	)

	// 列表
	routerGroup.GET("/role", controller.List)
	// 创建
	routerGroup.POST("/role", controller.Create)
	// 更新
	routerGroup.PUT("/role/:id", controller.Update)
	// 删除
	routerGroup.DELETE("/role/:id", controller.Delete)
	// 详情
	routerGroup.GET("/role/:id", controller.Detail)
}
