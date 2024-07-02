package routers

import (
	v1 "gin-base/app/controller/v1"
	"github.com/gin-gonic/gin"
)

func LoginRoutes(router *gin.RouterGroup) {
	// 路由分组
	// group := router.Group("/login/v1")
	// loginController := new(controller.LoginController)
	loginController := v1.LoginController{}
	// 登录
	router.POST("/login", loginController.Login)
}
