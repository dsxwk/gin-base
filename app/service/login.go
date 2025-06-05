package service

import (
	"errors"
	"gin-base/app/model"
	"gin-base/common"
	"gin-base/common/extend/event"
	"gin-base/common/global"
	"gin-base/helper/utils"
	"gorm.io/gorm"
)

type LoginService struct {
	common.BaseService
}

// Login 登录
// @param username string
// @param password string
// @return m model.User, error
func (s *LoginService) Login(username string, password string) (m model.User, err error) {
	if err = global.DB.
		Where("username = ?", username).
		Preload("UserRoles").
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return m, errors.New("登录账号错误")
		}
	}

	check := utils.BcryptCheck(password, m.Password)
	if !check {
		return m, errors.New("登录密码错误")
	}

	if m.Status != 1 {
		return m, errors.New("账号已被禁用")
	}

	// 发布事件
	e := event.Event{
		Name: "user_login",
		Data: map[string]interface{}{
			"username": username,
			"password": password,
		},
	}
	global.Event.Publish(e)

	return m, nil
}
