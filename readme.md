# 郑重提示

本文档由`大师兄`整理编写邮箱25076778@qq.com，严禁转载。
禁止将本项目用于含病毒、木马、色情、赌博、诈骗、违禁用品、假冒产品、虚假信息、数字货币、金融等违法违规业务

当前项目仅供个人学习测试，禁止一切线上商用行为，禁止一切违法使用！！！

## Gin框架介绍

Gin是一个用Go语言编写的Web框架。它具有简单、快速、高效等特点，被广泛应用于Go语言的Web开发中。

Gin框架的特性包括：

- 快速：Gin框架基于标准库`net/http`，使用`goroutine`和`channel`实现异步处理，提高性能。
- 简单：Gin框架提供了一系列的API和中间件，使得开发人员可以快速构建Web应用程序。
- 高效：Gin框架使用`sync.Pool`来缓存对象，减少内存分配和释放，提高性能。

Golang Gin 是一个轻量级且高效的 Golang Web 框架。它具有高性能、易用性和灵活性等特点，被广泛应用于各种 Web 应用程序的开发。

## 目录结构

```bash
├── app                                 # 应用程序代码
│   ├── controller                      # 控制期
│   ├── model                           # 模型
│   ├── service                         # 服务
│   ├── validate                        # 验证器
│   ├── middleware                      # 中间件
├── cli                                 # 命令行
├── common                              # 公共模块
│   ├── global                          # 全局变量
├── config                              # 配置文件
├── docs                                # 文档
├── database                            # 测试数据库文件
├── helper                              # 工具类
├── log                                 # 日志文件
├── resource                            # 静态资源
├── web                                 # Web服务
├── routers                             # 路由
├── vendor                              # 依赖包
```

## web目录结构
## 目录结构

```bash
├── app                                 # 应用
│   ├── modules                         # 模块
├── components                          # 组件
├── node_modules                        # npm包
├── public                              # 静态资源
├── docs                                # 文档
├── routers                             # 路由
├── src                                 # 核心js源码
├── views                               # 视图文件
```

## 使用
### 命令行生成模型 --path参数默认就行无需修改

```bash
go run ./cli/main.go --tableName=user --path=app/temp
go run ./cli/main.go --tableName=article --path=app/temp
go run ./cli/main.go --tableName=category --path=app/temp
go run ./cli/main.go --tableName=system_config --path=app/temp
```
## 登录接口

```bash
curl -X POST -H "Content-Type: application/json" -d '{"username":"admin","password":"123456"}' http://localhost:8080/api/v1/login
```

## jwt验证 header

```bash
curl -X GET -H "token: login-token -> eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6IkFkbWluIiwiaWF0IjoxNjc0NjYzMjM5LCJleHAiOjE2NzQ2NjYzMzl9.8W45GJQqV656
```

## 安装

使用 Golang 的包管理工具 `go get` 安装 Gin 框架：

```bash
go get -u github.com/gin-gonic/gin
```

## 基本用法

使用 Gin 框架创建一个简单的 Web 应用程序，包括路由和处理请求的函数。

```go
package main

import "github.com/gin-gonic/gin"

func main() {
    // 创建一个默认的 Gin 引擎
    r := gin.Default()

    // 定义一个路由和处理函数
    r.GET("/hello", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, World!",
        })
    })

    // 启动服务器并监听默认端口
    r.Run()
}
```

## 路由

Gin 框架支持多种类型的路由，包括 GET、POST、PUT、DELETE 等。可以使用 `r.GET()`、`r.POST()`、`r.PUT()`、`r.DELETE()` 等方法定义路由和处理函数。

```go
r.GET("/hello", func(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "Hello, World!",
    })
})

r.POST("/hello", func(c *gin.Context) {
    // 处理 POST 请求
})
```

## 中间件

Gin 框架支持使用中间件来处理请求的预处理和后处理逻辑。可以使用 `r.Use()` 方法定义中间件，并在路由处理函数之前或之后应用。

```go
r.Use(Logger())

func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 处理请求的预处理逻辑
    }
}

r.GET("/hello", func(c *gin.Context) {
    // 处理 GET 请求
})

r.Use(Recovery())

func Recovery() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 处理请求的后处理逻辑
    }
}
```

## 参数绑定

Gin 框架支持自动将请求中的参数绑定到处理函数的参数上。可以使用 `c.Bind()` 方法将请求中的参数绑定到结构体或 map 中。

```go
type User struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
}

r.POST("/user", func(c *gin.Context) {
    var user User
    if err := c.Bind(&user); err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }

    // 处理请求
})
```

## 文件上传

Gin 框架支持文件上传功能。可以使用 `c.SaveUploadedFile()` 方法将上传的文件保存到指定的路径。

```go
r.POST("/upload", func(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }

    // 保存文件到指定路径
    err = c.SaveUploadedFile(file, "path/to/save/file")
    if err != nil {
        c.JSON(500, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(200, gin.H{  
        "message": "File uploaded successfully",
    })
})
```
## 路由分组

Gin 框架支持路由分组功能，可以将多个路由分组到一个中间件中，方便管理和维护。

```go
v1 := r.Group("/v1")
{
    v1.GET("/hello", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, v1",
        })
    })

    v1.GET("/world", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "World, v1",
        })
    })
}

v2 := r.Group("/v2")
{
    v2.GET("/hello", func(c *gin.Context)
}
```