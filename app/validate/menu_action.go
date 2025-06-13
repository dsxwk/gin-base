package validate

import (
	"errors"
	validator "github.com/gookit/validate"
)

// MenuAction 菜单功能请求验证
type MenuAction struct {
	ID        int64  `json:"id" validate:"required" label:"ID"`
	MenuID    int64  `json:"menu_id" validate:"required" label:"菜单ID"`
	Type      int64  `json:"type" validate:"required" label:"类型 1=header 2=operation"`
	Label     string `json:"label" validate:"required" label:"功能名称"`
	AuthValue string `json:"authValue" label:"权限标识"`
	IsLink    string `json:"isLink" validate:"required" label:"是否为链接 1=是 2=否"`
	Sort      string `json:"sort" validate:"required" label:"排序"`
}

// GetValidate 请求验证
func (s MenuAction) GetValidate(data MenuAction, scene string) error {
	v := validator.Struct(data, scene)
	if !v.Validate(scene) {
		return errors.New(v.Errors.One())
	}

	return nil
}

// ConfigValidation 配置验证
// - 定义验证场景
// - 也可以添加验证设置
func (s MenuAction) ConfigValidation(v *validator.Validation) {
	v.WithScenes(validator.SValues{
		"list":   []string{"MenuID"},
		"create": []string{"MenuID", "Type", "Label", "AuthValue", "IsLink", "Sort"},
		"update": []string{"ID", "MenuID", "Type", "Label", "AuthValue", "IsLink", "Sort"},
		"delete": []string{"ID"},
	})
}

// Messages 您可以自定义验证器错误消息
func (s MenuAction) Messages() map[string]string {
	return validator.MS{
		"required": "字段 {field} 必填",
		"int":      "字段 {field} 必须为整数",
	}
}

// Translates 你可以自定义字段翻译
func (s MenuAction) Translates() map[string]string {
	return validator.MS{
		"ID":        "ID",
		"MenuID":    "菜单ID",
		"Type":      "类型",
		"Label":     "功能名称",
		"AuthValue": "权限标识",
		"IsLink":    "是否为链接",
		"Sort":      "排序",
	}
}
