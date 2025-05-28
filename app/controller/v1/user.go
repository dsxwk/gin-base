package v1

import (
	"gin-base/app/model"
	"gin-base/app/service"
	"gin-base/app/validate"
	"gin-base/common"
	"gin-base/common/global"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"strconv"
)

type UserController struct {
	common.BaseController
}

// List 列表
// @Router /api/v1/user [get]
func (s *UserController) List(c *gin.Context) {
	var (
		userService service.UserService
		req         validate.UserValidate
		search      validate.UserSearch
		pageData    global.PageData
	)

	err := c.ShouldBindQuery(&req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 搜索
	err = c.ShouldBindQuery(&search)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.GetUserValidate(req, "list")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	err = copier.Copy(&pageData, &req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
	}

	pageData, err = userService.List(pageData, search)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, pageData)
}

// Create 创建
// @Router /api/v1/user [post]
func (s *UserController) Create(c *gin.Context) {
	var (
		userService  service.UserService
		userValidate validate.UserValidate
		req          model.User
	)

	err := c.ShouldBind(&req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	err = copier.Copy(&userValidate, &req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.GetUserValidate(userValidate, "create")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	userModel, err := userService.Create(req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, nil, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "创建成功", userModel)
}

// Update 更新
// @Router /api/v1/user/{id} [put]
func (s *UserController) Update(c *gin.Context) {
	var (
		userService  service.UserService
		userValidate validate.UserValidate
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
		global.Log.Error(err.Error())
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
	err = validate.GetUserValidate(userValidate, "update")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	userModel, err := userService.Update(req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "更新成功", userModel)
}

// Detail 详情
// @Router /api/v1/user/{id} [get]
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
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, user)
}

// Delete 删除
// @Router /v1/user/{id} [delete]
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
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "删除成功", userModel)
}
