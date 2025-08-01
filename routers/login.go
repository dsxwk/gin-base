package routers

import (
	"gin-base/app/controller/v1"
	"github.com/gin-gonic/gin"
)

// LoginRouter 登录路由
type LoginRouter struct{}

// RegisterRoutes 实现 Router 接口
func (r LoginRouter) RegisterRoutes(routerGroup *gin.RouterGroup) {
	var (
		login v1.LoginController
	)

	// 登录
	routerGroup.POST("/login", login.Login)
	// 刷新token
	routerGroup.POST("/refresh-token", login.RefreshToken)
	// 获取验证码
	routerGroup.GET("/captcha", login.GetCaptcha)
	// 校验验证码
	routerGroup.POST("/captcha", login.CheckCaptcha)
}
