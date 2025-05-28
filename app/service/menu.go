package service

import (
	"errors"
	"fmt"
	"gin-base/app/model"
	"gin-base/common"
	"gin-base/common/global"
	"gin-base/helper/utils"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type MenuService struct {
	common.BaseService
}

// List 列表
// @return bool
func (s *MenuService) List() (menus []model.MenuQuery, err error) {
	var (
		menuModels []model.Menu
	)

	err = global.DB.
		Preload("MenuRoles").
		Preload("MenuAction", func(db *gorm.DB) *gorm.DB {
			return db.Preload("ActionRoles").
				Order("sort asc").
				Select("id,menu_id,type,name,is_link,sort,created_at,updated_at")
		}).
		Preload("MenuAction.ActionRoles").
		Order("sort asc").
		Find(&menuModels).Error

	if err != nil {
		return menus, err
	}
	err = copier.Copy(&menus, &menuModels)
	if err != nil {
		return menus, err
	}

	for k, m := range menuModels {
		menus[k].IsLink = m.GetIsLink()
		menus[k].Meta = m.GetMeta()
		menus[k].MenuRoles = m.MenuRoles
		menus[k].MenuAction = m.MenuAction
	}

	menus = s.MenuTree(menus, 0)

	return menus, nil
}

// MenuTree 数据子集递归
// @param menus []model.MenuQuery
// @param pid int64
// @return tree []model.MenuQuery
func (s *MenuService) MenuTree(menus []model.MenuQuery, pid int64) (tree []model.MenuQuery) {
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
// @param req model.MenuQuery
// @return menuModel model.Menu, err error
func (s *MenuService) Create(req model.MenuQuery) (menuModel model.Menu, err error) {
	var (
		menuRoles []model.MenuRoles
	)

	err = copier.Copy(&menuModel, &req)
	if err != nil {
		return menuModel, err
	}

	menuModel.Meta = menuModel.SetMeta(req.Meta)
	menuModel.IsLink = menuModel.SetIsLink(req.IsLink)

	tx := global.DB.Begin()
	err = tx.Create(&menuModel).Error
	if err != nil {
		return menuModel, err
	}

	for _, v := range req.Meta.Roles {
		menuRoles = append(menuRoles, model.MenuRoles{
			MenuID: menuModel.ID,
			RoleID: v,
		})
	}

	err = tx.Create(&menuRoles).Error
	if err != nil {
		return menuModel, err
	}
	tx.Commit()

	return menuModel, nil
}

// Update 更新
// @param req model.Menu
// @return menuModel model.Menu, err error
func (s *MenuService) Update(req model.MenuQuery) (menuModel model.Menu, err error) {
	var (
		menuRole  model.MenuRoles
		menuRoles []model.MenuRoles
	)

	for _, v := range req.Meta.Roles {
		menuRoles = append(menuRoles, model.MenuRoles{
			MenuID: req.ID,
			RoleID: v,
		})
	}

	err = copier.Copy(&menuModel, &req)
	if err != nil {
		return menuModel, err
	}

	menuModel.Meta = menuModel.SetMeta(req.Meta)
	menuModel.IsLink = menuModel.SetIsLink(req.IsLink)

	tx := global.DB.Begin()
	err = tx.Updates(&menuModel).Error
	if err != nil {
		return menuModel, err
	}

	err = tx.Where("menu_id", req.ID).Delete(&menuRole).Error
	if err != nil {
		return menuModel, err
	}

	err = tx.Create(&menuRoles).Error
	if err != nil {
		return menuModel, err
	}
	tx.Commit()

	return menuModel, nil
}

// Delete 删除
// @param id int64
// @return menuModel model.Menu, err error
func (s *MenuService) Delete(id int64) (menuModel model.Menu, err error) {
	err = global.DB.Delete(&menuModel, id).Error
	if err != nil {
		return menuModel, err
	}

	return menuModel, nil
}

// ActionList 功能列表
// @param menuId int64
// @return actionQuery []model.MenuActionQuery, err error
func (s *MenuService) ActionList(menuId int64) (actionQuery []model.MenuActionQuery, err error) {
	var (
		menuActionModels []model.MenuAction
	)

	err = global.DB.
		Where("menu_id = ?", menuId).
		Order("sort asc").
		Find(&menuActionModels).
		Scan(&actionQuery).
		Error
	if err != nil {
		return actionQuery, err
	}

	return actionQuery, nil
}

// ActionCreate 功能创建
// @param req model.MenuActionQuery
// @return menuActionModel model.MenuAction, err error
func (s *MenuService) ActionCreate(req model.MenuActionQuery) (menuActionModel model.MenuAction, err error) {
	err = copier.Copy(&menuActionModel, &req)
	if err != nil {
		return menuActionModel, err
	}

	db := global.DB.Create(&menuActionModel)

	err = db.Error
	if err != nil {
		return menuActionModel, err
	}

	return menuActionModel, nil
}

// ActionUpdate 功能更新
// @param req model.MenuActionQuery
// @return menuActionModel model.MenuAction, err error
func (s *MenuService) ActionUpdate(req model.MenuActionQuery) (menuActionModel model.MenuAction, err error) {
	err = copier.Copy(&menuActionModel, &req)
	if err != nil {
		return menuActionModel, err
	}

	err = global.DB.Updates(&menuActionModel).Error
	if err != nil {
		return menuActionModel, err
	}

	return menuActionModel, nil
}

// ActionDelete 功能删除
// @param id int64
// @return menuActionModel model.MenuAction, err error
func (s *MenuService) ActionDelete(id int64, menuID int64) (menuActionModel model.MenuAction, err error) {
	if err = global.DB.Where("id = ?", id).First(&menuActionModel).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return menuActionModel, err
		}
	}

	if menuActionModel.MenuID != menuID {
		return menuActionModel, fmt.Errorf("删除失败,该功能id【%d】不属于该菜单id【%d】", id, menuID)
	}

	err = global.DB.Delete(&menuActionModel, id).Error
	if err != nil {
		return menuActionModel, err
	}

	return menuActionModel, nil
}
