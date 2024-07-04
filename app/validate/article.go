package validate

import (
	"errors"
	Validator "github.com/gookit/validate"
)

// 文章列表请求验证
type ArticleListValidate struct {
	Page     int `form:"page" validate:"required|min:1" label:"页码"`
	PageSize int `form:"pageSize" validate:"required|min:1" label:"每页数量"`
}

// 获取文章列表请求验证
func GetArticleListValidate(data ArticleListValidate) error {
	v := Validator.Struct(data)
	// 或者使用
	// v := validates.New(u)
	if !v.Validate() {
		return errors.New(v.Errors.One())
	}

	return nil
}

// ConfigValidation 配置验证
// - 定义验证场景
// - 也可以添加验证设置
/*func (a ArticleListValidate) ConfigValidation(v *Validator.Validation) {
	// v.StringRule()

	v.WithScenes(Validator.SValues{
		//"list": []string{"Page", "PageSize"},
		//"create":    []string{"ExtInfo.Homepage", "Name", "Code"},
		//"update": []string{"ExtInfo.CityName", "Name"},
	})
}*/

// Messages 您可以自定义验证器错误消息
func (a ArticleListValidate) Messages() map[string]string {
	return Validator.MS{
		"required":     "字段 {field} 必填",
		"Page.min":     "页码最小为 1",
		"PageSize.min": "每页数量最小为 1",
	}
}

// Translates 你可以自定义字段翻译
func (a ArticleListValidate) Translates() map[string]string {
	return Validator.MS{
		"Page":     "页码",
		"PageSize": "每页数量",
	}
}
