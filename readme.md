# Important Notice
This document was compiled and written by Master Brother with email 25076778@qq.com. Unauthorized reproduction is strictly prohibited.
This project is prohibited from being used for illegal businesses such as viruses, trojans, pornography, gambling, fraud, prohibited items, counterfeit products, false information, cryptocurrencies, finance, etc.

The current project is for personal learning and testing only. All commercial online behaviors and illegal uses are strictly forbidden!!!

## Project Address
- Github: https://github.com/dsxwk/gin-base.git
- Gitee: https://gitee.com/dsxwk/gin-base.git

## Introduction to the Gin Framework
Gin is a web framework written in Go. It is known for its simplicity, speed, and efficiency, and is widely used in Go web development.

## Key features of the Gin framework include:

- Fast: Based on the standard library net/http, using goroutine and channel for asynchronous processing, enhancing performance.
- Simple: Provides a series of APIs and middleware, allowing developers to quickly build web applications.
- Efficient: Uses sync.Pool to cache objects, reducing memory allocation and release, thereby improving performance.

Golang Gin is a lightweight and efficient Golang web framework. It is widely used in the development of various web applications due to its high performance, ease of use, and flexibility.

## Project screenshot

![img.png](./img.png)
![img_1.png](./img_1.png)
![img_2.png](./img_2.png)

## Introduction to the Gin-Base Project
- Command line generation for quick creation of model, controller, service, validator, middleware
- Validators and custom validation scenarios
- Jwt authentication
- Air
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
go run main.go # OR air
```
#### Frontend
```bash
# Frontend Run Port:3000 Address:127.0.0.1:3000
cd path/to/your/frontend
npm install
npm run dev
```
### Command Generation

```bash
# Common Parameters --make=<model|controller|service|validate|middleware>
# Generative Model 
# --tableName=<user(your table name)> --camel=true|false(true:Generate camel hump field,false:Generate underline field)
go run ./cli/main.go --make=model --tableName=user --camel=true

# Generative Controller
# --fileName=</app/controller/v1/test(The generated file path)> --function=<List|Create|Update|Delete|Detail ...(Action name)> --method=<get|post|put|delete(request method)> --router=</v1/user(access route)> --description=<Action annotation>
go run ./cli/main.go --make=controller  --fileName=/app/controller/v1/test --function=List --method=get --router=/v1/list --description=test

# Generative Service
# --fileName=</app/service/test(The generated file path)> --function=<List|Create|Update|Delete|Detail ...(Action name)> --description=<Action annotation>
go run ./cli/main.go --make=service  --fileName=/app/service/test --function=List --description=test

# Generative validate
# --fileName=</app/validate/test(The generated file path)> --description=<Action annotation>
go run ./cli/main.go --make=validate  --fileName=/app/validate/test --description=test

# Generative Middleware
# --fileName=</app/middleware/test(The generated file path)> --description=<Action annotation>
go run ./cli/main.go --make=middleware  --fileName=/app/middleware/test --description=test 
```
### Example of Generating a Model Structure

```go
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

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

type ArticleQuery struct {
	ID         int64     `json:"id" comment:"ID"`                                                     // ID
	UID        int64     `json:"uid" comment:"用户id"`                                                  // 用户id
	User       *User     `json:"user" gorm:"foreignkey:uid;references:id" comment:"关联用户"`             // 关联用户
	Title      string    `json:"title" comment:"标题"`                                                  // 标题
	Content    string    `json:"content" comment:"内容"`                                                // 内容
	CategoryID int64     `json:"category_id" comment:"分类id"`                                          // 分类id
	DataSource int64     `json:"data_source" comment:"数据来源 1=文章库 2=自建"`                               // 数据来源 1=文章库 2=自建
	IsPublish  int64     `json:"is_publish" comment:"是否发布 1=已发布 2=未发布"`                               // 是否发布 1=已发布 2=未发布
	Category   *Category `json:"category" gorm:"foreignkey:category_id;references:id" comment:"关联分类"` // 关联分类
	Tag        []string  `json:"tag" comment:"标签"`                                                    // 标签
}

