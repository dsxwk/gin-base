package v1

import (
	"gin-base/app/middleware"
	"gin-base/app/model"
	"gin-base/app/service"
	"gin-base/app/validate"
	"gin-base/common"
	"gin-base/common/global"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	common.BaseController
}

// Login 登录
// @Summary 登录
// @Description 用户登录
// @Tags 登录
// @Accept json
// @Produce json
// @Param user body validate.LoginValidate true "请求参数" SchemaExample({"username": "admin", "password": "123456"})
// @Success 200 {object} global.Response{global.Success} "成功返回"
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误"
// @Failure 500 {object} global.Response{global.SystemError} "系统错误"
// @Router /api/v1/login [post]
func (s *LoginController) Login(c *gin.Context) {
	var (
		loginService  = service.LoginService{}
		loginValidate validate.LoginValidate
		req           model.User
		jwt           middleware.Jwt
	)

	err := c.ShouldBind(&req)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	loginValidate.Username = req.Username
	loginValidate.Password = req.Password

	// 验证
	err = validate.GetLoginValidate(loginValidate, "delete")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	userModel, err := loginService.Login(req.Username, req.Password)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	token, exp, err := jwt.Encode(userModel.ID, 4*60*60)
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	res := make(map[string]interface{})
	res["token"] = token
	res["exp"] = exp

	s.ApiResponse(c, global.Success, "登录成功", res)
}
