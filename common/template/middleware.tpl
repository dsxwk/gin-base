package {{.Package}}

import (
    "github.com/gin-gonic/gin"
    "gin-base/common"
)

type {{.Name}} struct {
	common.BaseMiddleware
}

// {{.Name}}Middleware {{.Description}}
func {{.Name}}Middleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Define your middleware logic here
        c.Next()
    }
}