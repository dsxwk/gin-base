package main

import (
	"fmt"
	_ "gin-base/app/controller/v1"
	"gin-base/app/middleware"
	"gin-base/common/extend/i18n"
	"gin-base/common/global"
	"gin-base/config"
	_ "gin-base/docs"
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

// @title Gin-base Swagger API
// @version 2.0
// @description Gin-base API 文档
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email 25076778@qq.com
// @host 127.0.0.1:8080
func main() {
	// 运行环境模式 debug模式, test测试模式, release生产模式, 默认是debug,根据当前配置文件读取
	gin.SetMode(global.Config.Service.Mode)

	// 初始化 i18n（使用 embed）
	i18n.Init(true)

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
	if global.Config.Cors.Enabled {
		router.Use(middleware.Cors{}.Handle())
	}

	// 全局日志中间件
	router.Use(middleware.Logger{}.Handle())

	// 订阅事件
	global.Event.Subscribe(onReceived)

	// redis 订阅
	_ = global.Redis.Subscribe("userLogin", func(channel, payload string) {
		fmt.Printf("收到 %s 消息：%s\n", channel, payload)
	})

	// 加载路由
	routers.LoadRouters(router)

	fmt.Println(`启动服务: 0.0.0.0:` + global.Config.Service.Port)
	_ = router.Run(`:` + global.Config.Service.Port)
}

// onReceived 接收事件
func onReceived(event config.Event, timestamp time.Time) {
	// todo 处理事件
	fmt.Printf("Event received at %s: name: %s, data: %v\n", timestamp.Format(time.RFC3339), event.Name, event.Data)
	if event.Name == "sendHttp" {
		config.SetHttpLog(event.Data)
	}
}
