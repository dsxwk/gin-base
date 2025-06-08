package validate

import (
	"errors"
	validator "github.com/gookit/validate"
)

// Dict 字典
type Dict struct {
	ID     int64  `json:"id" validate:"required" label:"ID"`
	PID    int64  `json:"pid" validate:"required" label:"父级ID"`
	Name   string `json:"name" form:"name" validate:"required" label:"字段名称(英文)"`
	Title  string `json:"title" validate:"required" label:"字段名称(中文)"`
	Label  string `json:"label" validate:"required" label:"映射值"`
	Status int64  `json:"status" validate:"required" label:"状态 1=启用 2=停用"`
}

// GetValidate 请求验证
func (s Dict) GetValidate(data Dict, scene string) error {
	v := validator.Struct(data, scene)
	if !v.Validate(scene) {
		return errors.New(v.Errors.One())
	}

	return nil
}

// ConfigValidation 配置验证
// - 定义验证场景
// - 也可以添加验证设置
func (s Dict) ConfigValidation(v *validator.Validation) {
	v.WithScenes(validator.SValues{
		"get":    []string{"Name"},
		"list":   []string{},
		"create": []string{"Name", "Title", "Status"},
		"update": []string{"ID", "Name", "Title", "Status"},
		"detail": []string{"ID"},
		"delete": []string{"ID"},
	})
}

// Messages 您可以自定义验证器错误消息
func (s Dict) Messages() map[string]string {
	return validator.MS{
		"required":    "字段 {field} 必填",
		"int":         "字段 {field} 必须为整数",
		"Page.gt":     "字段 {field} 需大于 0",
		"PageSize.gt": "字段 {field} 需大于 0",
	}
}

// Translates 你可以自定义字段翻译
func (s Dict) Translates() map[string]string {
	return validator.MS{
		"Page":     "页码",
		"PageSize": "每页数量",
		"ID":       "ID",
		"Title":    "标题",
		"Content":  "内容",
	}
}
