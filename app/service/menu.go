package service

import (
	"errors"
	"fmt"
	"gin-base/app/model"
	"gin-base/app/validate"
	"gin-base/common/base"
	"gin-base/common/global"
	"gin-base/helper/utils"
	"gorm.io/gorm"
)

type MenuService struct {
	base.BaseService
}

// List 列表
// @param search validate.MenuSearch
// @return models []model.Menu, err error
func (s *MenuService) List(search validate.MenuSearch) (models []model.Menu, err error) {
	models, err = s.All(search.RoleIds)

	if err != nil {
		return models, err
	}

	models = s.GetTree(models, 0)

	return models, nil
}

// All 获取所有菜单
// @param search validate.MenuSearch
// @return models []model.Menu, err error
func (s *MenuService) All(roleIds string) (models []model.Menu, err error) {
	var (
		roleIdsArr []int64
	)

	if roleIds != "" {
		roleIds = `[` + roleIds + `]`
		roleIdsArr = utils.ConvertToArr[int64](&roleIds)
	}

	db := global.DB.
		Preload("MenuRoles").
		Preload("MenuAction").
		Preload("MenuAction.ActionRoles").
		Joins("LEFT JOIN menu_roles ON menu_roles.menu_id = menu.id").
		Order("sort asc").
		Group("menu.id")

	if roleIds != "" {
		db = db.Where("menu_roles.role_id IN ?", roleIdsArr)
	}

	err = db.Find(&models).Error
	if err != nil {
		return models, err
	}

	return models, nil
}

// GetTree 数据子集递归
// @param menus []model.Menu
// @param pid int64
// @return tree []model.Menu
func (s *MenuService) GetTree(menus []model.Menu, pid int64) (tree []model.Menu) {
	for _, data := range menus {
		if data.Pid == pid {
			children := s.GetTree(menus, data.ID)
			data.Children = children
			data.Meta.Roles = utils.ArrayColumn(data.MenuRoles, func(m *model.MenuRoles) int64 { return m.RoleID })
			tree = append(tree, data)
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

	for _, v := range m.MenuRoles {
		menuRoles = append(menuRoles, model.MenuRoles{
			MenuID: m.ID,
			RoleID: v.RoleID,
			Name:   v.Name,
		})
	}

	if menuRoles != nil {
		err = tx.Create(&menuRoles).Error
		if err != nil {
			return m, err
		}
	}

	tx.Commit()

	return m, nil
}

// Update 更新
// @param m model.Menu
// @return model.Menu, error
func (s *MenuService) Update(m model.Menu) (model.Menu, error) {
	var (
		menuRoles []model.MenuRoles
	)

	tx := global.DB.Begin()
	err := tx.Updates(&m).Error
	if err != nil {
		return m, err
	}

	err = tx.Where("menu_id", m.ID).Delete(&model.MenuRoles{}).Error
	if err != nil {
		return m, err
	}

	for _, v := range m.MenuRoles {
		menuRoles = append(menuRoles, model.MenuRoles{
			MenuID: m.ID,
			RoleID: v.RoleID,
			Name:   v.Name,
		})
	}

	if menuRoles != nil {
		err = tx.Create(&menuRoles).Error
		if err != nil {
			return m, err
		}
	}

	tx.Commit()

	return m, nil
}

// Detail 详情
// @param id int64
// @return m model.Menu, err error
func (s *MenuService) Detail(id int64) (m model.Menu, err error) {
	var (
		menus  []model.Menu
		search validate.MenuSearch
	)

	menus, err = s.All(search.RoleIds)

	if err != nil {
		return m, err
	}

	err = global.DB.
		Preload("MenuRoles").
		Preload("MenuAction").
		Preload("MenuAction.ActionRoles").
		First(&m, id).Error
	if err != nil {
		return m, err
	}

	m.Children = s.GetTree(menus, m.ID)

	return m, nil
}

// Delete 删除
// @param id int64
// @return m model.Menu, err error
func (s *MenuService) Delete(id int64) (m model.Menu, err error) {
	var (
		detail      model.Menu
		menuActions []model.MenuAction
		actionRoles []model.ActionRoles
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
		return m, errors.New("存在子菜单，无法删除")
	}

	tx := global.DB.Begin()
	err = tx.Delete(&m, id).Error
	if err != nil {
		return m, err
	}

	// 菜单角色
	err = tx.Where("menu_id", id).Delete(&model.MenuRoles{}).Error
	if err != nil {
		return m, err
	}

	// 菜单功能
	err = tx.Where("menu_id", id).Find(&menuActions).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return m, err
		}
	}
	err = tx.Where("menu_id", id).Delete(&model.MenuAction{}).Error
	if err != nil {
		return m, err
	}

	// 功能角色
	if menuActions != nil {
		var menuActionIds []int64
		for _, v := range menuActions {
			menuActionIds = append(menuActionIds, v.ID)
		}

		err = tx.Where("action_id in ?", menuActionIds).Delete(&actionRoles).Error
		if err != nil {
			return m, err
		}
	}

	tx.Commit()

	return m, nil
}

