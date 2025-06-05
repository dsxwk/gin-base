package validate

import (
	"errors"
	validator "github.com/gookit/validate"
	"time"
)

// Cache 缓存验证
type Cache struct {
	Key    string        `json:"key" form:"key" validate:"required" label:"键"`
	Value  interface{}   `json:"value" validate:"required" label:"值"`
	Expire time.Duration `json:"expire" validate:"required" label:"过期时间"`
}

// GetValidate 请求验证
func (s Cache) GetValidate(data Cache, scene string) error {
	v := validator.Struct(data, scene)
	if !v.Validate(scene) {
		return errors.New(v.Errors.One())
	}

	return nil
}

// ConfigValidation 配置验证
// - 定义验证场景
// - 也可以添加验证设置
func (s Cache) ConfigValidation(v *validator.Validation) {
	v.WithScenes(validator.SValues{
		"setCache":    []string{"Key", "Value", "Expire"},
		"getCache":    []string{"Key"},
		"getAllCache": []string{},
		"deleteCache": []string{"Key"},
	})
}

// Messages 您可以自定义验证器错误消息
func (s Cache) Messages() map[string]string {
	return validator.MS{
		"required": "字段 {field} 必填",
	}
}

// Translates 你可以自定义字段翻译
func (s Cache) Translates() map[string]string {
	return validator.MS{
		"Key":    "键",
		"Value":  "值",
		"Expire": "过期时间",
	}
}
