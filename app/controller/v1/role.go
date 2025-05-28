package v1

import (
	"gin-base/app/service"
	"gin-base/app/validate"
	"gin-base/common"
	"gin-base/common/global"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type RoleController struct {
	common.BaseController
}

// List 列表
// @Router /api/v1/role [get]
func (s *RoleController) List(c *gin.Context) {
	var (
		roleService service.RoleService
		req         validate.RoleValidate
		search      validate.RoleSearchValidate
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
	err = validate.GetRoleValidate(req, "list")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	err = copier.Copy(&pageData, &req)
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
