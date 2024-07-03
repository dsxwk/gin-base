package v1

import (
	"gin-base/app/service"
	"gin-base/app/validate"
	"gin-base/common"
	"gin-base/common/global"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	common.BaseController
}

// @Tags    用户
// @Summary 列表
// @Router  /v1/user [get]
func (this *UserController) List(c *gin.Context) {
	var (
		userService service.LoginService
		requestData validate.ListValidate
	)

	err := c.ShouldBindQuery(&requestData)
	if err != nil {
		// 验证1
		msg := requestData.GetError(err.(validator.ValidationErrors))
		this.ApiResponse(c, global.Error, msg, nil)
		return
	}

	pageData, err := userService.List(requestData)
	if err != nil {
		global.Log.Error(err.Error())
		this.ApiResponse(c, global.SystemError, err.Error(), nil)
		return
	}

	this.ApiResponse(c, global.Success, "获取成功", pageData)
}
