package validate

import (
	"github.com/go-playground/validator/v10"
)

// 列表请求验证
type UserListRequest struct {
	Page     int `form:"page" binding:"required" comment:"页码"`
	PageSize int `form:"pageSize" binding:"required" comment:"每页数量"`
}

// 错误提示
func (this *UserListRequest) GetError(errs validator.ValidationErrors) string {
	for _, e := range errs {
		if e.Field() == "Page" {
			switch e.Tag() {
			case "required":
				return "分页page必传"
			}
		}
		if e.Field() == "PageSize" {
			switch e.Tag() {
			case "required":
				return "每页条数pageSize必传"
			}
		}
	}
	return "参数错误"
}
