package service

import (
	"gin-base/app/model"
	"gin-base/app/validate"
	"gin-base/common/base"
	"gin-base/common/global"
	"gin-base/helper/utils"
)

type UserService struct {
	base.BaseService
}

// List 列表
// @param: pageData global.PageData
// @param: search validate.UserSearch
// @return: global.PageData, error
func (s *UserService) List(pageData global.PageData, search validate.UserSearch) (global.PageData, error) {
	var (
		models []model.User
	)

	// 获取分页默认为第一页，每页10条记录
	offset, limit := utils.Pagination(pageData.Page, pageData.PageSize)
	db := global.DB.Model(&model.User{}).Scopes(model.Search(search)).Find(&models)

	// 获取总记录数
	err := db.Count(&pageData.Total).Error
	if err != nil {
		return pageData, err
	}

	// 执行分页查询
	err = db.
		Preload("UserRoles").
		Offset(offset).
		Limit(limit).
		Find(&models).Error
	if err != nil {
		return pageData, err
	}

	pageData.List = models

	return pageData, nil
}

// Create 创建
// @param: m model.User
// @return: model.User, error
func (s *UserService) Create(m model.User) (model.User, error) {
	var (
		userRoles []model.UserRoles
	)

	// 处理密码
	m.Password = utils.BcryptHash(m.Password)

	tx := global.DB.Begin()
	err := tx.Create(&m).Error
	if err != nil {
		return m, err
	}

	for _, v := range m.UserRoles {
		userRoles = append(userRoles, model.UserRoles{
			UserID: m.ID,
			RoleID: v.RoleID,
			Name:   v.Name,
		})
	}

	if userRoles != nil {
		err = tx.Create(&userRoles).Error
		if err != nil {
			return m, err
		}
	}

	tx.Commit()

	return m, nil
}

// Update 更新
// @param m model.User
// @return model.User, error
func (s *UserService) Update(m model.User) (model.User, error) {
	var (
		userRoles []model.UserRoles
	)

	for _, v := range m.UserRoles {
		userRoles = append(userRoles, model.UserRoles{
			UserID: m.ID,
			RoleID: v.RoleID,
			Name:   v.Name,
		})
	}

	data := utils.StructToMapFilter(m, []string{
		"createdAt", "updatedAt", "deletedAt", "password", "userRoles",
	})

	tx := global.DB.Begin()
	err := tx.Model(&m).
		Where("id = ?", m.ID).
		Updates(data).Error
	if err != nil {
		return m, err
	}

	err = tx.Where("user_id", m.ID).
		Delete(&model.UserRoles{}).Error
	if err != nil {
		return m, err
	}

	if userRoles != nil {
		err = tx.Create(&userRoles).Error
		if err != nil {
			return m, err
		}
	}

	tx.Commit()

	return m, nil
}

// Detail 详情
// @param id int64
// @return m model.User, err error
func (s *UserService) Detail(id int64) (m model.User, err error) {
	err = global.DB.
		Preload("UserRoles").
		First(&m, id).Error
	if err != nil {
		return m, err
	}

	return m, nil
}

// Delete 删除
// @param id int64
// @return m model.User, err error
func (s *UserService) Delete(id int64) (m model.User, err error) {
	err = global.DB.Delete(&m, id).Error
	if err != nil {
		return m, err
	}

	return m, nil
}
