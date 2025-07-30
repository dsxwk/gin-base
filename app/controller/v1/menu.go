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

type MenuController struct {
	base.BaseController
}

// List 列表
// @Tags 菜单管理
// @Summary 列表
// @Description 菜单列表
// @Param token header string true "认证Token"
// @Router /api/v1/menu [get]
// @Param data body object true "更新参数" SchemaExample({"title":"Go语言","content":"内容"})
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *MenuController) List(c *gin.Context) {
	var (
		menuService service.MenuService
	)

	data, err := menuService.List()
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, data)
}

// RoleMenu 角色菜单
// @Tags 菜单管理
// @Summary 角色菜单
// @Description 角色菜单
// @Param token header string true "认证Token"
// @Param roleIds query string true "角色ids"
// @Router /api/v1/menu/role-menu [get]
// @Param data body object true "更新参数" SchemaExample({"title":"Go语言","content":"内容"})
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *MenuController) RoleMenu(c *gin.Context) {
	var (
		menuService service.MenuService
		search      validate.MenuSearch
	)

	err := c.ShouldBindQuery(&search)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	data, err := menuService.RoleMenu(search.RoleIds)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, data)
}

// Create 创建
// @Tags 菜单管理
// @Summary 创建
// @Description 菜单创建
// @Param token header string true "认证Token"
// @Param data body object true "创建参数" SchemaExample({"pid":0,"name":"测试菜单","path":"/test","meta":{}})
// @Router /api/v1/menu [post]
// @Param data body object true "更新参数" SchemaExample({"title":"Go语言","content":"内容"})
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *MenuController) Create(c *gin.Context) {
	var (
		menuService  service.MenuService
		menuValidate validate.Menu
		req          model.Menu
	)

	err := c.ShouldBind(&req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	err = copier.CopyWithOption(&menuValidate, &req, copier.Option{DeepCopy: true})
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.Menu{}.GetValidate(menuValidate, "create")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	data, err := menuService.Create(req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, data)
}

// Update 更新
// @Tags 菜单管理
// @Summary 更新
// @Description 菜单更新
// @Param token header string true "认证Token"
// @Param id path int true "菜单ID"
// @Param data body object true "更新参数" SchemaExample({"pid":0,"name":"测试菜单","path":"/test","meta":{}})
// @Router /api/v1/menu/{id} [put]
// @Param data body object true "更新参数" SchemaExample({"title":"Go语言","content":"内容"})
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *MenuController) Update(c *gin.Context) {
	var (
		menuService  service.MenuService
		menuValidate validate.Menu
		req          model.Menu
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
	menuValidate.ID = id

	err = copier.CopyWithOption(&menuValidate, &req, copier.Option{DeepCopy: true})
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.Menu{}.GetValidate(menuValidate, "update")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	data, err := menuService.Update(req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "更新成功", data)
}

// Detail 详情
// @Tags 菜单管理
// @Summary 详情
// @Description 菜单详情
// @Param token header string true "认证Token"
// @Param id path int true "菜单ID"
// @Router /api/v1/menu/{id} [get]
// @Param data body object true "更新参数" SchemaExample({"title":"Go语言","content":"内容"})
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *MenuController) Detail(c *gin.Context) {
	var (
		menuService service.MenuService
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

	data, err := menuService.Detail(id)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, data)
}

// Delete 删除
// @Tags 菜单管理
// @Summary 删除
// @Description 菜单删除
// @Param token header string true "认证Token"
// @Param id path int true "菜单ID"
// @Router /api/v1/menu/{id} [delete]
// @Param data body object true "更新参数" SchemaExample({"title":"Go语言","content":"内容"})
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *MenuController) Delete(c *gin.Context) {
	var (
		menuService service.MenuService
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

	data, err := menuService.Delete(id)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "删除成功", data)
}

// ActionList 功能列表
// @Tags 菜单管理
// @Summary 功能列表
// @Description 菜单功能列表
// @Param token header string true "认证Token"
// @Param id path int true "菜单ID"
// @Router /api/v1/menu/{id}/action [get]
// @Param data body object true "更新参数" SchemaExample({"title":"Go语言","content":"内容"})
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *MenuController) ActionList(c *gin.Context) {
	var (
		menuService service.MenuService
	)

	menuIDParam := c.Param("id")
	if menuIDParam == "" {
		s.ApiResponse(c, global.ArgsError, "菜单id参数必传")
		return
	}

	menuID, err := strconv.ParseInt(menuIDParam, 10, 64)
	if err != nil {
		s.ApiResponse(c, global.ArgsError, "id参数格式错误")
		return
	}

	data, err := menuService.ActionList(menuID)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, data)
}

