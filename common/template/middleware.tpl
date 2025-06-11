package {{.Package}}

import (
    "github.com/gin-gonic/gin"
    "gin-base/common/base"
)

type {{.Name}} struct {
	base.BaseMiddleware
}

// {{.Name}}Middleware {{.Description}}
func (s {{.Name}}) {{.Name}}Middleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Define your middleware logic here
        c.Next()
    }
}