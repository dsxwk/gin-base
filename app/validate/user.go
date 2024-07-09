package validate

import (
	"errors"
	validator "github.com/gookit/validate"
)

// 用户请求验证
type UserValidate struct {
	Page     int `form:"page" validate:"required|min:1" label:"页码"`
	PageSize int `form:"pageSize" validate:"required|min:1" label:"每页数量"`
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
		"list": []string{"Page", "PageSize"},
	})
}

// Messages 您可以自定义验证器错误消息
func (s UserValidate) Messages() map[string]string {
	return validator.MS{
		"required":     "字段 {field} 必填",
		"Page.min":     "页码最小为 1",
		"PageSize.min": "每页数量最小为 1",
	}
}

// Translates 你可以自定义字段翻译
func (s UserValidate) Translates() map[string]string {
	return validator.MS{
		"Page":     "页码",
		"PageSize": "每页数量",
	}
}
