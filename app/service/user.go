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

// @function: List
// @description: 列表
// @param: req validate.UserValidate
// @return: global.PageData, error
func (s *UserService) List(req validate.UserValidate) (global.PageData, error) {
	var (
		userModel []model.User
		pageData  global.PageData
	)

	// 获取分页默认为第一页，每页10条记录
	offset, limit := utils.Pagination(req.Page, req.PageSize)

	db := global.DB.Find(&userModel)

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
	pageData.List = userModel

	return pageData, nil
}
