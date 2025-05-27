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
- Cache
- Event
- Swagger
- Air
- …

## Backend Technologies Used
- Gin
- Gorm
- Jwt
- Mysql
- Validator
- Cache
- Event
- Swagger

## Frontend Technologies Used
- vue-next-admin
- Vue3
- Vite
- Element-Plus
- Element-Plus-Table
- ...

### 感谢vue-next-admin提供的前端模版
项目地址: https://gitee.com/lyt-top/vue-next-admin

## Go-Base Directory Structure
### Backend

```bash
├── app                                 # Applicant
│   ├── controller                      # Controller
│   ├── model                           # Model
│   ├── service                         # Service
│   ├── validate                        # Validation
│   ├── middleware                      # Middleware
├── cli                                 # Command
├── common                              # Common Module
│   ├── extend                          # Extend
│   ├──├── cache                        # Cache
│   ├──├── event                        # Event
│   ├── global                          # Global
│   ├── template                        # Template
├── config                              # Configuration
│   ├── database                        # Database
├── docs                                # Documents
├── helper                              # Utils
├── log                                 # Log
├── resource                            # Resource
├── web                                 # Web Service
├── routers                             # Router
├── vendor                              # Vendor
```

## Web Directory Structure
### Frontend

```bash
├── api                                 # Api
├── components                          # Modules
├── dict                                # Dictionary
├── directive                           # Directive
├── layouts                             # Layout
├── public                              # Public
├── router                              # Router
├── static                              # Static Resource
├── stores                              # Cache
├── them                                # Them
├── docs                                # Documents
├── types                               # ts
├── utils                               # Utils
│   ├── error                           # Global Error Handle
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

# Generative Router
# --fileName=</routers/test(The generated file path)>
go run ./cli/main.go --make=router  --fileName=/routers/test
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
	ID         int64          `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true;comment:ID" json:"id"`  // ID
	UID        int64          `gorm:"column:uid;type:int(11);not null;comment:用户id" json:"uid"`                            // 用户id
	User       *User          `json:"user" gorm:"foreignkey:uid;references:id"`                                            // 关联用户
	Title      string         `gorm:"column:title;type:varchar(50);not null;comment:标题" json:"title"`                      // 标题
	Content    string         `gorm:"column:content;type:varchar(255);not null;comment:内容" json:"content"`                 // 内容
	CategoryID int64          `gorm:"column:category_id;type:int(11);not null;comment:分类id" json:"category_id"`            // 分类id
	DataSource int64          `gorm:"column:data_source;type:int(11);not null;comment:数据来源 1=文章库 2=自建" json:"data_source"` // 数据来源 1=文章库 2=自建
	IsPublish  int64          `gorm:"column:is_publish;type:int(11);not null;comment:是否发布 1=已发布 2=未发布" json:"is_publish"`  // 是否发布 1=已发布 2=未发布
	Category   *Category      `json:"category" gorm:"foreignkey:category_id;references:id"`                                // 关联分类
	Tag        *string        `gorm:"column:tag;type:json;comment:标签" json:"tag"`                                          // 标签
	CreatedAt  *JsonTime      `gorm:"column:created_at;type:datetime;comment:创建时间" json:"created_at"`                      // 创建时间
	UpdatedAt  *JsonTime      `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`                      // 更新时间
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间" json:"deleted_at"`                      // 删除时间
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
	CreatedAt  *JsonTime `gorm:"column:created_at;type:datetime;comment:创建时间" json:"created_at"`      // 创建时间
	UpdatedAt  *JsonTime `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`      // 更新时间
}

// TableName Article's table name
func (*Article) TableName() string {
	return TableNameArticle
}

// BeforeCreate Before creation
func (s *Article) BeforeCreate(tx *gorm.DB) (err error) {
	now := JsonTime(time.Now())
	s.CreatedAt = &now
	s.UpdatedAt = &now
	return nil
}

// BeforeUpdate Before updating
func (s *Article) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	s.UpdatedAt = &now
	return nil
}
```

## Setter and Getter
```go
// GetTag Get tags
func (s *Article) GetTag() []string {
	if s != nil && s.Tag != nil {
		var tagJson []string
		json.Unmarshal([]byte(*s.Tag), &tagJson)
		return tagJson
	}
	return nil
}

// SetTag Set tags
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
// List list
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

// Update update
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

## Cache usage Support for memory caching and Redis caching needs to be specified in YAML
```yaml
# 缓存
cache:
  type: "redis"  # memory OR "redis"
  redis:
    address: "127.0.0.1:6379"
    password: ""  # If the password is empty, it is not necessary
    db: 1
```
### Cache call example
```go
package service

import (
	"gin-base/common"
	"gin-base/common/global"
	"time"
)

type CacheService struct {
	common.BaseService
}

// SetCache Set cache
// @param: key string, value interface{}, expire time.Duration
// @return: bool, error
func (s *CacheService) SetCache(key string, value interface{}, expire time.Duration) (bool, error) {
	err := global.Cache.SetCache(key, value, expire)
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetCache Get cache
// @param: key string
// @return: interface{}, bool
func (s *CacheService) GetCache(key string) (interface{}, bool) {
	res, ok := global.Cache.GetCache(key)
	if ok {
		return res, ok
	}

	return false, ok
}

// DeleteCache Delete cache
// @param: key string
// @return: bool, error
func (s *CacheService) DeleteCache(key string) (bool, error) {
	err := global.Cache.DeleteCache(key)
	if err != nil {
		return false, err
	}

	return true, nil
}

```

