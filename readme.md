English | [中文](readme_zh.md)
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
![img_3.png](./img_3.png)
![img_4.png](./img_4.png)

## Introduction to the Gin-Base Project
- Command line generation
- - Model
- - Controller
- - Service
- - Validate
- - Middleware
- - Route
- Middleware
- - Cors
- - Language
- - JWT
- - Log
- Validators
- - custom validation scenarios
- Jwt authentication
- Cache
- - Memory Cache
- - Redis Cache
- - Disk Cache
- Event
- - Sql Listener
- - Http Listener
- Log
- - Error Info
- - Trace Info
- - Sql Record
- - Http Record
- Language
- - At present, only login related modules have cases, supporting only Chinese and English. If you need other modules or languages, please expand them yourself
- Swagger
- Air
- …

## Backend Technologies Used
- Gin
- Gorm
- Jwt
- Mysql
- Middleware
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

### Thank you for providing the front-end template by vue-next-admin
Project Address: https://gitee.com/lyt-top/vue-next-admin

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
│   ├── base                            # Base
│   ├── extend                          # Extend
│   ├──├── cache                        # Cache
│   ├──├── context                      # Context
│   ├──├── i18n                         # Language
│   ├──├──├── locales                   # Translate Files
│   ├── global                          # Global
│   ├── template                        # Template
├── config                              # Configuration
├── database                            # Database
├── docs                                # Documents
├── helper                              # Utils
├── storage                             # Storage
│   ├── cache                           # Disk Cache
│   ├── logs                            # Logs
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
cd path/gin-base
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go get -u
go mod tidy
go mod vendor 
go run main.go # OR air
```
#### Frontend
```bash
# Frontend Run Port:3000 Address:127.0.0.1:3000
cd path/gin-base/web
npm install
npm run dev
```
### Command Generation

```bash
# Command make:<model|controller|service|validate|middleware>
# Generative Model 
# --table=<user(your table name)> --camel=true|false(true:Generate camel hump field,false:Generate underline field) --file=<admin/test(The generated file path)>
go run ./cli/main.go make:model --table=user --camel=true --file=admin/test

# Generative Controller
# --file=<v1/test(The generated file path)> --func=<List|Create|Update|Delete|Detail ...(Action name)> --method=<get|post|put|delete(request method)> --router=</v1/user(access route)> --desc=<Action annotation>
go run ./cli/main.go make:controller --file=v1/test --func=List --method=get --router=/v1/list --desc=test

# Generative Service
# --file=<admin/test(The generated file path)> --func=<List|Create|Update|Delete|Detail ...(Action name)> --desc=<Action annotation>
go run ./cli/main.go make:service --file=admin/test --func=List --desc=test

# Generative validate
# --file=<test(The generated file path)> --desc=<Action annotation>
go run ./cli/main.go make:validate --file=test --desc=test

# Generative Middleware
# --file=<admin/test(The generated file path)> --desc=<Action annotation>
go run ./cli/main.go make:middleware  --file=admin/test --desc=test 

# Generative Router
# --file=<admin/test(The generated file path)>
go run ./cli/main.go make:router  --file=admin/test
```
### Example of Generating a Model Structure

```go
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"gorm.io/gorm"
)

const TableNameArticle = "article"

// Article mapped from table <article>
type Article struct {
	ID         int64          `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	UID        int64          `gorm:"column:uid;type:int(11);not null;comment:用户id" json:"uid"`                           // 用户id
	User       *User          `json:"user" gorm:"foreignkey:uid;references:id"`                                           // 关联用户
	Title      string         `gorm:"column:title;type:varchar(50);not null;comment:标题" json:"title"`                     // 标题
	Content    string         `gorm:"column:content;type:varchar(255);not null;comment:内容" json:"content"`                // 内容
	CategoryID int64          `gorm:"column:category_id;type:int(11);not null;comment:分类id" json:"categoryId"`            // 分类id
	DataSource int64          `gorm:"column:data_source;type:int(11);not null;comment:数据来源 1=文章库 2=自建" json:"dataSource"` // 数据来源 1=文章库 2=自建
	IsPublish  int64          `gorm:"column:is_publish;type:int(11);not null;comment:是否发布 1=已发布 2=未发布" json:"IsPublish"`  // 是否发布 1=已发布 2=未发布
	Category   *Category      `json:"category" gorm:"foreignkey:category_id;references:id"`                               // 关联分类
	Tag        JsonString     `gorm:"column:tag;type:json;comment:标签" json:"tag"`                                         // 标签
	CreatedAt  *JsonTime      `gorm:"column:created_at;type:datetime;comment:创建时间" json:"createdAt"`                      // 创建时间
	UpdatedAt  *JsonTime      `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updatedAt"`                      // 更新时间
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间" json:"deletedAt"`                      // 删除时间
}

