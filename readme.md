# Important Notice
This document was compiled and written by Master Brother with email 25076778@qq.com. Unauthorized reproduction is strictly prohibited.
This project is prohibited from being used for illegal businesses such as viruses, trojans, pornography, gambling, fraud, prohibited items, counterfeit products, false information, cryptocurrencies, finance, etc.

The current project is for personal learning and testing only. All commercial online behaviors and illegal uses are strictly forbidden!!!

## Project Address
Github: https://github.com/dsxwk/gin-base.git
Gitee: https://gitee.com/wang-ke/gin-base.git

## Introduction to the Gin Framework
Gin is a web framework written in Go. It is known for its simplicity, speed, and efficiency, and is widely used in Go web development.

## Key features of the Gin framework include:

- Fast: Based on the standard library net/http, using goroutine and channel for asynchronous processing, enhancing performance.
- Simple: Provides a series of APIs and middleware, allowing developers to quickly build web applications.
- Efficient: Uses sync.Pool to cache objects, reducing memory allocation and release, thereby improving performance.

Golang Gin is a lightweight and efficient Golang web framework. It is widely used in the development of various web applications due to its high performance, ease of use, and flexibility.

## Introduction to the Gin-Base Project
- Command model generation
- Validators and custom validation scenarios
- Jwt authentication
- …

## Backend Technologies Used
- Gin
- Gorm
- Jwt
- Mysql
- Validator

## Frontend Technologies Used
- Vue3
- Vite
- Element-Plus
- Element-Plus-Table

## Go-Base Directory Structure
### Backend

```bash
├── app                                 # Application
│   ├── controller                      # Controller
│   ├── model                           # Model
│   ├── service                         # Service
│   ├── validate                        # Validation
│   ├── middleware                      # Middleware
├── cli                                 # Command
├── common                              # Common
│   ├── global                          # Global
├── config                              # Configuration
├── docs                                # Documents
├── database                            # Database
├── helper                              # Helper
├── log                                 # Log
├── resource                            # Resource
├── web                                 # Web
├── routers                             # Router
├── vendor                              # Vendor
```

## Web Directory Structure
### Frontend

```bash
├── app                                 # Application
│   ├── modules                         # Modules
├── components                          # Components
├── config                              # Configuration
├── enums                               # Enums
├── node_modules                        # Node Modules
├── public                              # Public
├── docs                                # Documents
├── routers                             # Router
├── src                                 # Src
├── utils                               # Utils
├── views                               # Views
```

## Usage
### Running Frontend and Backend
#### Backend
```bash
# Backend Run Port:8080 Address:127.0.0.1:8080/api/v1/...
cd path/to/your/backend
go mod download
go mod tidy vendor
go run main.go
```
#### Frontend
```bash
# Frontend Run Port:3000 Address:127.0.0.1:3000
cd path/to/your/frontend
npm install
npm run dev
```
### Command Model Generation --path Parameter Default is Fine,No Need to Modify

```bash
go run ./cli/main.go --tableName=user --path=app/temp
go run ./cli/main.go --tableName=article --path=app/temp
go run ./cli/main.go --tableName=category --path=app/temp
go run ./cli/main.go --tableName=system_config --path=app/temp
```
### Example of Generating a Model Structure

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

## Login API Endpoint

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

## JWT Authentication and Fetching Article List Example

```http
GET /api/v1/article?page=1&pageSize=1 HTTP/1.1
Host: :8080
token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6IkFkbWluIiwiaWF0IjoxNjc0NjYzMjM5LCJleHAiOjE2NzQ2NjYzMzl9.8W45GJQqV656
```

## Controller Example

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

// @Tags    Article
// @Summary List
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

	// validate
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

	s.ApiResponse(c, global.Success, "success", pageData)
}
```

## Validators and Validation Scenarios

```go
package validate

import (
	"errors"
	validator "github.com/gookit/validate"
)

// Article Request Validation
type ArticleValidate struct {
	Page     int    `form:"page" validate:"required|int|gt:0" label:"页码"`
	PageSize int    `form:"pageSize" validate:"required|int|gt:0" label:"每页数量"`
	ID       int64  `json:"id" validate:"required" label:"ID"`
	Title    string `json:"title" validate:"required" label:"标题"`
	Content  string `json:"content" validate:"required" label:"内容"`
}

// Request Validation
func GetArticleValidate(data ArticleValidate, scene string) error {
	v := validator.Struct(data, scene)
	if !v.Validate(scene) {
		return errors.New(v.Errors.One())
	}

	return nil
}

// ConfigValidation 
// Configuration Validation
// - Defining Validation Scenarios
// - Adding Validation Settings
func (a ArticleValidate) ConfigValidation(v *validator.Validation) {
	v.WithScenes(validator.SValues{
		"list":   []string{"Page", "PageSize"},
		"create": []string{"Title", "Content"}, // []string{"User.FullName", "Title"}
		"update": []string{"ID", "Title", "Content"},
		"detail": []string{"ID"},
		"delete": []string{"ID"},
	})
}

// Messages 
// Customizing Validator Error Messages
func (s ArticleValidate) Messages() map[string]string {
	return validator.MS{
		"required":    "Filed {field} Required",
		"int":         "Filed {field} Must be an integer",
		"Page.gt":     "Filed {field} Must be greater than 0",
		"PageSize.gt": "Filed {field} Must be greater than 0",
	}
}

// Translates 
// Customizing Field Translations
func (s ArticleValidate) Translates() map[string]string {
	return validator.MS{
		"Page":     "Page",
		"PageSize": "Page Size",
		"ID":       "ID",
		"Title":    "Title",
		"Content":  "Content",
	}
}
```

## Setting Up Custom Routes and Middleware

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

// Load routes
func LoadRouters(router *gin.Engine) {
	/*// Unified routing grouping
	v1 := router.Group("api/v1")

	// Load routes...
	// Login
	LoginRoutes(v1)*/

	// Login
	login_controller := controller.LoginController{}

	// User
	user_controller := controller.UserController{}

	// Article
	article_controller := controller.ArticleController{}

	// Unified routing grouping
	v1 := router.Group("api/v1")
	{
		// No token verification required
		// User login
		v1.POST("/login", login_controller.Login)

		// Token verification is required
		v1.Use(jwtMiddleware)
		{
			// User List
			v1.GET("/user", user_controller.List)

			// Create User
			v1.POST("/user", user_controller.Create)

			// Update User
			v1.PUT("/user/:id", user_controller.Update)

			// User Detail
			v1.GET("/user/:id", user_controller.Detail)

			// Delete User
			v1.DELETE("/user/:id", user_controller.Delete)

			// Article List
			v1.GET("/article", article_controller.List)

			// Create Article
			v1.POST("/article", article_controller.Create)

			// Update Article
			v1.PUT("/article/:id", article_controller.Update)

			// Article Detail
			v1.GET("/article/:id", article_controller.Detail)

			// Delete Article
			v1.DELETE("/article/:id", article_controller.Delete)
		}
	}
}
```