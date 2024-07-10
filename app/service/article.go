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

// @function: List
// @description: 列表
// @param: req validate.ArticleValidate
// @return: global.PageData, error
func (s *ArticleService) List(req validate.ArticleValidate) (global.PageData, error) {
	// scan
	/*type field struct {
		ID         int64          `json:"id"`
		UID        int64          `json:"uid"`
		Title      string         `json:"title"`
		Content    string         `json:"content"`
		CategoryID int64          `json:"category_id"`
		Username   string         `json:"username"`
		Name       string         `json:"name"`
		CreatedAt  string         `json:"created_at"`
		UpdatedAt  string         `json:"updated_at"`
		DeletedAt  gorm.DeletedAt `json:"deleted_at"`
	}*/

	var (
		articleModel []model.Article
		pageData     global.PageData
		//fields       []field
	)

	// 获取分页默认为第一页，每页10条记录
	offset, limit := utils.Pagination(req.Page, req.PageSize)

	// join
	//db := global.DB.Joins("LEFT JOIN user ON article.uid = user.id LEFT JOIN category ON article.category_id = category.id").Select("article.*, user.username, category.name").Find(&articleModel).scan(&fields)

	db := global.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username, full_name, nickname, email, gender, age")
	}).Preload("Category", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).Find(&articleModel)

	err := db.Count(&pageData.Total).Error
	if err != nil {
		return pageData, err
	}

	err = db.Offset(offset).Limit(limit).Error
	if err != nil {
		return pageData, err
	}

	pageData.Page = req.Page
	pageData.PageSize = req.PageSize
	pageData.List = articleModel

	return pageData, nil
}

// @function: Create
// @description: 创建
// @param: req model.Article
// @return: model.Article, error
func (s *ArticleService) Create(req model.Article) (model.Article, error) {
	err := global.DB.Create(&req).Error
	if err != nil {
		return req, err
	}

	return req, nil
}

// @function: Update
// @description: 更新
// @param: req model.Article
// @return: model.Article, error
func (this *ArticleService) Update(req model.Article) (model.Article, error) {
	err := global.DB.Updates(&req).Error
	if err != nil {
		return req, err
	}

	return req, nil
}

// @function: Detail
// @description: 详情
// @param: id int64
// @return: model.Article, error
func (s *ArticleService) Detail(id int64) (model.Article, error) {
	var (
		articleModel model.Article
	)

	err := global.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username, full_name, nickname, email, gender, age")
	}).Preload("Category", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).First(&articleModel, id).Error

	if err != nil {
		return articleModel, err
	}

	return articleModel, nil
}

// @function: Delete
// @description: 删除
// @param: id int64
// @return: model.Article, error
func (s *ArticleService) Delete(id int64) (model.Article, error) {
	var (
		articleModel model.Article
	)

	err := global.DB.Delete(&articleModel, id).Error
	if err != nil {
		return articleModel, err
	}

	return articleModel, nil
}
