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

type ArticleController struct {
	base.BaseController
}

// List 列表
// @Tags 文章管理
// @Summary 列表
// @Description 文章列表
// @Param token header string true "认证Token"
// @Param page query string true "页码"
// @Param pageSize query string true "分页大小"
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
// @Router /api/v1/article [get]
func (s *ArticleController) List(c *gin.Context) {
	var (
		articleService  service.ArticleService
		articleValidate validate.Article
		pageData        global.PageData
	)

	err := c.ShouldBindQuery(&articleValidate)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.Article{}.GetValidate(articleValidate, "list")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	err = copier.Copy(&pageData, &articleValidate)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
	}

	pageData, err = articleService.List(pageData)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, pageData)
}

// Create 创建
// @Tags 文章管理
// @Summary 创建
// @Description 文章创建
// @Accept json
// @Produce json
// @Param token header string true "认证Token"
// @Param data body object true "创建参数" SchemaExample({"title":"Go语言","content":"内容"})
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
// @Router /api/v1/article [post]
func (s *ArticleController) Create(c *gin.Context) {
	var (
		articleService  service.ArticleService
		articleValidate validate.Article
		req             model.Article
	)

	err := c.ShouldBind(&req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	req.UID = s.GetUserId(c)
	err = copier.Copy(&articleValidate, &req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.Article{}.GetValidate(articleValidate, "create")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	articleModel, err := articleService.Create(req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "创建成功", articleModel)
}

// Update 更新
// @Tags 文章管理
// @Summary 更新
// @Description 文章更新
// @Param token header string true "认证Token"
// @Param id path int true "文章ID"
// @Param data body object true "更新参数" SchemaExample({"title":"Go语言","content":"内容"})
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
// @Router /api/v1/article/{id} [put]
func (s *ArticleController) Update(c *gin.Context) {
	var (
		articleService  service.ArticleService
		articleValidate validate.Article
		req             model.Article
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
	err = copier.Copy(&articleValidate, &req)
	if err != nil {
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	// 验证
	err = validate.Article{}.GetValidate(articleValidate, "create")
	if err != nil {
		s.ApiResponse(c, global.ArgsError, err.Error())
		return
	}

	articleModel, err := articleService.Update(req)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "更新成功", articleModel)
}

// Detail 详情
// @Tags 文章管理
// @Summary 详情
// @Description 文章详情
// @Param token header string true "认证Token"
// @Param id path int true "文章ID"
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
// @Router /api/v1/article/{id} [get]
func (s *ArticleController) Detail(c *gin.Context) {
	var (
		articleService service.ArticleService
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

	article, err := articleService.Detail(id)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, article)
}

// Delete 删除
// @Tags 文章管理
// @Summary 删除
// @Description 文章删除
// @Param token header string true "认证Token"
// @Param id path int true "文章ID"
// @Success 200 {object} global.Response{global.Success} "成功返回" Example({"code":0,"msg":"Success","data":[]})
// @Failure 400 {object} global.Response{global.ArgsError} "参数错误" Example({"code":400,"msg":"参数错误","data":[]})
// @Failure 500 {object} global.Response{global.SystemError} "系统错误" Example({"code":500,"msg":"系统错误","data":[]})
// @Router /api/v1/article/{id} [delete]
func (s *ArticleController) Delete(c *gin.Context) {
	var (
		articleService service.ArticleService
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

	articleModel, err := articleService.Delete(id)
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, "删除成功", articleModel)
}
