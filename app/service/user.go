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

// 列表
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

	// 处理时间
	for key, value := range userModel {
		if value.CreatedAt != nil {
			createdAt := utils.FormatTime(*value.CreatedAt)
			userModel[key].CreatedAt = &createdAt
		}

		if value.UpdatedAt != nil {
			updatedAt := utils.FormatTime(*value.UpdatedAt)
			userModel[key].UpdatedAt = &updatedAt
		}

		if value.DeletedAt != nil {
			deletedAt := utils.FormatTime(*value.DeletedAt)
			userModel[key].DeletedAt = &deletedAt
		}
	}

	pageData.Page = req.Page
	pageData.PageSize = req.PageSize
	pageData.List = userModel

	return pageData, nil
}