// TableName Article's table name
func (*Article) TableName() string {
	return TableNameArticle
}

```

## Setter and Getter
```go
type Meta struct {
    Title       string  `gorm:"column:title;type:json;comment:菜单名称" json:"title"`
    Icon        string  `gorm:"column:icon;type:json;comment:菜单图标" json:"icon"`
    IsHide      bool    `gorm:"column:isHide;type:json;comment:是否隐藏" json:"isHide"`
    IsKeepAlive bool    `gorm:"column:isKeepAlive;type:json;comment:是否缓存" json:"isKeepAlive"`
    IsAffix     bool    `gorm:"column:isAffix;type:json;comment:是否固定" json:"isAffix"`
    IsLink      string  `gorm:"column:isLink;type:json;comment:外链/内嵌时链接地址" json:"isLink"` // 外链/内嵌时链接地址（http:xxx.com），开启外链条件，`1、isLink: 链接地址不为空`
    IsIframe    bool    `gorm:"column:isIframe;type:json;comment:是否内嵌" json:"isIframe"`   // 是否内嵌，开启条件，`1、isIframe:true 2、isLink：链接地址不为空`
    Roles       []int64 `gorm:"column:roles;type:json;comment:菜单角色" json:"roles"`         // 权限标识，取角色管理
}

func (m Meta) Value() (driver.Value, error) {
    return json.Marshal(m)
}

// Implement the sql.Scanner interface (for reading from the database)
func (m *Meta) Scan(value interface{}) error {
    if value == nil {
        *m = Meta{}
        return nil
    }
    var data []byte
    switch v := value.(type) {
    case string:
    data = []byte(v)
    case []byte:
    data = v
    default:
        return fmt.Errorf("cannot scan type %T into Meta", value)
    }
    return json.Unmarshal(data, m)
}
```
## Model Dynamic Filtering
### Structure Filtering
```go
package validate

...
// UserSearchQuery query
type UserSearchQuery struct {
	Username string `form:"username" validate:"required" label:"用户名"`
	FullName string `form:"fullName" validate:"required" label:"姓名"`
	Nickname string `form:"nickname" validate:"required" label:"昵称"`
	Gender   int64  `form:"gender" validate:"required|int" label:"性别"`
	// JSON fields such as 'user' is JSON field in the database, user = {gender:1,age:18}
	UserGender string `form:"user.gender" validate:"required" label:"性别"`
}

// UserSearchBody body
type UserSearchBody struct {
	Username string `json:"username" validate:"required" label:"用户名"`
	FullName string `json:"fullName" validate:"required" label:"姓名"`
	Nickname string `json:"nickname" validate:"required" label:"昵称"`
	Gender   int    `json:"gender" validate:"required|int" label:"性别"`
	// JSON fields such as 'user' is JSON field in the database, user = {gender:1,age:18}
	UserGender string `json:"user.gender" validate:"required" label:"性别"`
}

...
package service
import (
    "gin-base/app/model"
)

...
var (
	models []model.user
)

