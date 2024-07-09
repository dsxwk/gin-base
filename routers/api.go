package routers

import (
	controller "gin-base/app/controller/v1"
	"github.com/gin-gonic/gin"
)

// 加载路由
func LoadRouters(router *gin.Engine) {
	/*// 统一路由分组
	v1 := router.Group("api/v1")

	// 加载路由...
	// 登录
	LoginRoutes(v1)*/

	// 登录
	login_controller := controller.LoginController{}

	// 用户
	user_controller := controller.UserController{}

	// 文章
	article_controller := controller.ArticleController{}

	v1 := router.Group("api/v1")
	{
		v1.GET("/login", login_controller.Login)
		v1.GET("/user", user_controller.List)
		v1.GET("/article", article_controller.List)
		v1.POST("/article", article_controller.Create)
		v1.PUT("/article/:id", article_controller.Update)
		v1.GET("/article/:id", article_controller.Detail)
		v1.DELETE("/article/:id", article_controller.Delete)
	}
}
