package routers

import (
	"gin-base/app/middleware"
	"github.com/gin-gonic/gin"
)

var (
	jwtMiddleware = middleware.Jwt{}.JwtMiddleware()
)

// Router 路由接口
type Router interface {
	RegisterRoutes(router *gin.RouterGroup)
}

// LoadRouters 加载路由
func LoadRouters(router *gin.Engine) {
	var (
		// 统一路由分组
		v1      = router.Group("api/v1")
		login   LoginRouter
		user    UserRouter
		article ArticleRouter
		cache   CacheRouter
	)

	// 登录
	login.RegisterRoutes(v1) // new(LoginRouter).RegisterRoutes(v1)

	// 需要权限
	auth := v1.Group("", jwtMiddleware)
	{
		// 用户
		user.RegisterRoutes(auth)
		// 文章
		article.RegisterRoutes(auth)
		// 缓存
		cache.RegisterRoutes(auth)
	}
}
