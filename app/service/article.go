package service

import (
	"gin-base/app/model"
	"gin-base/common/base"
	"gin-base/common/global"
	"gin-base/helper"
	"gorm.io/gorm"
)

type ArticleService struct {
	base.BaseService
}

// List 列表
// @param pageData global.PageData
// @return global.PageData, error
func (s *ArticleService) List(pageData global.PageData) (global.PageData, error) {
	var (
		models []model.Article
	)

	// 获取分页默认为第一页，每页10条记录
	offset, limit := helper.Pagination(pageData.Page, pageData.PageSize)

	// join
	// db := global.DB.Joins("LEFT JOIN user ON article.uid = user.id LEFT JOIN category ON article.category_id = category.id").Select("article.*, user.username, category.name").Find(&models)

	db := global.DB.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, username, full_name, nickname, email, gender, age")
		}).Preload("Category", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).
		Find(&models)

	// 获取总记录数
	err := db.Count(&pageData.Total).Error
	if err != nil {
		return pageData, err
	}

	// 执行分页查询
	err = db.Offset(offset).
		Limit(limit).
		Find(&models).Error
	if err != nil {
		return pageData, err
	}

	pageData.List = models

	return pageData, nil
}

// Create 创建
// @param m model.Article
// @return model.Article, error
func (s *ArticleService) Create(m model.Article) (model.Article, error) {
	err := global.DB.Create(&m).Error
	if err != nil {
		return m, err
	}

	return m, nil
}

// Update 更新
// @param m model.Article
// @return model.Article, error
func (s *ArticleService) Update(m model.Article) (model.Article, error) {
	err := global.DB.Updates(&m).Error
	if err != nil {
		return m, err
	}

	return m, nil
}

// Detail 详情
// @param id int64
// @return model.Article, error
func (s *ArticleService) Detail(id int64) (model.Article, error) {
	var (
		m model.Article
	)

	err := global.DB.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, username, full_name, nickname, email, gender, age")
		}).Preload("Category", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).
		First(&m, id).Error

	if err != nil {
		return m, err
	}

	return m, nil
}

// Delete 删除
// @param id int64
// @return model.Article, error
func (s *ArticleService) Delete(id int64) (model.Article, error) {
	err := global.DB.Delete(&model.Article{}, id).Error
	if err != nil {
		return model.Article{}, err
	}

	return model.Article{}, nil
}
