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
	// 缓存
	cache_controller := controller.CacheController{}

	// 统一路由分组
	v1 := router.Group("api/v1")
	{
		// 不需要token验证
		// 用户登录
		v1.POST("/login", login_controller.Login)

		// 需要token验证
		// 用户
		user := v1.Group("user", jwtMiddleware)
		{
			// 列表
			user.GET("", user_controller.List)
			// 创建用户
			user.POST("", user_controller.Create)
			// 更新用户
			user.PUT("/:id", user_controller.Update)
			// 用户详情
			user.GET("/:id", user_controller.Detail)
			// 删除用户
			user.DELETE("/:id", user_controller.Delete)
		}

		// 文章
		article := v1.Use(jwtMiddleware)
		{
			// 列表
			article.GET("/article", article_controller.List)
			// 创建
			article.POST("/article", article_controller.Create)
			// 更新
			article.PUT("/article/:id", article_controller.Update)
			// 详情
			article.GET("/article/:id", article_controller.Detail)
			// 删除
			article.DELETE("/article/:id", article_controller.Delete)
		}

		// 缓存
		cache := v1.Use(jwtMiddleware)
		{
			// 设置缓存
			cache.POST("/cache", cache_controller.SetCache)
			// 获取缓存
			cache.GET("/cache", cache_controller.GetCache)
			// 删除缓存
			cache.DELETE("/cache", cache_controller.DeleteCache)
		}
	}
}
