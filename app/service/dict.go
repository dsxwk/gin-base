package service

import (
	"errors"
	"gin-base/app/model"
	"gin-base/common"
	"gin-base/common/global"
	"gorm.io/gorm"
)

type DictService struct {
	common.BaseService
}

// List 列表
// @return models []model.Dict, err error
func (s *DictService) List() (models []model.Dict, err error) {
	models, err = s.All()

	if err != nil {
		return models, err
	}

	models = s.GetTree(models, 0)

	return models, nil
}

// All 获取所有字典
// @return models []model.Dict, err error
func (s *DictService) All() (models []model.Dict, err error) {
	err = global.DB.
		Order("sort asc").
		Find(&models).Error

	if err != nil {
		return models, err
	}

	return models, nil
}

// GetTree 数据子集递归
// @param dicts []model.Dict
// @param pid int64
// @return tree []model.Dict
func (s *DictService) GetTree(dicts []model.Dict, pid int64) (tree []model.Dict) {
	for _, data := range dicts {
		if data.Pid == pid {
			children := s.GetTree(dicts, data.ID)
			data.Children = children
			tree = append(tree, data)
		}
	}
	return tree
}

// Create 创建
// @param m model.Dict
// @return model.Dict, error
func (s *DictService) Create(m model.Dict) (model.Dict, error) {
	err := global.DB.Create(&m).Error
	if err != nil {
		return m, err
	}

	return m, nil
}

// Update 更新
// @param m model.Dict
// @return model.Dict, error
func (s *DictService) Update(m model.Dict) (model.Dict, error) {
	err := global.DB.Updates(&m).Error
	if err != nil {
		return m, err
	}

	return m, nil
}

// Detail 详情
// @param id int64
// @return m model.Dict, err error
func (s *DictService) Detail(id int64) (m model.Dict, err error) {
	var (
		dicts []model.Dict
	)

	dicts, err = s.All()

	if err != nil {
		return m, err
	}

	err = global.DB.
		First(&m, id).Error
	if err != nil {
		return m, err
	}

	m.Children = s.GetTree(dicts, m.ID)

	return m, nil
}

// Delete 删除
// @param id int64
// @return m model.Dict, err error
func (s *DictService) Delete(id int64) (m model.Dict, err error) {
	var (
		detail model.Dict
	)

	err = global.DB.
		Where("pid = ?", id).
		First(&detail).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return m, err
		}
	}

	if detail.ID > 0 {
		return m, errors.New("存在子集，无法删除")
	}

	err = global.DB.Delete(&m, id).Error
	if err != nil {
		return m, err
	}

	return m, nil
}
