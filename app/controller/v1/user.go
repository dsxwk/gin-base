package v1

import (
	"gin-base/app/service"
	"gin-base/app/validate"
	"gin-base/common"
	"gin-base/common/global"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	common.BaseController
}

// @Tags    用户
// @Summary 列表
// @Router  /v1/user [get]
func (s *UserController) List(c *gin.Context) {
	var (
		userService service.UserService
		req         validate.UserValidate
	)

	err := c.ShouldBindQuery(&req)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	// 验证
	err = validate.GetUserValidate(req, "list")
	if err != nil {
		s.ApiResponse(c, global.Error, err.Error(), nil)
		return
	}

	pageData, err := userService.List(req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error(), nil)
		return
	}

	s.ApiResponse(c, global.Success, "获取成功", pageData)
}
