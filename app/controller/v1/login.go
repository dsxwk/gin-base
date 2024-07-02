package v1

import (
	"fmt"
	"gin-base/app/service"
	"gin-base/common"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	common.BaseController
}

// 登录
func (this *LoginController) Login(c *gin.Context) {
	var (
		loginService = service.LoginService{}
	)

	err := loginService.Login(0, 0)
	if err != nil {
		fmt.Printf("登录失败：%v", err)
	}

	res := make(map[string]interface{})
	res["token"] = "token"
	res["exp"] = 60 * 60 * 2
	this.ApiResponse(c, 200, "登录成功", res)
}
