package routers

import (
	v1 "gin-base/app/controller/v1"
	"github.com/gin-gonic/gin"
)

type MenuRouter struct{}

// RegisterRoutes 实现 Router 接口
func (r MenuRouter) RegisterRoutes(routerGroup *gin.RouterGroup) {
	var (
		controller v1.MenuController
	)

	// 列表
	routerGroup.GET("/menu", controller.List)

	// 创建
	routerGroup.POST("/menu", controller.Create)

	// 更新
	routerGroup.PUT("/menu/:id", controller.Update)

	// 删除
	routerGroup.DELETE("/menu/:id", controller.Delete)

	// 功能列表
	routerGroup.GET("/menu/:id/action", controller.ActionList)

	// 功能创建
	routerGroup.POST("/menu/:id/action", controller.ActionCreate)

	// 功能更新
	routerGroup.PUT("/menu/:id/action/:actionId", controller.ActionUpdate)

	// 功能删除
	routerGroup.DELETE("/menu/:id/action/:actionId", controller.ActionDelete)

	// 功能详情
	routerGroup.GET("/menu/:id/action/:actionId", controller.ActionDetail)
}
