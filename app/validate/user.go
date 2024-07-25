package validate

import (
	"errors"
	validator "github.com/gookit/validate"
)

// 用户请求验证
type UserValidate struct {
	ID       int64  `json:"id" validate:"required|int|gt:0" label:"ID"`
	Page     int    `form:"page" validate:"required|int|gt:0" label:"页码"`
	PageSize int    `form:"pageSize" validate:"required|int|gt:0" label:"每页数量"`
	Username string `json:"username" validate:"required" label:"用户名"`
	FullName string `json:"full_name" validate:"required" label:"姓名"`
	Nickname string `json:"nickname" validate:"required" label:"昵称"`
	Gender   int    `json:"gender" validate:"required|int" label:"性别"`
	Password string `json:"password" validate:"required" label:"密码"`
}

// 用户搜索
type UserSearch struct {
	Username string `form:"username" validate:"required" label:"用户名"`
	FullName string `form:"full_name" validate:"required" label:"姓名"`
	Nickname string `form:"nickname" validate:"required" label:"昵称"`
	Gender   int    `form:"gender" validate:"required|int" label:"性别"`
}

// 用户搜索验证
type UserSearchValidate struct {
	Username string `form:"username" validate:"required" label:"用户名"`
	FullName string `form:"full_name" validate:"required" label:"姓名"`
	Nickname string `form:"nickname" validate:"required" label:"昵称"`
	Gender   int    `form:"gender" validate:"required|int" label:"性别"`
}

// 请求验证
func GetUserValidate(data UserValidate, scene string) error {
	v := validator.Struct(data, scene)
	if !v.Validate(scene) {
		return errors.New(v.Errors.One())
	}

	return nil
}

// ConfigValidation 配置验证
// - 定义验证场景
// - 也可以添加验证设置
func (a UserValidate) ConfigValidation(v *validator.Validation) {
	v.WithScenes(validator.SValues{
		"list":   []string{"Page", "PageSize"},
		"detail": []string{"ID"},
		"create": []string{"Username", "FullName", "Nickname", "Gender", "Password"},
		"update": []string{"ID"},
		"delete": []string{"ID"},
	})
}

// Messages 您可以自定义验证器错误消息
func (s UserValidate) Messages() map[string]string {
	return validator.MS{
		"required":    "字段 {field} 必填",
		"int":         "字段 {field} 必须为整数",
		"Page.gt":     "字段 {field} 需大于 0",
		"PageSize.gt": "字段 {field} 需大于 0",
	}
}

// Translates 你可以自定义字段翻译
func (s UserValidate) Translates() map[string]string {
	return validator.MS{
		"Page":     "页码",
		"PageSize": "每页数量",
	}
}
