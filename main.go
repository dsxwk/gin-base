package main

import (
	"fmt"
	"gin-base/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go get -u
//go:generate go mod tidy
//go:generate go mod download
//go:generate go mod vendor

func main() {
	//gin.SetMode(gin.ReleaseMode) // 生产模式
	gin.SetMode(gin.DebugMode) // debug模式
	//gin.SetMode(gin.TestMode) // 测试模式

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "pong",
			"data":    []string{},
		})
	})

	// 静态文件
	router.StaticFS("/resource", http.Dir("./resource"))
	// 设置 HTTP 请求处理文件上传时可以使用的最大内存为 90MB
	router.MaxMultipartMemory = 90 << 20

	// 加载路由
	routers.LoadRouters(router)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("启动服务失败，错误信息为：", err)
	}
}
