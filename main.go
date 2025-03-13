package main

import (
	"fmt"
	_ "gin-base/app/controller/v1"
	"gin-base/app/middleware"
	"gin-base/common/global"
	"gin-base/helper"
	"gin-base/routers"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go get -u
//go:generate go mod tidy
//go:generate go mod download
//go:generate go mod vendor

func main() {
	// 运行环境模式 debug模式, test测试模式, release生产模式, 默认是debug,根据当前配置文件读取
	gin.SetMode(global.Config.Env.Mode)

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "pong",
			"data":    []string{},
		})
	})

	// Swagger 路由
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 静态文件
	router.StaticFS("/resource", http.Dir("./resource"))
	// 设置 HTTP 请求处理文件上传时可以使用的最大内存为 90MB
	router.MaxMultipartMemory = 90 << 20

	// 设置跨域
	router.Use(Cors())

	// 全局日志中间件
	router.Use(middleware.LoggerMiddleware())

	// 注册所有事件
	global.Event.RegisterAllEvent(onEventReceived)

	// 加载路由
	routers.LoadRouters(router)

	err := router.Run(`:` + global.Config.Env.Port)
	if err != nil {
		fmt.Println("启动服务失败，错误信息为：", err)
	}
}

// onEventReceived 接收事件
func onEventReceived(event helper.Event, timestamp time.Time) {
	// todo 处理事件
	fmt.Printf("Event received at %s: name: %s, data: %v\n", timestamp.Format(time.RFC3339), event.Name, event.Data)
}

// Cors 跨域请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", global.Config.Cors.AllowOrigin)
		c.Header("Access-Control-Allow-Headers", global.Config.Cors.AllowHeaders)
		c.Header("Access-Control-Expose-Headers", global.Config.Cors.ExposeHeaders)
		c.Header("Access-Control-Allow-Methods", global.Config.Cors.AllowMethods)
		c.Header("Access-Control-Allow-Credentials", global.Config.Cors.AllowCredentials)

		// 放行所有OPTIONS方法
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}
