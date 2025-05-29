package service

import (
	"errors"
	"fmt"
	"gin-base/app/model"
	"gin-base/common"
	"gin-base/common/global"
	"gin-base/helper/utils"
	"gorm.io/gorm"
)

type MenuService struct {
	common.BaseService
}

// List 列表
// @return bool
func (s *MenuService) List() (models []model.Menu, err error) {
	err = global.DB.
		Preload("MenuRoles").
		Preload("MenuAction", func(db *gorm.DB) *gorm.DB {
			return db.Preload("ActionRoles").
				Order("sort asc").
				Select("id,menu_id,type,name,is_link,sort,created_at,updated_at")
		}).
		Preload("MenuAction.ActionRoles").
		Order("sort asc").
		Find(&models).Error

	if err != nil {
		return models, err
	}

	models = s.MenuTree(models, 0)

	return models, nil
}

// MenuTree 数据子集递归
// @param menus []model.Menu
// @param pid int64
// @return tree []model.Menu
func (s *MenuService) MenuTree(menus []model.Menu, pid int64) (tree []model.Menu) {
	for _, menu := range menus {
		if menu.Pid == pid {
			children := s.MenuTree(menus, menu.ID)
			menu.Children = children
			menu.Meta.Roles = utils.ArrayColumn(menu.MenuRoles, func(m *model.MenuRoles) int64 { return m.RoleID })
			tree = append(tree, menu)
		}
	}
	return tree
}

// Create 创建
// @param m model.Menu
// @return model.Menu, error
func (s *MenuService) Create(m model.Menu) (model.Menu, error) {
	var (
		menuRoles []model.MenuRoles
	)

	tx := global.DB.Begin()
	err := tx.Create(&m).Error
	if err != nil {
		return m, err
	}

	for _, v := range m.Meta.Roles {
		menuRoles = append(menuRoles, model.MenuRoles{
			MenuID: m.ID,
			RoleID: v,
		})
	}

	err = tx.Create(&menuRoles).Error
	if err != nil {
		return m, err
	}
	tx.Commit()

	return m, nil
}

// Update 更新
// @param m model.Menu
// @return model.Menu, error
func (s *MenuService) Update(m model.Menu) (model.Menu, error) {
	var (
		menuRole  model.MenuRoles
		menuRoles []model.MenuRoles
	)

	for _, v := range m.Meta.Roles {
		menuRoles = append(menuRoles, model.MenuRoles{
			MenuID: m.ID,
			RoleID: v,
		})
	}

	tx := global.DB.Begin()
	err := tx.Updates(&m).Error
	if err != nil {
		return m, err
	}

	err = tx.Where("menu_id", m.ID).Delete(&menuRole).Error
	if err != nil {
		return m, err
	}

	err = tx.Create(&menuRoles).Error
	if err != nil {
		return m, err
	}
	tx.Commit()

	return m, nil
}

// Delete 删除
// @param id int64
// @return m model.Menu, err error
func (s *MenuService) Delete(id int64) (m model.Menu, err error) {
	err = global.DB.Delete(&m, id).Error
	if err != nil {
		return m, err
	}

	return m, nil
}

// ActionList 功能列表
// @param menuId int64
// @return models []model.MenuAction, err error
func (s *MenuService) ActionList(menuId int64) (models []model.MenuAction, err error) {
	err = global.DB.
		Where("menu_id = ?", menuId).
		Order("sort asc").
		Find(&models).
		Error
	if err != nil {
		return models, err
	}

	return models, nil
}

// ActionCreate 功能创建
// @param m model.MenuAction
// @return model.MenuAction, error
func (s *MenuService) ActionCreate(m model.MenuAction) (model.MenuAction, error) {
	db := global.DB.Create(&m)

	err := db.Error
	if err != nil {
		return m, err
	}

	return m, nil
}

// ActionUpdate 功能更新
// @param m model.MenuAction
// @return model.MenuAction, error
func (s *MenuService) ActionUpdate(m model.MenuAction) (model.MenuAction, error) {
	err := global.DB.Updates(&m).Error
	if err != nil {
		return m, err
	}

	return m, nil
}

// ActionDelete 功能删除
// @param id int64
// @param menuID int64
// @return m model.MenuAction, err error
func (s *MenuService) ActionDelete(id int64, menuID int64) (m model.MenuAction, err error) {
	if err = global.DB.Where("id = ?", id).First(&m).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return m, err
		}
	}

	if m.MenuID != menuID {
		return m, fmt.Errorf("删除失败,该功能id【%d】不属于该菜单id【%d】", id, menuID)
	}

	err = global.DB.Delete(&m, id).Error
	if err != nil {
		return m, err
	}

	return m, nil
}
