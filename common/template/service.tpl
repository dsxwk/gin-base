package {{.Package}}

import (
    "github.com/gin-gonic/gin"
    "gin-base/common"
)

type {{.Name}}Service struct {
    common.BaseService
}

// @function: {{.Function}}
// @description: {{.Description}}
// @param: YourParam ParamType
// @return: bool
func (s *{{.Name}}Service) {{.Function}}(c *gin.Context) bool {
    // Define your service function here
	return true
}