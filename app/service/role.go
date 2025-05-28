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

// List 角色列表
// @param pageData global.PageData, search validate.RoleSearchValidate
// @return interface{}, error
func (s *RoleService) List(pageData global.PageData, search validate.RoleSearchValidate) (interface{}, error) {
	var (
		roles []model.RolesQuery
	)

	// 获取where条件和参数
	where, args := utils.BuildWhereClause(search, "form")

	db := global.DB.Model(&model.Roles{}).Find(&roles)
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
			Find(&roles).Error
		if err != nil {
			return pageData, err
		}
	}

	if pageData.IsPage == nil || *pageData.IsPage {
		pageData.List = roles
		return pageData, nil
	}

	return roles, nil
}
