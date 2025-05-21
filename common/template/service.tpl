package {{.Package}}

import (
    "gin-base/common"
)

type {{.Name}}Service struct {
    common.BaseService
}

// {{.Function}} {{.Description}}
// @param YourParam string
// @return bool
func (s *{{.Name}}Service) {{.Function}}(YourParam string) bool {
    // Define your service function here
	return true
}