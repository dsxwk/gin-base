package v1

import (
	"gin-base/app/service"
	"gin-base/app/validate"
	"gin-base/common"
	"gin-base/common/global"
	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	common.BaseController
}

// @Tags    文章
// @Summary 列表
// @Router  /v1/user [get]
func (this *ArticleController) List(c *gin.Context) {
	var (
		articleService service.ArticleService
		requestData    validate.ArticleListValidate
	)

	err := c.ShouldBindQuery(&requestData)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	// 验证2
	err = validate.GetArticleListValidate(requestData)
	if err != nil {
		this.ApiResponse(c, global.Error, err.Error(), nil)
		return
	}

	pageData, err := articleService.List(requestData)
	if err != nil {
		global.Log.Error(err.Error())
		this.ApiResponse(c, global.SystemError, err.Error(), nil)
		return
	}

	this.ApiResponse(c, global.Success, "获取成功", pageData)
}
