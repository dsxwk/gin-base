package routers

import (
	controller "gin-base/app/controller/v1"
	"gin-base/app/middleware"
	"github.com/gin-gonic/gin"
)

var (
	jwtMiddleware = middleware.Jwt{}.JwtMiddleware()
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

	// 统一路由分组
	v1 := router.Group("api/v1")
	{
		// 不需要token验证
		// 用户登录
		v1.POST("/login", login_controller.Login)

		// 需要token验证
		v1.Use(jwtMiddleware)
		{
			// 用户列表
			v1.GET("/user", user_controller.List)

			// 用户详情
			v1.GET("/user/:id", user_controller.Detail)

			// 文章列表
			v1.GET("/article", article_controller.List)

			// 创建文章
			v1.POST("/article", article_controller.Create)

			// 更新文章
			v1.PUT("/article/:id", article_controller.Update)

			// 文章详情
			v1.GET("/article/:id", article_controller.Detail)

			// 删除文章
			v1.DELETE("/article/:id", article_controller.Delete)
		}
	}
}
