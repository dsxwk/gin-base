package v1

import (
	"encoding/json"
	"gin-base/app/model"
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

// List
// @Tags 用户
// @Summary 列表
// @Router /v1/user [get]
func (s *UserController) List(c *gin.Context) {
	var (
		userService service.UserService
		req         validate.UserValidate
		search      validate.UserSearch
	)

	err := c.ShouldBindQuery(&req)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	// 搜索
	err = c.ShouldBindQuery(&search)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	// 验证
	err = validate.GetUserValidate(req, "list")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error(), nil)
		return
	}

	pageData, err := userService.List(req, search)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error(), nil)
		return
	}

	s.ApiResponse(c, global.Success, "获取成功", pageData)
}

// Create
// @Tags 用户
// @Summary 创建
// @Router /v1/user [post]
func (s *UserController) Create(c *gin.Context) {
	var (
		userService  service.UserService
		userValidate validate.UserValidate
		req          model.User
	)

	err := c.ShouldBind(&req)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	// 序列化为 JSON
	jsonData, err := json.Marshal(req)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	// 反序列化到响应结构体
	err = json.Unmarshal(jsonData, &userValidate)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	// 验证
	err = validate.GetUserValidate(userValidate, "create")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error(), nil)
		return
	}

	userModel, err := userService.Create(req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error(), nil)
		return
	}

	s.ApiResponse(c, global.Success, "创建成功", userModel)
}

// Update
// @Tags 用户
// @Summary 更新
// @Router /v1/user/:id [put]
func (s *UserController) Update(c *gin.Context) {
	var (
		userService  service.UserService
		userValidate validate.UserValidate
		req          model.User
	)

	idParam := c.Param("id")
	if idParam == "" {
		s.ApiResponse(c, global.ArgsError, "id参数必传", nil)
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		s.ApiResponse(c, global.ArgsError, "id参数格式错误", nil)
		return
	}

	err = c.ShouldBind(&req)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	req.ID = id
	userValidate.ID = id
	// 序列化为 JSON
	jsonData, err := json.Marshal(req)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	// 反序列化到响应结构体
	err = json.Unmarshal(jsonData, &userValidate)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	// 验证
	err = validate.GetUserValidate(userValidate, "update")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error(), nil)
		return
	}

	userModel, err := userService.Update(req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error(), nil)
		return
	}

	s.ApiResponse(c, global.Success, "更新成功", userModel)
}

// Detail
// @Tags 用户
// @Summary 详情
// @Router /v1/user/:id [get]
func (s *UserController) Detail(c *gin.Context) {
	var (
		userService service.UserService
		req         validate.UserValidate
	)

	idParam := c.Param("id")
	if idParam == "" {
		s.ApiResponse(c, global.ArgsError, "id参数必传", nil)
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		s.ApiResponse(c, global.ArgsError, "id参数格式错误", nil)
		return
	}

	req.ID = id

	// 验证
	err = validate.GetUserValidate(req, "detail")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error(), nil)
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

// Delete
// @Tags 用户
// @Summary 删除
// @Router /v1/user/:id [delete]
func (s *UserController) Delete(c *gin.Context) {
	var (
		userService  service.UserService
		userValidate validate.UserValidate
	)

	idParam := c.Param("id")
	if idParam == "" {
		s.ApiResponse(c, global.ArgsError, "id参数必传", nil)
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		s.ApiResponse(c, global.ArgsError, "id参数格式错误", nil)
		return
	}

	userValidate.ID = id

	// 验证
	err = validate.GetUserValidate(userValidate, "delete")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error(), nil)
		return
	}

	userModel, err := userService.Delete(id)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error(), nil)
		return
	}

	s.ApiResponse(c, global.Success, "删除成功", userModel)
}
