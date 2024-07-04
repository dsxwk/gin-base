package service

import (
	"gin-base/app/model"
	"gin-base/app/validate"
	"gin-base/common"
	"gin-base/common/global"
	"gin-base/helper/utils"
	"gorm.io/gorm"
)

type ArticleService struct {
	common.BaseService
}

// 列表
func (this *ArticleService) List(requestData validate.ArticleListValidate) (global.PageData, error) {
	var (
		articleModel []model.Article
		pageData     global.PageData
	)

	// 获取分页默认为第一页，每页10条记录
	offset, limit := utils.Pagination(requestData.Page, requestData.PageSize)

	db := global.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username, full_name, nickname, email, gender, age")
	}).Preload("Category", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).Find(&articleModel)

	err := db.Count(&pageData.Total).Error
	if err != nil {
		return pageData, err
	}
	err = db.Offset(offset).Limit(limit).Find(&articleModel).Error
	if err != nil {
		return pageData, err
	}

	// 处理时间
	for key, value := range articleModel {
		if value.CreatedAt != nil {
			createdAt := utils.FormatTime(*value.CreatedAt)
			articleModel[key].CreatedAt = &createdAt
		}

		if value.UpdatedAt != nil {
			updatedAt := utils.FormatTime(*value.UpdatedAt)
			articleModel[key].UpdatedAt = &updatedAt
		}

		if value.DeletedAt != nil {
			deletedAt := utils.FormatTime(*value.DeletedAt)
			articleModel[key].DeletedAt = &deletedAt
		}
	}

	pageData.Page = requestData.Page
	pageData.PageSize = requestData.PageSize
	pageData.List = articleModel

	return pageData, nil
}
