package validate

import (
	"errors"
	validator "github.com/gookit/validate"
)

// RoleValidate Your Description
type RoleValidate struct {
	Page     int    `form:"page" validate:"required|int|gt:0" label:"页码"`
	PageSize int    `form:"pageSize" validate:"required|int|gt:0" label:"每页数量"`
	IsPage   *bool  `form:"isPage" validate:"required|bool" label:"是否分页"`
	ID       int64  `json:"id" validate:"required" label:"ID"`
	Name     string `json:"name" validate:"required" label:"角色名称"`
	Desc     string `json:"desc" validate:"required" label:"描述"`
}

// RoleSearchValidate 角色搜索验证
type RoleSearchValidate struct {
	Name   string `form:"name" validate:"required" label:"角色名称"`
	Status int64  `form:"status" validate:"required" label:"角色状态 1=启用 2=禁用"`
}

// GetRoleValidate 请求验证
func GetRoleValidate(data RoleValidate, scene string) error {
	v := validator.Struct(data, scene)
	if !v.Validate(scene) {
		return errors.New(v.Errors.One())
	}

	return nil
}

// ConfigValidation 配置验证
// - 定义验证场景
// - 也可以添加验证设置
func (s RoleValidate) ConfigValidation(v *validator.Validation) {
	v.WithScenes(validator.SValues{
		"list":   []string{"Page", "PageSize"},
		"create": []string{"Name"},
		"update": []string{"ID", "Name"},
		"detail": []string{"ID"},
		"delete": []string{"ID"},
	})
}

// Messages 您可以自定义验证器错误消息
func (s RoleValidate) Messages() map[string]string {
	return validator.MS{
		"required":    "字段 {field} 必填",
		"int":         "字段 {field} 必须为整数",
		"Page.gt":     "字段 {field} 需大于 0",
		"PageSize.gt": "字段 {field} 需大于 0",
	}
}

// Translates 你可以自定义字段翻译
func (s RoleValidate) Translates() map[string]string {
	return validator.MS{
		"Page":     "页码",
		"PageSize": "每页数量",
		"ID":       "ID",
		"Name":     "角色名称",
	}
}
