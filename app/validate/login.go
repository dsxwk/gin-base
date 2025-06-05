package validate

import (
	"errors"
	validator "github.com/gookit/validate"
)

// LoginValidate YourDesc
type LoginValidate struct {
	CaptchaID string `json:"captchaId" validate:"required" label:"验证码ID"`
	Code      string `json:"code" validate:"required" label:"验证码"`
	Username  string `json:"username" validate:"required" label:"用户名"`
	Password  string `json:"password" validate:"required" label:"密码"`
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
		"login":  []string{"Username", "Password", "CaptchaID", "Code"},
		"verify": []string{"CaptchaID", "Code"},
	})
}

// Messages 您可以自定义验证器错误消息
func (s LoginValidate) Messages() map[string]string {
	return validator.MS{
		"required": "字段 {field} 必填",
	}
}

// Translates 你可以自定义字段翻译
func (s LoginValidate) Translates() map[string]string {
	return validator.MS{
		"CaptchaID": "验证码ID",
		"Code":      "验证码",
	}
}
