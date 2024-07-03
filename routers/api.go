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

	v1 := router.Group("api/v1")
	{
		v1.GET("/login", login_controller.Login)
		v1.GET("/user", user_controller.List)
	}
}
