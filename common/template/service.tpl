package {{.Package}}

import (
    "gin-base/common/base"
)

type {{.Name}}Service struct {
    base.BaseService
}

// {{.Function}} {{.Description}}
func (s *{{.Name}}Service) {{.Function}}() {
    // todo
}