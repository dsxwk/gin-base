package validate

import (
	"errors"
	validator "github.com/gookit/validate"
)

// MenuValidate 菜单请求验证
type MenuValidate struct {
	ID             int64  `json:"id" validate:"required" label:"ID"`
	PID            int64  `json:"pid" validate:"required" label:"父级ID"`
	Title          string `json:"title" validate:"required" label:"菜单名称"`
	Name           string `json:"name" validate:"required" label:"路由名称"`
	Path           string `json:"path" validate:"required" label:"路由路径"`
	Redirect       string `json:"redirect" validate:"required" label:"重定向"`
	Icon           string `json:"icon" validate:"required" label:"菜单图标"`
	ComponentAlias string `json:"component_alias" validate:"required" label:"组件路径"`
	IsLink         string `json:"is_link" validate:"required" label:"链接地址"`
	IsHide         string `json:"is_hide" validate:"required" label:"是否隐藏"`
	Status         string `json:"status" validate:"required" label:"状态"`
	Sort           string `json:"sort" validate:"required" label:"排序"`
}

// GetMenuValidate 请求验证
func GetMenuValidate(data MenuValidate, scene string) error {
	v := validator.Struct(data, scene)
	if !v.Validate(scene) {
		return errors.New(v.Errors.One())
	}

	return nil
}

// ConfigValidation 配置验证
// - 定义验证场景
// - 也可以添加验证设置
func (s MenuValidate) ConfigValidation(v *validator.Validation) {
	v.WithScenes(validator.SValues{
		"list":   []string{},
		"create": []string{"Title", "Name", "Path"},
		"update": []string{"ID", "Title", "Name", "Path"},
		"delete": []string{"ID"},
	})
}

// Messages 您可以自定义验证器错误消息
func (s MenuValidate) Messages() map[string]string {
	return validator.MS{
		"required": "字段 {field} 必填",
		"int":      "字段 {field} 必须为整数",
	}
}

// Translates 你可以自定义字段翻译
func (s MenuValidate) Translates() map[string]string {
	return validator.MS{
		"ID":    "ID",
		"PID":   "父级ID",
		"Title": "菜单名称",
		"Name":  "路由名称",
	}
}
