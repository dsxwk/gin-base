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

type DictController struct {
	base.BaseController
}

// List 列表
// @Tags 字典管理
// @Summary 列表
// @Description 字典列表
// @Router /api/v1/dict [get]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *DictController) List(c *gin.Context) {
	var (
		dictService service.DictService
	)

	data, err := dictService.List()
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, data)
}

// Create 创建
// @Tags 字典管理
// @Summary 创建
// @Description 字典创建
// @Param token header string true "认证Token"
// @Param data body object true "创建参数" SchemaExample({"name":"状态","title":"状态","status":1})
// @Router /api/v1/dict [post]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *DictController) Create(c *gin.Context) {
	var (
		dictService  service.DictService
		dictValidate validate.Dict
		req          model.Dict
	)

	err := c.ShouldBind(&req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	err = copier.Copy(&dictValidate, &req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.Dict{}.GetValidate(dictValidate, "create")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	data, err := dictService.Create(req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, data)
}

// Update 更新
// @Tags 字典管理
// @Summary 更新
// @Description 字典更新
// @Param token header string true "认证Token"
// @Param id path int true "字典ID"
// @Param data body object true "更新参数" SchemaExample({"name":"状态","title":"状态","status":1})
// @Router /api/v1/dict/{id} [put]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *DictController) Update(c *gin.Context) {
	var (
		dictService  service.DictService
		dictValidate validate.Dict
		req          model.Dict
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
	dictValidate.ID = id
	err = copier.Copy(&dictValidate, &req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.Dict{}.GetValidate(dictValidate, "update")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	data, err := dictService.Update(req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "更新成功", data)
}

// Detail 详情
// @Tags 字典管理
// @Summary 详情
// @Description 字典详情
// @Param token header string true "认证Token"
// @Param id path int true "字典ID"
// @Router /api/v1/dict/{id} [get]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *DictController) Detail(c *gin.Context) {
	var (
		dictService service.DictService
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

	data, err := dictService.Detail(id)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, data)
}

// Delete 删除
// @Tags 字典管理
// @Summary 删除
// @Description 字典删除
// @Param token header string true "认证Token"
// @Param id path int true "字典ID"
// @Router /api/v1/dict/{id} [delete]
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
func (s *DictController) Delete(c *gin.Context) {
	var (
		dictService service.DictService
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

	data, err := dictService.Delete(id)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "删除成功", data)
}
