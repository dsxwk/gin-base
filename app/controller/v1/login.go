package v1

import (
	"gin-base/app/middleware"
	"gin-base/app/service"
	"gin-base/app/validate"
	"gin-base/common/base"
	"gin-base/common/global"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

type LoginController struct {
	base.BaseController
}

// Login 登录
// @Tags 登录相关
// @Summary 登录
// @Description 用户登录
// @Accept json
// @Produce json
// @Param data body validate.Login true "登录参数"
// @Router /api/v1/login [post]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *LoginController) Login(c *gin.Context) {
	var (
		loginService  service.LoginService
		loginValidate validate.Login
		jwt           middleware.Jwt
	)

	err := c.ShouldBind(&loginValidate)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.Login{}.GetValidate(loginValidate, "login")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	// 验证码校验
	b := s.verify(loginValidate.CaptchaID, loginValidate.Code)
	if !b {
		s.ApiResponse(c, global.ArgsError, "验证码错误")
		return
	}

	userModel, err := loginService.Login(loginValidate.Username, loginValidate.Password)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	token, rToken, tExpire, rExpire, err := jwt.WithRefresh(userModel.ID, 2*60*60, 2*24*60*60)
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "登录成功", map[string]interface{}{
		"token": map[string]interface{}{
			"accessToken":   token,
			"refreshToken":  rToken,
			"expire":        tExpire,
			"refreshExpire": rExpire,
		},
		"user": userModel,
	})
}

// RefreshToken 刷新token
// @Tags 登录相关
// @Summary 刷新token
// @Description 刷新token
// @Accept json
// @Produce json
// @Param token header string true "刷新Token"
// @Router /api/v1/refresh-token [post]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *LoginController) RefreshToken(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" || token == "null" {
		global.Log.Error(global.ArgsError.Message)
		s.ApiResponse(c, global.ArgsError)
		return
	}

	j := middleware.Jwt{}
	claims, err := j.Decode(token)
	if err != nil || claims["typ"] != "refresh" {
		global.ApiResponse(c, global.Unauthorized, "无效的refresh token")
		return
	}

	uid := int64(claims["id"].(float64))
	aToken, rToken, tExpire, rExpire, err := j.WithRefresh(uid, 2*60*60, 2*24*60*60)
	if err != nil {
		global.ApiResponse(c, global.SystemError)
		return
	}

	s.ApiResponse(c, global.Success, map[string]interface{}{
		"accessToken":   aToken,
		"refreshToken":  rToken,
		"expire":        tExpire,
		"refreshExpire": rExpire,
	})
}

// generateCharset 生成字符串
func (s *LoginController) generateCharset() string {
	charset := ""

	// 添加数字 0-9
	for i := '0'; i <= '9'; i++ {
		charset += string(i)
	}

	// 添加小写字母 a-z
	for i := 'a'; i <= 'z'; i++ {
		charset += string(i)
	}

	// 添加大写字母 A-Z
	for i := 'A'; i <= 'Z'; i++ {
		charset += string(i)
	}

	return charset
}

// GetCaptcha 获取验证码
// @Tags 登录相关
// @Summary 获取验证码
// @Description 获取验证码
// @Accept json
// @Produce json
// @Router /api/v1/captcha [get]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
func (s *LoginController) GetCaptcha(c *gin.Context) {
	// 配置验证码
	driver := base64Captcha.NewDriverString(
		60,                  // 高度
		240,                 // 宽度
		0,                   // 噪点数量
		0,                   // base64Captcha.OptionShowHollowLine|base64Captcha.OptionShowSlimeLine, // 显示线条选项
		4,                   // 验证码长度
		s.generateCharset(), // 验证码字符集
		&color.RGBA{R: 255, G: 255, B: 255, A: 255}, // 背景颜色（白色）
		base64Captcha.DefaultEmbeddedFonts,          // 字体存储
		nil,                                         // []string{"wqy-microfiche.ttf"},              // 字体名称
	)

	// 生成验证码
	id, b64s, _, err := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore).Generate()
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	// 返回验证码 ID 和图片的 Base64 数据
	s.ApiResponse(c, global.Success, map[string]interface{}{
		"captchaId":    id,
		"captchaImage": b64s,
	})
}

// CheckCaptcha 校验验证码
// @Tags 登录相关
// @Summary 校验验证码
// @Description 校验验证码
// @Router /api/v1/captcha [post]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
func (s *LoginController) CheckCaptcha(c *gin.Context) {
	var (
		loginValidate validate.Login
	)

	err := c.ShouldBind(&loginValidate)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.Login{}.GetValidate(loginValidate, "verify")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, s.verify(loginValidate.CaptchaID, loginValidate.Code))
}

// verify 验证
// @return bool
func (s *LoginController) verify(captchaId, value string) bool {
	return base64Captcha.DefaultMemStore.Verify(captchaId, value, true)
}
