package routers

import (
	"gin-base/app/controller/v1"
	"github.com/gin-gonic/gin"
)

// ArticleRouter 文章路由
type ArticleRouter struct{}

// RegisterRoutes 实现 Router 接口
func (r ArticleRouter) RegisterRoutes(routerGroup *gin.RouterGroup) {
	var (
		controller = v1.ArticleController{}
	)

	// 列表
	routerGroup.GET("/user", controller.List)
	// 创建
	routerGroup.POST("/user", controller.Create)
	// 更新
	routerGroup.PUT("/user/:id", controller.Update)
	// 删除
	routerGroup.DELETE("/user/:id", controller.Delete)
	// 详情
	routerGroup.GET("/user/:id", controller.Detail)
}
