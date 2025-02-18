package {{.Package}}

import (
    "github.com/gin-gonic/gin"
)

// @Description {{.Description}}
func {{.Name}}Middleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Define your middleware logic here
        c.Next()
    }
}