package validate

import (
	"errors"
	validator "github.com/gookit/validate"
)

// MenuActionValidate 菜单功能请求验证
type MenuActionValidate struct {
	ID     int64  `json:"id" validate:"required" label:"ID"`
	MenuID int64  `json:"menu_id" validate:"required" label:"菜单ID"`
	Type   int64  `json:"type" validate:"required" label:"类型 1=header 2=operation"`
	Name   string `json:"name" validate:"required" label:"功能名称"`
	IsLink string `json:"isLink" validate:"required" label:"是否为链接 1=是 2=否"`
	Sort   string `json:"sort" validate:"required" label:"排序"`
}

// GetMenuActionValidate 请求验证
func GetMenuActionValidate(data MenuActionValidate, scene string) error {
	v := validator.Struct(data, scene)
	if !v.Validate(scene) {
		return errors.New(v.Errors.One())
	}

	return nil
}

// ConfigValidation 配置验证
// - 定义验证场景
// - 也可以添加验证设置
func (s MenuActionValidate) ConfigValidation(v *validator.Validation) {
	v.WithScenes(validator.SValues{
		"list":   []string{"MenuID"},
		"create": []string{"MenuID", "Type", "Name", "IsLink", "Sort"},
		"update": []string{"ID", "MenuID", "Type", "Name", "IsLink", "Sort"},
		"delete": []string{"ID"},
	})
}

// Messages 您可以自定义验证器错误消息
func (s MenuActionValidate) Messages() map[string]string {
	return validator.MS{
		"required": "字段 {field} 必填",
		"int":      "字段 {field} 必须为整数",
	}
}

// Translates 你可以自定义字段翻译
func (s MenuActionValidate) Translates() map[string]string {
	return validator.MS{
		"ID":     "ID",
		"MenuID": "菜单ID",
		"Type":   "类型",
		"Name":   "功能名称",
		"IsLink": "是否为链接",
		"Sort":   "排序",
	}
}
