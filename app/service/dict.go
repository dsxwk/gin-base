package service

import (
	"gin-base/app/model"
	"gin-base/common"
	"gin-base/common/global"
)

type DictService struct {
	common.BaseService
}

// List 列表
// @return models []model.Dict, err error
func (s *DictService) List() (models []model.Dict, err error) {
	err = global.DB.
		Order("sort asc").
		Find(&models).Error

	if err != nil {
		return models, err
	}

	models = s.GetTree(models, 0)

	return models, nil
}

// GetTree 数据子集递归
// @param menus []model.Dict
// @param pid int64
// @return tree []model.Dict
func (s *DictService) GetTree(menus []model.Dict, pid int64) (tree []model.Dict) {
	for _, menu := range menus {
		if menu.Pid == pid {
			children := s.GetTree(menus, menu.ID)
			menu.Children = children
			tree = append(tree, menu)
		}
	}
	return tree
}
