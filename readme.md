# 郑重提示

本文档由`大师兄`整理编写邮箱25076778@qq.com，严禁转载。
禁止将本项目用于含病毒、木马、色情、赌博、诈骗、违禁用品、假冒产品、虚假信息、数字货币、金融等违法违规业务

当前项目仅供个人学习测试，禁止一切线上商用行为，禁止一切违法使用！！！

## 项目地址
- Github: https://github.com/dsxwk/gin-base.git
- Gitee: https://gitee.com/wangke/gin-base.git

## Gin框架介绍

Gin是一个用Go语言编写的Web框架。它具有简单、快速、高效等特点，被广泛应用于Go语言的Web开发中。

Gin框架的特性包括：

- 快速：Gin框架基于标准库`net/http`，使用`goroutine`和`channel`实现异步处理，提高性能。
- 简单：Gin框架提供了一系列的API和中间件，使得开发人员可以快速构建Web应用程序。
- 高效：Gin框架使用`sync.Pool`来缓存对象，减少内存分配和释放，提高性能。

Golang Gin 是一个轻量级且高效的 Golang Web 框架。它具有高性能、易用性和灵活性等特点，被广泛应用于各种 Web 应用程序的开发。

## Gin-Base项目介绍
- 命令行生成模型
- 验证器以及自定义验证场景
- Jwt鉴权
- ...

## 后端使用技术

- Gin
- Gorm
- Jwt
- Mysql
- Validator

## 前端使用技术

- Vue3
- Vite
- Element-Plus
- Element-Plus-Table

## go-base目录结构

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
├── config                              # 配置
├── enums                               # 枚举
├── node_modules                        # npm包
├── public                              # 静态资源
├── docs                                # 文档
├── routers                             # 路由
├── src                                 # 核心js源码
├── utils                               # 工具助手类
├── views                               # 视图文件
```

## 使用
### 运行前后端

```bash
# 后端运行 端口8080 接口地址:127.0.0.1:8080/api/v1/...
go run main.go
```

```bash
# 前端运行 端口3000 访问地址:127.0.0.1:3000
npm run dev
```
### 命令行生成模型 --path参数默认就行无需修改

```bash
go run ./cli/main.go --tableName=user --path=app/temp
go run ./cli/main.go --tableName=article --path=app/temp
go run ./cli/main.go --tableName=category --path=app/temp
go run ./cli/main.go --tableName=system_config --path=app/temp
```
### 生成模型结构示例

```go
package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

const TableNameArticle = "article"

// Article mapped from table <article>
type Article struct {
	ID         int64     `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true;comment:ID" json:"id"`  // ID
	UID        int64     `gorm:"column:uid;type:int(11);not null;comment:用户id" json:"uid"`                            // 用户id
	User       *User     `json:"user" gorm:"foreignkey:uid;references:id"`                                            // 关联用户
	Title      string    `gorm:"column:title;type:varchar(50);not null;comment:标题" json:"title"`                      // 标题
	Content    string    `gorm:"column:content;type:varchar(255);not null;comment:内容" json:"content"`                 // 内容
	CategoryID int64     `gorm:"column:category_id;type:int(11);not null;comment:分类id" json:"category_id"`            // 分类id
	DataSource int64     `gorm:"column:data_source;type:int(11);not null;comment:数据来源 1=文章库 2=自建" json:"data_source"` // 数据来源 1=文章库 2=自建
	IsPublish  int64     `gorm:"column:is_publish;type:int(11);not null;comment:是否发布 1=已发布 2=未发布" json:"is_publish"`  // 是否发布 1=已发布 2=未发布
	Category   *Category `json:"category" gorm:"foreignkey:category_id;references:id"`                                // 关联分类
	Tag        *string   `gorm:"column:tag;type:json;comment:标签" json:"tag"`                                          // 标签
	CreatedAt  *string   `gorm:"column:created_at;type:datetime;comment:创建时间" json:"created_at"`                      // 创建时间
	UpdatedAt  *string   `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`                      // 更新时间
	DeletedAt  *string   `gorm:"column:deleted_at;type:datetime;comment:删除时间" json:"deleted_at"`                      // 删除时间
}
```

## 登录接口

```http
POST /api/v1/login HTTP/1.1
Host: :8080
Content-Type: application/json
Content-Length: 56

{
    "username": "admin",
    "password": "123456"
}
```

## jwt验证header,获取文章列表示例

```http
GET /api/v1/article?page=1&pageSize=1 HTTP/1.1
Host: :8080
token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6IkFkbWluIiwiaWF0IjoxNjc0NjYzMjM5LCJleHAiOjE2NzQ2NjYzMzl9.8W45GJQqV656
```

## 控制器示例

```go
package v1

import (
	"encoding/json"
	"gin-base/app/model"
	"gin-base/app/service"
	"gin-base/app/validate"
	"gin-base/common"
	"gin-base/common/global"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ArticleController struct {
	common.BaseController
}

// @Tags    文章
// @Summary 列表
// @Router  /v1/article [get]
func (s *ArticleController) List(c *gin.Context) {
	var (
		articleService service.ArticleService
		req            validate.ArticleValidate
	)

	err := c.ShouldBindQuery(&req)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	// 验证
	err = validate.GetArticleValidate(req, "list")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error(), nil)
		return
	}

	pageData, err := articleService.List(req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error(), nil)
		return
	}

	s.ApiResponse(c, global.Success, "获取成功", pageData)
}
```

## 验证器以及验证场景

```go
package validate

import (
	"errors"
	validator "github.com/gookit/validate"
)

// 文章请求验证
type ArticleValidate struct {
	Page     int    `form:"page" validate:"required|int|gt:0" label:"页码"`
	PageSize int    `form:"pageSize" validate:"required|int|gt:0" label:"每页数量"`
	ID       int64  `json:"id" validate:"required" label:"ID"`
	Title    string `json:"title" validate:"required" label:"标题"`
	Content  string `json:"content" validate:"required" label:"内容"`
}

// 请求验证
func GetArticleValidate(data ArticleValidate, scene string) error {
	v := validator.Struct(data, scene)
	if !v.Validate(scene) {
		return errors.New(v.Errors.One())
	}

	return nil
}

// ConfigValidation 配置验证
// - 定义验证场景
// - 也可以添加验证设置
func (a ArticleValidate) ConfigValidation(v *validator.Validation) {
	v.WithScenes(validator.SValues{
		"list":   []string{"Page", "PageSize"},
		"create": []string{"Title", "Content"}, // []string{"User.FullName", "Title"}
		"update": []string{"ID", "Title", "Content"},
		"detail": []string{"ID"},
		"delete": []string{"ID"},
	})
}

// Messages 您可以自定义验证器错误消息
func (s ArticleValidate) Messages() map[string]string {
	return validator.MS{
		"required":    "字段 {field} 必填",
		"int":         "字段 {field} 必须为整数",
		"Page.gt":     "字段 {field} 需大于 0",
		"PageSize.gt": "字段 {field} 需大于 0",
	}
}

// Translates 你可以自定义字段翻译
func (s ArticleValidate) Translates() map[string]string {
	return validator.MS{
		"Page":     "页码",
		"PageSize": "每页数量",
		"ID":       "ID",
		"Title":    "标题",
		"Content":  "内容",
	}
}
```

## 路由以及中间件

```go
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

			// 创建用户
			v1.POST("/user", user_controller.Create)

			// 更新用户
			v1.PUT("/user/:id", user_controller.Update)

			// 用户详情
			v1.GET("/user/:id", user_controller.Detail)

			// 删除用户
			v1.DELETE("/user/:id", user_controller.Delete)

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
```