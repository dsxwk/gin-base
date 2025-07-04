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

type RoleController struct {
	base.BaseController
}

// List 列表
// @Tags 角色管理
// @Summary 列表
// @Description 角色列表
// @Param token header string true "认证Token"
// @Param page query string true "页码"
// @Param pageSize query string true "分页大小"
// @Router /api/v1/role [get]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *RoleController) List(c *gin.Context) {
	var (
		roleService  service.RoleService
		roleValidate validate.Role
		search       validate.RoleSearch
		pageData     global.PageData
	)

	err := c.ShouldBindQuery(&roleValidate)
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
	err = validate.Role{}.GetValidate(roleValidate, "list")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	err = copier.Copy(&pageData, &roleValidate)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
	}

	data, err := roleService.List(pageData, search)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, data)
}

// Create 创建
// @Tags 角色管理
// @Summary 创建
// @Description 角色创建
// @Param token header string true "认证Token"
// @Param data body object true "创建参数" SchemaExample({"name":"角色名称","desc":"描述","status":1})
// @Router /api/v1/role [post]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *RoleController) Create(c *gin.Context) {
	var (
		roleService  service.RoleService
		roleValidate validate.Role
		req          model.Roles
	)

	err := c.ShouldBind(&req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	err = copier.Copy(&roleValidate, &req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.Role{}.GetValidate(roleValidate, "create")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	roleModel, err := roleService.Create(req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, nil, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "创建成功", roleModel)
}

// Update 更新
// @Tags 角色管理
// @Summary 更新
// @Description 角色更新
// @Param token header string true "认证Token"
// @Param id path int true "角色ID"
// @Param data body object true "更新参数" SchemaExample({"name":"角色名称","desc":"描述","status":1})
// @Router /api/v1/role/{id} [put]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *RoleController) Update(c *gin.Context) {
	var (
		roleService  service.RoleService
		roleValidate validate.Role
		req          model.Roles
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
	roleValidate.ID = id
	err = copier.Copy(&roleValidate, &req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.Role{}.GetValidate(roleValidate, "update")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	roleModel, err := roleService.Update(req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "更新成功", roleModel)
}

// Detail 详情
// @Tags 角色管理
// @Summary 详情
// @Description 角色详情
// @Param token header string true "认证Token"
// @Param id path int true "角色ID"
// @Router /api/v1/role/{id} [get]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *RoleController) Detail(c *gin.Context) {
	var (
		roleService service.RoleService
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

	role, err := roleService.Detail(id)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, role)
}

// Delete 删除
// @Tags 角色管理
// @Summary 删除
// @Description 角色删除
// @Param token header string true "认证Token"
// @Param id path int true "角色ID"
// @Router /v1/role/{id} [delete]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *RoleController) Delete(c *gin.Context) {
	var (
		roleService service.RoleService
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

	roleModel, err := roleService.Delete(id)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "删除成功", roleModel)
}
