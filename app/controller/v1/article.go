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

type ArticleController struct {
	common.BaseController
}

// List
// @Tags 文章
// @Summary 列表
// @Router /v1/article [get]
func (s *ArticleController) List(c *gin.Context) {
	var (
		articleService service.ArticleService
		req            validate.ArticleValidate
	)

	err := c.ShouldBindQuery(&req)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	// 验证
	err = validate.GetArticleValidate(req, "list")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error(), nil)
		return
	}

	pageData, err := articleService.List(req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error(), nil)
		return
	}

	s.ApiResponse(c, global.Success, "获取成功", pageData)
}

// Create
// @Tags 文章
// @Summary 创建
// @Router /v1/article [post]
func (s *ArticleController) Create(c *gin.Context) {
	var (
		articleService  service.ArticleService
		articleValidate validate.ArticleValidate
		req             model.ArticleQuery
	)

	err := c.ShouldBind(&req)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	req.UID = s.GetUserId(c)
	// 序列化为 JSON
	jsonData, err := json.Marshal(req)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	// 反序列化到响应结构体
	err = json.Unmarshal(jsonData, &articleValidate)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	// 验证
	err = validate.GetArticleValidate(articleValidate, "create")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error(), nil)
		return
	}

	articleModel, err := articleService.Create(req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error(), nil)
		return
	}

	s.ApiResponse(c, global.Success, "创建成功", articleModel)
}

// Update
// @Tags 文章
// @Summary 更新
// @Router /v1/article/:id [put]
func (s *ArticleController) Update(c *gin.Context) {
	var (
		articleService  service.ArticleService
		articleValidate validate.ArticleValidate
		req             model.ArticleQuery
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
	// 序列化为 JSON
	jsonData, err := json.Marshal(req)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	// 反序列化到响应结构体
	err = json.Unmarshal(jsonData, &articleValidate)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	// 验证
	err = validate.GetArticleValidate(articleValidate, "create")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error(), nil)
		return
	}

	articleModel, err := articleService.Update(req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error(), nil)
		return
	}

	s.ApiResponse(c, global.Success, "更新成功", articleModel)
}

// Detail
// @Tags 文章
// @Summary 详情
// @Router /v1/article/:id [get]
func (s *ArticleController) Detail(c *gin.Context) {
	var (
		articleService  service.ArticleService
		articleValidate validate.ArticleValidate
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

	articleValidate.ID = id

	// 验证
	err = validate.GetArticleValidate(articleValidate, "detail")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error(), nil)
		return
	}

	articleModel, err := articleService.Detail(id)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error(), nil)
		return
	}

	s.ApiResponse(c, global.Success, "查询成功", articleModel)
}

// Delete
// @Tags 文章
// @Summary 删除
// @Router /v1/article/:id [delete]
func (s *ArticleController) Delete(c *gin.Context) {
	var (
		articleService  service.ArticleService
		articleValidate validate.ArticleValidate
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

	articleValidate.ID = id

	// 验证
	err = validate.GetArticleValidate(articleValidate, "delete")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error(), nil)
		return
	}

	articleModel, err := articleService.Delete(id)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error(), nil)
		return
	}

	s.ApiResponse(c, global.Success, "删除成功", articleModel)
}
