package service

import (
	"errors"
	"gin-base/app/model"
	"gin-base/common"
	"gin-base/common/global"
	"gin-base/helper/utils"
	"gorm.io/gorm"
)

type LoginService struct {
	common.BaseService
}

// @function: Login
// @description: 登录
// @param: username string
// @param: password string
// @return: model.User, error
func (s *LoginService) Login(username string, password string) (model.User, error) {
	var (
		userModel model.User
	)

	if err := global.DB.Where("username = ?", username).First(&userModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return userModel, errors.New("登录账号错误")
		}
	}

	check := utils.BcryptCheck(password, userModel.Password)
	if !check {
		return userModel, errors.New("登录密码错误")
	}

	return userModel, nil
}
