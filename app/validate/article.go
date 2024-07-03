package validate

import (
	"errors"
	"fmt"
	validates "github.com/gookit/validate"
)

// 文章列表请求验证
type ArticleListValidate struct {
	Page     int `form:"page" validate:"required|get=0" label:"页码"`
	PageSize int `form:"pageSize" validate:"required|get=0" label:"每页数量"`
}

// 获取文章列表请求验证
func GetArticleListValidate(data ArticleListValidate) error {
	v := validates.Struct(data)
	// 或者使用
	// v := validates.New(u)
	if !v.Validate() {
		fmt.Printf("验证失败: %v\n", v.Errors)
		return errors.New(v.Errors.One())
	} else {
		fmt.Printf("验证通过: %v\n", v.Errors)
	}

	return nil
}
