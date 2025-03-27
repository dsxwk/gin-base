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
// @Param username body string true "用户名" SchemaExample("admin")
// @Param password body string true "密码" SchemaExample("123456")
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
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
