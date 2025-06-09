package {{.Package}}

import (
    "gin-base/common/base"
)

type {{.Name}}Service struct {
    base.BaseService
}

// {{.Function}} {{.Description}}
// @param YourParam string
// @return bool
func (s *{{.Name}}Service) {{.Function}}(YourParam string) bool {
    // Define your service function here
	return true
}