package v1

import (
	"gin-base/app/service"
	"gin-base/app/validate"
	"gin-base/common"
	"gin-base/common/global"
	"github.com/gin-gonic/gin"
	"strconv"
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

// @Tags    用户
// @Summary 详情
// @Router  /v1/user/:id [get]
func (s *UserController) Detail(c *gin.Context) {
	var (
		userService service.UserService
		req         validate.UserValidate
	)

	idParam := c.Param("id")
	if idParam == "" {
		s.ApiResponse(c, global.Error, "id参数必传", nil)
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		s.ApiResponse(c, global.Error, "id参数格式错误", nil)
		return
	}

	req.ID = id

	// 验证
	err = validate.GetUserValidate(req, "detail")
	if err != nil {
		s.ApiResponse(c, global.Error, err.Error(), nil)
		return
	}

	pageData, err := userService.Detail(id)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error(), nil)
		return
	}

	s.ApiResponse(c, global.Success, "获取成功", pageData)
}
