package service

import (
	"errors"
	"gin-base/app/model"
	"gin-base/common"
	"gin-base/common/global"
	"gin-base/helper/utils"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"time"
)

type ArticleService struct {
	common.BaseService
}

// List 列表
// @param pageData global.PageData
// @return global.PageData, error
func (s *ArticleService) List(pageData global.PageData) (global.PageData, error) {
	var (
		articleModel []model.Article
		articleQuery []model.ArticleQuery
	)

	// 获取分页默认为第一页，每页10条记录
	offset, limit := utils.Pagination(pageData.Page, pageData.PageSize)

	// join
	// db := global.DB.Joins("LEFT JOIN user ON article.uid = user.id LEFT JOIN category ON article.category_id = category.id").Select("article.*, user.username, category.name").Find(&articleModel)

	db := global.DB.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, username, full_name, nickname, email, gender, age")
		}).Preload("Category", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).
		Find(&articleModel)

	// 获取总记录数
	err := db.Count(&pageData.Total).Error
	if err != nil {
		return pageData, err
	}

	// 执行分页查询
	err = db.Offset(offset).
		Limit(limit).
		Find(&articleModel).Error
	if err != nil {
		return pageData, err
	}

	err = copier.Copy(&articleQuery, &articleModel)
	if err != nil {
		return pageData, err
	}

	for k, m := range articleModel {
		articleQuery[k].Tag = m.GetTag()
	}

	pageData.List = articleQuery

	return pageData, nil
}

// Create 创建
// @param req model.Article
// @return model.Article, error
func (s *ArticleService) Create(req model.ArticleQuery) (model.Article, error) {
	var (
		articleModel model.Article
	)

	err := copier.Copy(&articleModel, &req)
	if err != nil {
		return articleModel, err
	}
	articleModel.Tag = articleModel.SetTag(req.Tag)

	err = global.DB.Create(&articleModel).Error
	if err != nil {
		return articleModel, err
	}

	return articleModel, nil
}

// Update 更新
// @param req model.Article
// @return model.Article, error
func (this *ArticleService) Update(req model.ArticleQuery) (model.Article, error) {
	var (
		articleModel model.Article
	)

	err := copier.Copy(&articleModel, &req)
	if err != nil {
		return articleModel, err
	}
	articleModel.Tag = articleModel.SetTag(req.Tag)

	ok, err := global.Redis.Lock("test:lock", 20*time.Second)
	if err != nil {
		return articleModel, err
	}
	if !ok {
		return articleModel, errors.New("请稍后尝试")
	}

	// 模拟耗时
	time.Sleep(3 * time.Second)

	err = global.DB.Updates(&articleModel).Error
	if err != nil {
		return articleModel, err
	}

	err = global.Redis.UnLock("test:lock")
	if err != nil {
		return articleModel, err
	}

	return articleModel, nil
}

// Detail 详情
// @param id int64
// @return model.ArticleQuery, error
func (s *ArticleService) Detail(id int64) (model.ArticleQuery, error) {
	var (
		articleModel model.Article
		articleQuery model.ArticleQuery
	)

	err := global.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username, full_name, nickname, email, gender, age")
	}).Preload("Category", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).First(&articleModel, id).Scan(&articleQuery).Error

	if err != nil {
		return articleQuery, err
	}

	return articleQuery, nil
}

// Delete 删除
// @param id int64
// @return model.Article, error
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
