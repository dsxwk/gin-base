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
func (this *LoginService) List(requestData validate.UserListRequest) (global.PageData, error) {
	var (
		userModel []model.User
		pageData  global.PageData
	)

	// 获取分页默认为第一页，每页10条记录
	offset, limit := utils.Pagination(requestData.Page, requestData.PageSize)

	db := global.DB.Find(&userModel)

	err := db.Count(&pageData.Total).Error
	if err != nil {
		return pageData, err
	}
	err = db.Offset(offset).Limit(limit).Find(&userModel).Error
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

	pageData.Page = requestData.Page
	pageData.PageSize = requestData.PageSize
	pageData.List = userModel

	return pageData, nil
}
