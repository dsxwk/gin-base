package service

import (
	"gin-base/app/model"
	"gin-base/app/validate"
	"gin-base/common/base"
	"gin-base/common/global"
	"gin-base/helper/utils"
)

type RoleService struct {
	base.BaseService
}

// List 角色列表
// @param pageData global.PageData
// @param search validate.RoleSearchValidate
// @return interface{}, error
func (s *RoleService) List(pageData global.PageData, search validate.RoleSearch) (interface{}, error) {
	var (
		models []model.Roles
	)

	// 获取where条件和参数
	where, args := utils.BuildWhereClause(search, "form")

	db := global.DB.Model(&model.Roles{}).Find(&models)
	// 根据 where 子句添加条件
	if where != "" {
		db = db.Where(where, args...)
	}

	// 获取总记录数
	err := db.Count(&pageData.Total).Error
	if err != nil {
		return pageData, err
	}

	if pageData.IsPage == nil || *pageData.IsPage {
		// 获取分页默认为第一页，每页10条记录
		offset, limit := utils.Pagination(pageData.Page, pageData.PageSize)
		// 执行分页查询
		err = db.Offset(offset).
			Limit(limit).
			Find(&models).Error
		if err != nil {
			return pageData, err
		}
	}

	if pageData.IsPage == nil || *pageData.IsPage {
		pageData.List = models
		return pageData, nil
	}

	return models, nil
}

// Create 创建
// @param: m model.Roles
// @return: model.Roles, error
func (s *RoleService) Create(m model.Roles) (model.Roles, error) {
	err := global.DB.Create(&m).Error
	if err != nil {
		return m, err
	}

	return m, nil
}

// Update 更新
// @param m model.Roles
// @return model.Roles, error
func (s *RoleService) Update(m model.Roles) (model.Roles, error) {
	err := global.DB.Updates(&m).Error
	if err != nil {
		return m, err
	}

	return m, nil
}

// Detail 详情
// @param id int64
// @return m model.Roles, err error
func (s *RoleService) Detail(id int64) (m model.Roles, err error) {
	err = global.DB.
		First(&m, id).Error
	if err != nil {
		return m, err
	}

	return m, nil
}

// Delete 删除
// @param id int64
// @return m model.Roles, err error
func (s *RoleService) Delete(id int64) (m model.Roles, err error) {
	err = global.DB.Delete(&m, id).Error
	if err != nil {
		return m, err
	}

	return m, nil
}