// TableName Article's table name
func (*Article) TableName() string {
	return TableNameArticle
}

// Before creation
func (s *Article) BeforeCreate(tx *gorm.DB) (err error) {
	if s.CreatedAt == nil {
		createdAt := time.Now().Format("2006-01-02 15:04:05")
		s.CreatedAt = &createdAt
	}

	if s.UpdatedAt == nil {
		updatedAt := time.Now().Format("2006-01-02 15:04:05")
		s.UpdatedAt = &updatedAt
	}

	return
}

// Before updating
func (s *Article) BeforeUpdate(tx *gorm.DB) (err error) {
	if s.UpdatedAt == nil {
		updatedAt := time.Now().Format("2006-01-02 15:04:05")
		s.UpdatedAt = &updatedAt
	}

	return
}

// After querying
func (s *Article) AfterFind(tx *gorm.DB) (err error) {
	// Time format conversion
	if s.CreatedAt != nil {
		createdAt, _ := time.Parse(time.RFC3339, *s.CreatedAt)
		// Format time The Time type is a string, and it is reassigned to a field of type * string
		formattedCreatedAt := createdAt.Format("2006-01-02 15:04:05")
		s.CreatedAt = &formattedCreatedAt
	}

	if s.UpdatedAt != nil {
		updatedAt, _ := time.Parse(time.RFC3339, *s.UpdatedAt)
		// Format time The Time type is a string, and it is reassigned to a field of type * string
		formattedUpdatedAt := updatedAt.Format("2006-01-02 15:04:05")
		s.UpdatedAt = &formattedUpdatedAt
	}

	if s.DeletedAt != nil {
		deletedAt, _ := time.Parse(time.RFC3339, *s.DeletedAt)
		// Format time The Time type is a string, and it is reassigned to a field of type * string
		formattedDeletedAt := deletedAt.Format("2006-01-02 15:04:05")
		s.DeletedAt = &formattedDeletedAt
	}

	return
}

// Before deletion
func (s *Article) BeforeDelete(tx *gorm.DB) (err error) {
	if s.DeletedAt == nil {
		deletedAt := time.Now().Format("2006-01-02 15:04:05")
		s.DeletedAt = &deletedAt
	}

	return
}
```

## Setter and Getter
```go
// Get tags
func (s *Article) GetTag() []string {
	if s != nil && s.Tag != nil {
		var tagJson []string
		json.Unmarshal([]byte(*s.Tag), &tagJson)
		return tagJson
	}
	return nil
}

// Set tags
func (s *Article) SetTag(tag []string) *string {
	var (
		model Article
	)

	if tag != nil {
		tagJSON, _ := json.Marshal(tag)
		tagStr := string(tagJSON)
		model.Tag = &tagStr
	}
	return model.Tag
}
```

### The use of the setter and getter
```go
// @function: List
// @description: list
// @param: req validate.ArticleValidate
// @return: global.PageData, error
func (s *ArticleService) List(req validate.ArticleValidate) (global.PageData, error) {
	var (
		articleModel []model.Article
		articleQuery []model.ArticleQuery
		pageData     global.PageData
		//fields       []field
	)

	// Get pagination defaults to the first page, with 10 records per page
	offset, limit := utils.Pagination(req.Page, req.PageSize)

	// join
	//db := global.DB.Joins("LEFT JOIN user ON article.uid = user.id LEFT JOIN category ON article.category_id = category.id").Select("article.*, user.username, category.name").Find(&articleModel).scan(&fields)

	db := global.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username, full_name, nickname, email, gender, age")
	}).Preload("Category", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).Find(&articleModel)

	// Obtain the total number of records
	err := db.Count(&pageData.Total).Error
	if err != nil {
		return pageData, err
	}

	// Execute pagination query
	err = db.Offset(offset).Limit(limit).Find(&articleModel).Scan(&articleQuery).Error
	if err != nil {
		return pageData, err
	}

	for k, m := range articleModel {
		articleQuery[k].User = m.User
		articleQuery[k].Category = m.Category
		articleQuery[k].Tag = m.GetTag()
	}

	pageData.Page = req.Page
	pageData.PageSize = req.PageSize
	pageData.List = articleQuery

	return pageData, nil
}

