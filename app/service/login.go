package service

import (
	"errors"
	"fmt"
	"gin-base/app/model"
	"gin-base/common/base"
	"gin-base/common/global"
	"gin-base/config"
	"gin-base/helper"
	"gorm.io/gorm"
)

type LoginService struct {
	base.BaseService
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
			return m, errors.New("login.accountErr")
		}
	}

	check := helper.BcryptCheck(password, m.Password)
	if !check {
		return m, errors.New("login.pwdErr")
	}

	if m.Status != 1 {
		return m, errors.New("login.accountDisabled")
	}

	// 发布事件
	global.Event.Publish(config.Event{
		Name: "userLogin",
		Data: map[string]interface{}{
			"username": username,
			"password": password,
		},
	})

	// redis 发布消息
	_ = global.Redis.Publish("userLogin", fmt.Sprintf("redis event for user login: %s", username))

	return m, nil
}
