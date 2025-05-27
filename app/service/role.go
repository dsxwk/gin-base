package service

import (
	"gin-base/app/model"
	"gin-base/app/validate"
	"gin-base/common"
	"gin-base/common/global"
	"gin-base/helper/utils"
)

type RoleService struct {
	common.BaseService
}

// Your Function Name Your Description
// @param YourParam string
// @return bool
func (s *RoleService) List(req validate.RoleValidate, search validate.RoleSearchValidate) (global.PageData, error) {
	var (
		roleModels []model.Roles
		roles      []model.RolesQuery
		pageData   global.PageData
	)

	// 获取where条件和参数
	where, args := utils.BuildWhereClause(search, "form")

	// 获取分页默认为第一页，每页10条记录
	offset, limit := utils.Pagination(req.Page, req.PageSize)

	db := global.DB.Find(&roleModels)
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
	err = db.Find(&roleModels).
		Scan(&roles).
		Offset(offset).
		Limit(limit).Error
	if err != nil {
		return pageData, err
	}

	pageData.Page = req.Page
	pageData.PageSize = req.PageSize
	pageData.List = roles

	return pageData, nil
}
