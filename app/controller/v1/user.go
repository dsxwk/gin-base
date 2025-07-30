package v1

import (
	"gin-base/app/model"
	"gin-base/app/service"
	"gin-base/app/validate"
	"gin-base/common/base"
	"gin-base/common/global"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"strconv"
)

type UserController struct {
	base.BaseController
}

// List 列表
// @Tags 用户管理
// @Summary 列表
// @Description 用户列表
// @Param token header string true "认证Token"
// @Param page query string true "页码"
// @Param pageSize query string true "分页大小"
// @Router /api/v1/user [get]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *UserController) List(c *gin.Context) {
	var (
		userService  service.UserService
		userValidate validate.User
		search       validate.UserSearch
		pageData     global.PageData
	)

	err := c.ShouldBindQuery(&userValidate)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 搜索
	err = c.ShouldBindQuery(&search)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.User{}.GetValidate(userValidate, "list")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	err = copier.Copy(&pageData, &userValidate)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
	}

	pageData, err = userService.List(pageData, search)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, pageData)
}

// Create 创建
// @Tags 用户管理
// @Summary 创建
// @Description 用户创建
// @Param token header string true "认证Token"
// @Param data body object true "创建参数" SchemaExample({"username":"用户名","fullName":"姓名","password":"密码"})
// @Router /api/v1/user [post]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *UserController) Create(c *gin.Context) {
	var (
		userService  service.UserService
		userValidate validate.User
		req          model.User
	)

	err := c.ShouldBind(&req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	err = copier.Copy(&userValidate, &req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.User{}.GetValidate(userValidate, "create")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	userModel, err := userService.Create(req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, nil, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "创建成功", userModel)
}

// Update 更新
// @Tags 用户管理
// @Summary 更新
// @Description 用户更新
// @Param token header string true "认证Token"
// @Param id path int true "用户ID"
// @Param data body object true "更新参数" SchemaExample({"username":"用户名","fullName":"姓名","password":"密码"})
// @Router /api/v1/user/{id} [put]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *UserController) Update(c *gin.Context) {
	var (
		userService  service.UserService
		userValidate validate.User
		req          model.User
	)

	idParam := c.Param("id")
	if idParam == "" {
		s.ApiResponse(c, global.ArgsError, "id参数必传")
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		s.ApiResponse(c, global.ArgsError, "id参数格式错误")
		return
	}

	err = c.ShouldBind(&req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	req.ID = id
	userValidate.ID = id
	err = copier.Copy(&userValidate, &req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.User{}.GetValidate(userValidate, "update")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	userModel, err := userService.Update(req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "更新成功", userModel)
}

// Detail 详情
// @Tags 用户管理
// @Summary 详情
// @Description 用户详情
// @Param token header string true "认证Token"
// @Param id path int true "用户ID"
// @Router /api/v1/user/{id} [get]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *UserController) Detail(c *gin.Context) {
	var (
		userService service.UserService
	)

	idParam := c.Param("id")
	if idParam == "" {
		s.ApiResponse(c, global.ArgsError, "id参数必传")
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		s.ApiResponse(c, global.ArgsError, "id参数格式错误")
		return
	}

	user, err := userService.Detail(id)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, user)
}

// Delete 删除
// @Tags 用户管理
// @Summary 删除
// @Description 用户删除
// @Param token header string true "认证Token"
// @Param id path int true "用户ID"
// @Router /v1/user/{id} [delete]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *UserController) Delete(c *gin.Context) {
	var (
		userService service.UserService
	)

	idParam := c.Param("id")
	if idParam == "" {
		s.ApiResponse(c, global.ArgsError, "id参数必传")
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		s.ApiResponse(c, global.ArgsError, "id参数格式错误")
		return
	}

	userModel, err := userService.Delete(id)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "删除成功", userModel)
}
