package routers

import (
	"gin-base/app/middleware"
	"github.com/gin-gonic/gin"
)

var (
	jwtMiddleware  = middleware.Jwt{}.Handle()
	I18nMiddleware = middleware.I18n{}.Handle()
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
		menu    MenuRouter
		role    RoleRouter
		dict    DictRouter
	)

	// 登录
	login.RegisterRoutes(v1.Group("", I18nMiddleware)) // new(LoginRouter).RegisterRoutes(v1)

	// 需要权限
	auth := v1.Group("", I18nMiddleware, jwtMiddleware)
	{
		// 用户
		user.RegisterRoutes(auth)
		// 文章
		article.RegisterRoutes(auth)
		// 缓存
		cache.RegisterRoutes(auth)
		// 菜单
		menu.RegisterRoutes(auth)
		// 角色
		role.RegisterRoutes(auth)
		// 字典
		dict.RegisterRoutes(auth)
	}
}