data := global.DB.Scopes(model.Search(search)).Find(&models) 
fmt.Printf("data:%v",data)
...
```
### Map filtering suggests using post to place body requests for complex filtering, while getting requires assembling front-end data by oneself
```json
// operator: >,<, >=, <=, =, !=, in, not in, between, not between, is null, is not null, like, not like, left like, right like
// Front end request 
//	__search: {
//	   "field": {"operator": ">", "value": 10},
//	   "field": {"operator": "<", "value": 10},
//	   "field": {"operator": ">=", "value": 10},
//	   "field": {"operator": "<=", "value": 10},
//	   "field": {"operator": "=", "value": 10},
//	   "field": {"operator": "!=", "value": 10},
//	   "field": {"operator": "in", "value": [1,2,3]},
//	   "field": {"operator": "not in", "value": [1,2,3]},
//	   "field": {"operator": "between", "value": [1,2]},
//	   "field": {"operator": "not between", "value": [1,2]},
//	   "field": {"operator": "is null"},
//	   "field": {"operator": "is not null"},
//	   "field": {"operator": "like", "value": "admin"},
//	   "field": {"operator": "left like", "value": "test"},
//	   "field": {"operator": "right like", "value": "138"},
//     "or":    [
//    		"field": {"operator": ">", "value": 10},
//	   		"field": {"operator": "<", "value": 10},
//	   		"field": {"operator": ">=", "value": 10},
//          ...
//     ]
//     "and":   [
//    		"field": {"operator": ">", "value": 10},
//	   		"field": {"operator": "<", "value": 10},
//	   		"field": {"operator": ">=", "value": 10},
//          ...
//     ]
//	   "jsonField.field": {"operator": ">", "value": 10},
//	   ...
//	}
__search: {
  "username": {
    "operator": "like", 
    "value": "admin"
  },
  "fullName": {
    "operator": "=",
    "value": "zhangsan"  
  },
  "nickname": {
    "operator": "=",
    "value": "nickname"   
  },
  "gender": {
    "operator": "=",
    "value": 1
  },
  // JSON fields such as 'user' is JSON field in the database, user = {gender:1,age:18}
  "user.age": {
    "operator": ">",
    "value": 18
  }
}
```

```go
package service

...
var (
	search map[string]interface{}
	models []model.user
)

...
data := global.DB.Scopes(model.SearchMap(search)).Find(&models)
fmt.Printf("data:%v",data)

...
```

## Cache usage Support for memory caching and disk caching and Redis caching needs to be specified in YAML
```yaml
# Cache
cache:
  type: "redis"  # memory|disk|redis
  redis:
    address: "127.0.0.1:6379"
    password: ""  # If the password is empty, it is not necessary
    db: 1
```
### Cache call example
```go
package service

import (
	"gin-base/common/base"
	"gin-base/common/global"
	"time"
)

type CacheService struct {
	base.BaseService
}

// SetCache Set cache
// @param key string
// @param value interface{}
// @param expire time.Duration
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
	"gin-base/common/base"
	"gin-base/common/global"
	"gin-base/helper"
	"gorm.io/gorm"
)

type LoginService struct {
	base.BaseService
}

// Login Login
// @param username string
// @param password string
// @return m model.User, error
func (s *LoginService) Login(username string, password string) (m model.User, err error) {
	if err = global.DB.Where("username = ?", username).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return m, errors.New("Login account error")
		}
	}

	check := helper.BcryptCheck(password, m.Password)
	if !check {
		return m, errors.New("Login password error")
	}

	if m.Status != 1 {
		return m, errors.New("Account has been disabled")
	}

	// Publish an event
	e := global.Config.Event
	e.Name = "userLogin"
	e.Data = map[string]interface{}{
		"username": username,
		"password": password,
	}
	global.Event.Publish(e)

	return m, nil
}

```

### Received event
```go
package main