## Event usage
### Release event note: Taking login as an example, you only need to add the release event where you want to transmit data
```go
package service

import (
	"errors"
	"gin-base/app/model"
	"gin-base/common"
	"gin-base/common/global"
	"gin-base/helper"
	"gin-base/helper/utils"
	"gorm.io/gorm"
)

type LoginService struct {
	common.BaseService
}

// Login Login
// @param: username string, password string
// @return: model.User, error
func (s *LoginService) Login(username string, password string) (model.User, error) {
	var (
		userModel model.User
	)

	if err := global.DB.Where("username = ?", username).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userModel, errors.New("Login account error")
		}
	}

	check := utils.BcryptCheck(password, userModel.Password)
	if !check {
		return userModel, errors.New("Login password error")
	}

	if userModel.Status != 1 {
		return userModel, errors.New("Account has been disabled")
	}

	// Publish an event
	event := helper.Event{
		Name: "user_login",
		Data: map[string]interface{}{
			"username": username,
			"password": password,
		},
	}
	global.Event.Publish(event)

	return userModel, nil
}

```

### Received event
```go
package main

import (
	"fmt"
	"gin-base/app/middleware"
	"gin-base/common/global"
	"gin-base/helper"
	"gin-base/routers"
	"github.com/gin-gonic/gin"
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
	// Run environment mode debug mode, Test mode, Release production mode, default is debug, read based on the current configuration file
	gin.SetMode(global.Config.Env.Mode)

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "pong",
			"data":    []string{},
		})
	})

	// Static files
	router.StaticFS("/resource", http.Dir("./resource"))
	// Set the maximum memory that can be used for HTTP request processing file upload to 90MB
	router.MaxMultipartMemory = 90 << 20

	// Set up cross domain settings
	router.Use(Cors())

	// Global log middleware
	router.Use(middleware.LoggerMiddleware())

	// Register all events
	global.Event.RegisterAllEvent(onEventReceived)

	// Load route
	routers.LoadRouters(router)

	err := router.Run(`:` + global.Config.Env.Port)
	if err != nil {
		fmt.Println("Service startup failed with error message:", err)
	}
}

// onEventReceived Receive events
func onEventReceived(event helper.Event, timestamp time.Time) {
	// todo Process Event
	fmt.Printf("Event received at %s: name: %s, data: %v\n", timestamp.Format(time.RFC3339), event.Name, event.Data)
}

// Cors cross-domain requests
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", global.Config.Cors.AllowOrigin)
		c.Header("Access-Control-Allow-Headers", global.Config.Cors.AllowHeaders)
		c.Header("Access-Control-Expose-Headers", global.Config.Cors.ExposeHeaders)
		c.Header("Access-Control-Allow-Methods", global.Config.Cors.AllowMethods)
		c.Header("Access-Control-Allow-Credentials", global.Config.Cors.AllowCredentials)

		// Release all OPTIONS methods
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// Processing requests
		c.Next()
	}
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

// List List
// @Router /v1/article [get]
func (s *ArticleController) List(c *gin.Context) {
	var (
		articleService service.ArticleService
		req            validate.ArticleValidate
	)

	err := c.ShouldBindQuery(&req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// validate
	err = validate.GetArticleValidate(req, "list")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	pageData, err := articleService.List(req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, pageData)
}
```

## Validators and Validation Scenarios

```go
package validate

import (
	"errors"
	validator "github.com/gookit/validate"
)

// ArticleValidate Article Request Validation
type ArticleValidate struct {
	Page     int    `form:"page" validate:"required|int|gt:0" label:"页码"`
	PageSize int    `form:"pageSize" validate:"required|int|gt:0" label:"每页数量"`
	ID       int64  `json:"id" validate:"required" label:"ID"`
	Title    string `json:"title" validate:"required" label:"标题"`
	Content  string `json:"content" validate:"required" label:"内容"`
}

// GetArticleValidate
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

## Register Route

```go
package routers

import (
	"gin-base/app/controller/v1"
	"github.com/gin-gonic/gin"
)

// UserRouter User Router
type UserRouter struct{}

// RegisterRoutes Implement Router interface
func (r UserRouter) RegisterRoutes(routerGroup *gin.RouterGroup) {
	var (
		controller v1.UserController
	)

	// List
	routerGroup.GET("/user", controller.List)
	// Create
	routerGroup.POST("/user", controller.Create)
	// Update
	routerGroup.PUT("/user/:id", controller.Update)
	// Delete
	routerGroup.DELETE("/user/:id", controller.Delete)
	// Detail
	routerGroup.GET("/user/:id", controller.Detail)
}

```

## Cache

```go
package cache

import (
	"fmt"
	"time"
)

// Test Cache Test
func Test() {
	// Set Cache
	global.Cache.SetCache("test", "test", 10*time.Second)
	// Get Cache
	res := global.Cache.GetCache("test")
	// Delete Cache
	global.Cache.DelCache("test")
	fmt.Printf("%v\n", res)
}
```

## Swagger Document generation note: Currently, only login is required. If you need to expand all documents, you can do so yourself
```bash
# Quickly generate document command
swag init -g main.go --exclude cli
```
