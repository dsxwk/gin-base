package validate

import (
	"errors"
	validator "github.com/gookit/validate"
)

// LoginValidate 登录请求验证
type LoginValidate struct {
	Username string `json:"username" validate:"required|minLen:3|maxLen:10" label:"用户名"`
	Password string `json:"password" validate:"required|minLen:6" label:"密码"`
}

// GetLoginValidate 请求验证
func GetLoginValidate(data LoginValidate, scene string) error {
	v := validator.Struct(data, scene)
	if !v.Validate(scene) {
		return errors.New(v.Errors.One())
	}

	return nil
}

// ConfigValidation 配置验证
// - 定义验证场景
// - 也可以添加验证设置
func (s LoginValidate) ConfigValidation(v *validator.Validation) {
	v.WithScenes(validator.SValues{
		"login": []string{"Username", "Password"},
	})
}

// Messages 您可以自定义验证器错误消息
func (s LoginValidate) Messages() map[string]string {
	return validator.MS{
		"required":        "字段 {field} 必填",
		"int":             "字段 {field} 必须为整数",
		"Username.minLen": "字段 {field} 最小长度为 3",
		"Username.maxLen": "字段 {field} 最大长度为 10",
		"Password.minLen": "字段 {field} 最小长度为 6",
	}
}

// Translates 你可以自定义字段翻译
func (s LoginValidate) Translates() map[string]string {
	return validator.MS{
		"Page":     "页码",
		"PageSize": "每页数量",
	}
}