// @function: Update
// @description: update
// @param: req model.Article
// @return: model.Article, error
func (this *ArticleService) Update(req model.ArticleQuery) (model.Article, error) {
    var (
        articleModel model.Article
    )
    
    err := copier.Copy(&articleModel, &req)
    if err != nil {
        return articleModel, err
    }
    articleModel.Tag = articleModel.SetTag(req.Tag)
    
    err = global.DB.Updates(&articleModel).Error
    if err != nil {
        return articleModel, err
    }
    
    return articleModel, nil
}
```

## air
### Updating the code will automatically restart without the need for a restart 
```shell
 E:\www\dsx\gin-base> air

  __    _   ___
 / /\  | | | |_)
/_/--\ |_| |_| \_ v1.60.0, built with Go go1.23.1

watching .
watching app
watching app\controller
watching app\controller\v1
watching app\middleware
watching app\model
watching app\service
watching app\validate
watching cli
watching common
watching common\global
watching common\template
watching config
watching database
watching helper
watching helper\utils
watching log
watching resource
watching resource\images
watching routers
!exclude tmp
!exclude vendor
building...
!exclude .git
!exclude .git
running...
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> main.main.func1 (3 handlers)
[GIN-debug] GET    /resource/*filepath       --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /resource/*filepath       --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] POST   /api/v1/login             --> gin-base/app/controller/v1.(*LoginController).Login-fm (4 handlers)
[GIN-debug] GET    /api/v1/user/             --> gin-base/app/controller/v1.(*UserController).List-fm (5 handlers)
[GIN-debug] POST   /api/v1/user/             --> gin-base/app/controller/v1.(*UserController).Create-fm (5 handlers)
[GIN-debug] PUT    /api/v1/user/:id          --> gin-base/app/controller/v1.(*UserController).Update-fm (5 handlers)
[GIN-debug] GET    /api/v1/user/:id          --> gin-base/app/controller/v1.(*UserController).Detail-fm (5 handlers)
[GIN-debug] DELETE /api/v1/user/:id          --> gin-base/app/controller/v1.(*UserController).Delete-fm (5 handlers)
[GIN-debug] GET    /api/v1/article/          --> gin-base/app/controller/v1.(*ArticleController).List-fm (5 handlers)
[GIN-debug] POST   /api/v1/article/          --> gin-base/app/controller/v1.(*ArticleController).Create-fm (5 handlers)
[GIN-debug] PUT    /api/v1/article/:id       --> gin-base/app/controller/v1.(*ArticleController).Update-fm (5 handlers)
[GIN-debug] GET    /api/v1/article/:id       --> gin-base/app/controller/v1.(*ArticleController).Detail-fm (5 handlers)
[GIN-debug] DELETE /api/v1/article/:id       --> gin-base/app/controller/v1.(*ArticleController).Delete-fm (5 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
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
		// User
		user := v1.Group("user").Use(jwtMiddleware)
		{
			// List
			user.GET("/", user_controller.List)
			// Create
			v1.POST("/", user_controller.Create)
			// Update
			v1.PUT("/:id", user_controller.Update)
			// Detail
			v1.GET("/:id", user_controller.Detail)
			// Delete
			v1.DELETE("/:id", user_controller.Delete)
		}
		
		// Article
		article := v1.Group("article").Use(jwtMiddleware)
		{
			// List
			article.GET("/", article_controller.List)
			// Create
			v1.POST("/", article_controller.Create)
			// Update
			v1.PUT("/:id", article_controller.Update)
			// Detail
			v1.GET("/:id", article_controller.Detail)
			// Delete
			v1.DELETE("/:id", article_controller.Delete)
        }
	}
}
```