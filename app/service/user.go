package service

import (
	"gin-base/app/model"
	"gin-base/app/validate"
	"gin-base/common"
	"gin-base/common/global"
	"gin-base/helper/utils"
)

type UserService struct {
	common.BaseService
}

// List 列表
// @param: req validate.UserValidate, search validate.UserSearch
// @return: global.PageData, error
func (s *UserService) List(req validate.UserValidate, search validate.UserSearch) (global.PageData, error) {
	var (
		userModel []model.User
		userQuery []model.UserQuery
		pageData  global.PageData
	)

	// 获取where条件和参数
	where, args := utils.BuildWhereClause(search, "form")

	// 获取分页默认为第一页，每页10条记录
	offset, limit := utils.Pagination(req.Page, req.PageSize)

	db := global.DB.Find(&userModel)
	// 根据 where 子句添加条件
	if where != "" {
		db = db.Where(where, args...)
	}

	// 获取总记录数
	err := db.Count(&pageData.Total).Error
	if err != nil {
		return pageData, err
	}

	// 执行分页查询
	err = db.Offset(offset).Limit(limit).Find(&userModel).Scan(&userQuery).Error
	if err != nil {
		return pageData, err
	}

	for K, m := range userModel {
		userQuery[K].CreatedAt = utils.FormatTime(m.CreatedAt)
		userQuery[K].UpdatedAt = utils.FormatTime(m.UpdatedAt)
	}

	pageData.Page = req.Page
	pageData.PageSize = req.PageSize
	pageData.List = userQuery

	return pageData, nil
}

// Create 创建
// @param: req model.User
// @return: model.User, error
func (s *UserService) Create(req model.User) (model.User, error) {
	// 处理密码
	req.Password = utils.BcryptHash(req.Password)

	err := global.DB.Create(&req).Error
	if err != nil {
		return req, err
	}

	return req, nil
}

// Update 更新
// @param: req model.User
// @return: model.User, error
func (this *UserService) Update(req model.User) (model.User, error) {
	err := global.DB.Updates(&req).Error
	if err != nil {
		return req, err
	}

	return req, nil
}

// Detail 详情
// @param: id int64
// @return: model.UserQuery, error
func (s *UserService) Detail(id int64) (model.UserQuery, error) {
	var (
		userModel model.User
		userQuery model.UserQuery
	)

	err := global.DB.First(&userModel, id).Scan(&userQuery).Error
	if err != nil {
		return userQuery, err
	}

	userQuery.CreatedAt = utils.FormatTime(userModel.CreatedAt)
	userQuery.UpdatedAt = utils.FormatTime(userModel.UpdatedAt)

	return userQuery, nil
}

// Delete 删除
// @param: id int64
// @return: model.User, error
func (s *UserService) Delete(id int64) (model.User, error) {
	var (
		userModel model.User
	)

	err := global.DB.Delete(&userModel, id).Error
	if err != nil {
		return userModel, err
	}

	return userModel, nil
}