import (
	"fmt"
	"gin-base/app/middleware"
	"gin-base/common/global"
	"gin-base/config"
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
	gin.SetMode(global.Config.Service.Mode)

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
	if global.Config.Cors.Enabled {
		router.Use(middleware.Cors{}.Handle())
	}

	// Global log middleware
	router.Use(middleware.Logger{}.Handle())

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
func onEventReceived(event config.Event, timestamp time.Time) {
	// todo Process Event
	fmt.Printf("Event received at %s: name: %s, data: %v\n", timestamp.Format(time.RFC3339), event.Name, event.Data)
}
```

## Language
```go
package v1

import (
    "gin-base/app/middleware"
    "gin-base/app/service"
    "gin-base/app/validate"
    "gin-base/common/base"
    "gin-base/common/extend/i18n"
    "gin-base/common/global"
    "github.com/gin-gonic/gin"
)

// Login 登录
// @Tags 登录相关
// @Summary 登录
// @Description 用户登录
// @Accept json
// @Produce json
// @Param data body validate.Login true "登录参数"
// @Router /api/v1/login [post]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *LoginController) Login(c *gin.Context) {
    var (
        loginService  service.LoginService
        loginValidate validate.Login
        jwt           middleware.Jwt
    )
    
    err := c.ShouldBind(&loginValidate)
    if err != nil {
        s.ApiResponse(c, global.SystemError, err.Error())
        return
    }
    
    // 验证
    err = validate.Login{}.GetValidate(c, loginValidate, "login")
        if err != nil {
        s.ApiResponse(c, global.ArgsError, err.Error())
        return
    }
    
    // 验证码校验
    b := s.verify(loginValidate.CaptchaID, loginValidate.Code)
    if !b {
        s.ApiResponse(c, global.ArgsError, i18n.T(c, "login.codeErr", nil))
        return
    }
    
    userModel, err := loginService.Login(loginValidate.Username, loginValidate.Password)
    if err != nil {
        s.ApiResponse(c, global.ArgsError, err.Error())
        return
    }
    
    token, rToken, tExpire, rExpire, err := jwt.WithRefresh(userModel.ID, 2*60*60, 2*24*60*60)
    if err != nil {
        s.ApiResponse(c, global.ArgsError, err.Error())
        return
    }

	s.ApiResponse(c, global.Success, i18n.T(c, "login.codeErr", map[string]interface{}{
		"name": userModel.Username,
	}), map[string]interface{}{
		"token": map[string]interface{}{
			"accessToken":   token,
			"refreshToken":  rToken,
			"expire":        tExpire,
			"refreshExpire": rExpire,
		},
		"user": userModel,
	})
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
	"gin-base/common/base"
	"gin-base/common/global"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ArticleController struct {
	base.BaseController
}

// List List
// @Tags Article
// @Summary List
// @Description Article Lisy
// @Param token header string true "Auth Token"
// @Param page query string true "Page"
// @Param pageSize query string true "Page Size"
// @Success 200 {object} global.Response{global.Success} "SUCCESS" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "Arg Error" Example({"code":400,"msg":"Arg Error","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "Sys Error" Example({"code":500,"msg":"Sys Error","data":[]})
// @Router /api/v1/article [get]
func (s *ArticleController) List(c *gin.Context) {
	var (
		articleService  service.ArticleService
		articleValidate validate.Article
	)

	err := c.ShouldBindQuery(&articleValidate)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// validate
	err = validate.Article{}.GetValidate(articleValidate, "list")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	pageData, err := articleService.List(articleValidate)
	if err != nil {
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

// Article Article Request Validation
type Article struct {
	Page     int    `form:"page" validate:"required|int|gt:0" label:"页码"`
	PageSize int    `form:"pageSize" validate:"required|int|gt:0" label:"每页数量"`
	ID       int64  `json:"id" validate:"required" label:"ID"`
	Title    string `json:"title" validate:"required" label:"标题"`
	Content  string `json:"content" validate:"required" label:"内容"`
}

// GetValidate
// Request Validation
func (s Article) GetValidate(data Article, scene string) error {
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
func (s Article) ConfigValidation(v *validator.Validation) {
	v.WithScenes(validator.SValues{
		"list":   []string{"Page", "PageSize"},
		"create": []string{"Title", "Content"}, 
		"update": []string{"ID", "Title", "Content"},
		"detail": []string{"ID"},
		"delete": []string{"ID"},
	})
}

// Messages 
// Customizing Validator Error Messages
func (s Article) Messages() map[string]string {
	return validator.MS{
		"required":    "Filed {field} Required",
		"int":         "Filed {field} Must be an integer",
		"Page.gt":     "Filed {field} Must be greater than 0",
		"PageSize.gt": "Filed {field} Must be greater than 0",
	}
}

// Translates 
// Customizing Field Translations
func (s Article) Translates() map[string]string {
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

## Swagger Document generation
```bash
# Quickly generate document command
go install github.com/swaggo/swag/cmd/swag@latest
swag init -g main.go --exclude cli,app/service
```
![img.png](./img.png)