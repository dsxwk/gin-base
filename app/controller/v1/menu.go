package v1

import (
	"fmt"
	"gin-base/app/model"
	"gin-base/app/service"
	"gin-base/app/validate"
	"gin-base/common"
	"gin-base/common/global"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"strconv"
)

type MenuController struct {
	common.BaseController
}

// List 列表
// @Router /api/v1/menu [get]
func (s *MenuController) List(c *gin.Context) {
	var (
		menuService service.MenuService
	)

	data, err := menuService.List()
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, data)
}

// Create 创建
// @Router /api/v1/menu [post]
func (s *MenuController) Create(c *gin.Context) {
	var (
		menuService  service.MenuService
		menuValidate validate.MenuValidate
		req          model.MenuQuery
	)

	err := c.ShouldBind(&req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}
	fmt.Printf("%+v", req)
	err = copier.Copy(&menuValidate, &req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}
	fmt.Printf("%+v", menuValidate)
	// 验证
	err = validate.GetMenuValidate(menuValidate, "create")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	data, err := menuService.Create(req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, data)
}

// Update 更新
// @Router /api/v1/menu/{id} [put]
func (s *MenuController) Update(c *gin.Context) {
	var (
		menuService  service.MenuService
		menuValidate validate.MenuValidate
		req          model.MenuQuery
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
	menuValidate.ID = id
	err = copier.Copy(&menuValidate, &req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.GetMenuValidate(menuValidate, "update")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	data, err := menuService.Update(req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "更新成功", data)
}

// Delete 删除
// @Router /api/v1/menu/{id} [delete]
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
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "删除成功", data)
}

// ActionList 功能列表
// @Router /api/v1/menu/{menu_id}/action [get]
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
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, data)
}

// ActionCreate 功能创建
// @Router /api/v1/menu/{id}/action [post]
func (s *MenuController) ActionCreate(c *gin.Context) {
	var (
		menuService        service.MenuService
		menuActionValidate validate.MenuActionValidate
		req                model.MenuActionQuery
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
		global.Log.Error(err.Error())
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
	err = validate.GetMenuActionValidate(menuActionValidate, "create")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	data, err := menuService.ActionCreate(req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
	}

	s.ApiResponse(c, global.Success, "创建成功", data)
}

// ActionUpdate 功能更新
// @Router /api/v1/menu/{id}/action/{action_id} [put]
func (s *MenuController) ActionUpdate(c *gin.Context) {
	var (
		menuService        service.MenuService
		menuActionValidate validate.MenuActionValidate
		req                model.MenuActionQuery
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

	IDParam := c.Param("action_id")
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
		global.Log.Error(err.Error())
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

	err = validate.GetMenuActionValidate(menuActionValidate, "update")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	data, err := menuService.ActionUpdate(req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "更新成功", data)
}

// ActionDelete 功能删除
// @Router /api/v1/menu/{id}/action/{action_id} [delete]
func (s *MenuController) ActionDelete(c *gin.Context) {
	var (
		menuService service.MenuService
	)

	iDParam := c.Param("action_id")
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
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "删除成功", data)
}