// ActionCreate 功能创建
// @Tags 菜单管理
// @Summary 功能创建
// @Description 菜单功能创建
// @Param token header string true "认证Token"
// @Param id path int true "菜单ID"
// @Param data body object true "创建参数" SchemaExample({"menuId":1,"type":1,"label":"功能名称","authValue":"btn.add"})
// @Router /api/v1/menu/{id}/action [post]
// @Param data body object true "更新参数" SchemaExample({"title":"Go语言","content":"内容"})
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *MenuController) ActionCreate(c *gin.Context) {
	var (
		menuService        service.MenuService
		menuActionValidate validate.MenuAction
		req                model.MenuAction
	)

	menuIDParam := c.Param("id")
	if menuIDParam == "" {
		s.ApiResponse(c, global.ArgsError, "菜单id参数必传")
		return
	}

	menuID, err := strconv.ParseInt(menuIDParam, 10, 64)
	if err != nil {
		s.ApiResponse(c, global.ArgsError, "菜单id参数格式错误")
		return
	}

	err = c.ShouldBind(&req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	req.MenuID = menuID
	err = copier.Copy(&menuActionValidate, &req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.MenuAction{}.GetValidate(menuActionValidate, "create")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	data, err := menuService.ActionCreate(req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
	}

	s.ApiResponse(c, global.Success, "创建成功", data)
}

// ActionUpdate 功能更新
// @Tags 菜单管理
// @Summary 功能创建
// @Description 菜单功能创建
// @Param token header string true "认证Token"
// @Param id path int true "菜单ID"
// @Param actionId path int true "功能ID"
// @Param data body object true "更新参数" SchemaExample({"menuId":1,"type":1,"label":"功能名称","authValue":"btn.add"})
// @Router /api/v1/menu/{id}/action/{actionId} [put]
// @Param data body object true "更新参数" SchemaExample({"title":"Go语言","content":"内容"})
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *MenuController) ActionUpdate(c *gin.Context) {
	var (
		menuService        service.MenuService
		menuActionValidate validate.MenuAction
		req                model.MenuAction
	)

	menuIDParam := c.Param("id")
	if menuIDParam == "" {
		s.ApiResponse(c, global.ArgsError, "菜单id参数必传")
		return
	}

	menuID, err := strconv.ParseInt(menuIDParam, 10, 64)
	if err != nil {
		s.ApiResponse(c, global.ArgsError, "菜单id参数格式错误")
		return
	}

	IDParam := c.Param("actionId")
	if IDParam == "" {
		s.ApiResponse(c, global.ArgsError, "功能id参数必传")
		return
	}

	ID, err := strconv.ParseInt(IDParam, 10, 64)
	if err != nil {
		s.ApiResponse(c, global.ArgsError, "功能id参数格式错误")
		return
	}

	err = c.ShouldBind(&req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	req.ID = ID
	req.MenuID = menuID
	err = copier.Copy(&menuActionValidate, &req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	err = validate.MenuAction{}.GetValidate(menuActionValidate, "update")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	data, err := menuService.ActionUpdate(req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "更新成功", data)
}

// ActionDelete 功能删除
// @Tags 菜单管理
// @Summary 功能删除
// @Description 菜单功能删除
// @Param token header string true "认证Token"
// @Param id path int true "菜单ID"
// @Param actionId path int true "功能ID"
// @Router /api/v1/menu/{id}/action/{actionId} [delete]
// @Param data body object true "更新参数" SchemaExample({"title":"Go语言","content":"内容"})
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *MenuController) ActionDelete(c *gin.Context) {
	var (
		menuService service.MenuService
	)

	iDParam := c.Param("actionId")
	if iDParam == "" {
		s.ApiResponse(c, global.ArgsError, "功能id参数必传")
		return
	}

	ID, err := strconv.ParseInt(iDParam, 10, 64)
	if err != nil {
		s.ApiResponse(c, global.ArgsError, "功能id参数格式错误")
		return
	}

	menuIDParam := c.Param("id")
	if menuIDParam == "" {
		s.ApiResponse(c, global.ArgsError, "菜单id参数必传")
		return
	}

	menuID, err := strconv.ParseInt(menuIDParam, 10, 64)
	if err != nil {
		s.ApiResponse(c, global.ArgsError, "菜单id参数格式错误")
		return
	}

	data, err := menuService.ActionDelete(ID, menuID)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "删除成功", data)
}

// ActionDetail 功能详情
// @Tags 菜单管理
// @Summary 功能详情
// @Description 菜单功能详情
// @Param token header string true "认证Token"
// @Param id path int true "菜单ID"
// @Param actionId path int true "功能ID"
// @Router /api/v1/menu/{id}/action/{actionId} [get]
// @Param data body object true "更新参数" SchemaExample({"title":"Go语言","content":"内容"})
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *MenuController) ActionDetail(c *gin.Context) {
	var (
		menuService service.MenuService
	)

	idParam := c.Param("actionId")
	if idParam == "" {
		s.ApiResponse(c, global.ArgsError, "功能id参数必传")
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		s.ApiResponse(c, global.ArgsError, "id参数格式错误")
		return
	}

	data, err := menuService.ActionDetail(id)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, data)
}
