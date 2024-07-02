package service

import (
	"fmt"
	"gin-base/common"
)

type LoginService struct {
	common.BaseService
}

// 登录
func (this *LoginService) Login(id int64, exp int64) error {
	fmt.Println("id:", id, "exp:", exp)
	return nil
}
