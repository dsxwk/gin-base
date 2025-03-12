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