// ActionList 功能列表
// @param menuId int64
// @return models []model.MenuAction, err error
func (s *MenuService) ActionList(menuId int64) (models []model.MenuAction, err error) {
	models, err = s.GetAll(menuId)

	if err != nil {
		return models, err
	}

	models = s.GetActionTree(models, 0)

	return models, nil
}

// GetAll 功能
// @param menuId int64
// @return models []model.MenuAction, err error
func (s *MenuService) GetAll(menuId int64) (models []model.MenuAction, err error) {
	err = global.DB.
		Preload("Parent").
		Preload("ActionRoles").
		Where("menu_id = ?", menuId).
		Order("sort asc").
		Find(&models).
		Error
	if err != nil {
		return models, err
	}

	return models, nil
}

// GetActionTree 数据子集递归
// @param rows []model.MenuAction
// @param pid int64
// @return tree []model.MenuAction
func (s *MenuService) GetActionTree(rows []model.MenuAction, pid int64) (tree []model.MenuAction) {
	for _, data := range rows {
		if data.Pid == pid {
			children := s.GetActionTree(rows, data.ID)
			data.Children = children
			tree = append(tree, data)
		}
	}
	return tree
}

// ActionCreate 功能创建
// @param m model.MenuAction
// @return model.MenuAction, error
func (s *MenuService) ActionCreate(m model.MenuAction) (model.MenuAction, error) {
	var (
		actionRoles []model.ActionRoles
	)

	tx := global.DB.Begin()
	err := tx.Create(&m).Error
	if err != nil {
		return m, err
	}

	for _, v := range m.ActionRoles {
		actionRoles = append(actionRoles, model.ActionRoles{
			ActionID: m.ID,
			RoleID:   v.RoleID,
			Name:     v.Name,
		})
	}

	if actionRoles != nil {
		err = tx.Create(&actionRoles).Error
		if err != nil {
			return m, err
		}
	}

	tx.Commit()

	return m, nil
}

// ActionUpdate 功能更新
// @param m model.MenuAction
// @return model.MenuAction, error
func (s *MenuService) ActionUpdate(m model.MenuAction) (model.MenuAction, error) {
	var (
		actionRoles []model.ActionRoles
	)

	tx := global.DB.Begin()
	err := tx.Updates(&m).Error
	if err != nil {
		return m, err
	}

	for _, v := range m.ActionRoles {
		actionRoles = append(actionRoles, model.ActionRoles{
			ActionID: m.ID,
			RoleID:   v.RoleID,
			Name:     v.Name,
		})
	}

	err = tx.Where("action_id", m.ID).Delete(&model.ActionRoles{}).Error
	if err != nil {
		return m, err
	}

	if actionRoles != nil {
		err = tx.Create(&actionRoles).Error
		if err != nil {
			return m, err
		}
	}

	tx.Commit()

	return m, nil
}

// ActionDelete 功能删除
// @param id int64
// @param menuID int64
// @return m model.MenuAction, err error
func (s *MenuService) ActionDelete(id int64, menuID int64) (m model.MenuAction, err error) {
	var (
		detail model.MenuAction
	)

	if err = global.DB.Where("id = ?", id).First(&m).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return m, err
		}
	}

	if m.MenuID != menuID {
		return m, fmt.Errorf("删除失败,该功能id【%d】不属于该菜单id【%d】", id, menuID)
	}

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

	tx := global.DB.Begin()
	err = tx.Delete(&m, id).Error
	if err != nil {
		return m, err
	}

	// 功能角色
	err = tx.Where("action_id", id).Delete(&model.ActionRoles{}).Error
	if err != nil {
		return m, err
	}

	tx.Commit()

	return m, nil
}

// ActionDetail 功能详情
// @param id int64
// @m model.MenuAction, err error
func (s *MenuService) ActionDetail(id int64) (m model.MenuAction, err error) {
	var (
		rows []model.MenuAction
	)

	err = global.DB.
		Preload("Parent").
		Preload("ActionRoles").
		First(&m, id).Error
	if err != nil {
		return m, err
	}

	rows, err = s.GetAll(m.MenuID)

	if err != nil {
		return m, err
	}

	m.Children = s.GetActionTree(rows, m.ID)

	return m, nil
}
