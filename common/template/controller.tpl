package {{.Package}}

import (
    "gin-base/common/base"
    "gin-base/common/global"
    "github.com/gin-gonic/gin"
)

type {{.Name}}Controller struct {
    base.BaseController
}

// {{.Function}} {{.Description}}
// @Router {{.Router}} [{{.Method}}]
func (s *{{.Name}}Controller) {{.Function}}(c *gin.Context) {
    // Define your controller function here
	s.ApiResponse(c, global.Success, "success", nil)
}