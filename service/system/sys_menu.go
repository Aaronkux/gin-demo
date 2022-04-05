package system

import (
	"errors"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/response"
	"gandi.icu/demo/model/system"
	systemReq "gandi.icu/demo/model/system/request"
	"gorm.io/gorm"
)

type MenuService struct{}

func (menuService *MenuService) CreateMenu(r systemReq.CreateMenu) (menuRes system.SysMenu, err error) {
	var parentMenu system.SysMenu
	if !errors.Is(global.AM_DB.Where("menu_key = ?", r.MenuKey).First(&menuRes).Error, gorm.ErrRecordNotFound) {
		return menuRes, &response.CusError{Msg: "该菜单key已存在"}
	}
	if *r.ParentId != 0 && errors.Is(global.AM_DB.Where("id = ?", r.ParentId).First(&parentMenu).Error, gorm.ErrRecordNotFound) {
		return menuRes, &response.CusError{Msg: "父级菜单不存在"}
	}

	newMenu := system.SysMenu{MenuName: r.MenuName, Path: r.Path, Hidden: *r.Hidden, ParentId: *r.ParentId, MenuKey: r.MenuKey}
	newMenu.ID = global.SnowflakeID(global.AM_SNOWFLAKE.Generate().Int64())

	err = global.AM_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&newMenu).Error; err != nil {
			return err
		}
		if *r.ParentId != 0 {
			if err := tx.Model(&parentMenu).Association("Children").Append(&newMenu); err != nil {
				return err
			}
		}
		return nil
	})
	return newMenu, err
}

func (menuService *MenuService) GetMenuList(r systemReq.SearchMenuParams) (list interface{}, total int64, err error) {
	var menuList []system.SysMenu
	limit := r.PageSize
	offset := r.PageSize * (r.Page - 1)
	db := global.AM_DB.Model(&system.SysMenu{})
	if r.MenuName != "" {
		db = db.Where("menu_name like ?", "%"+r.MenuName+"%")
	}
	err = db.Where("parent_id = ?", 0).Count(&total).Error
	if err != nil {
		return menuList, total, err
	}
	err = db.Limit(limit).Offset(offset).Where("parent_id = ?", 0).Find(&menuList).Error
	if len(menuList) > 0 {
		for k := range menuList {
			err = menuService.findChildrenMenu(&menuList[k])
		}
	}
	return menuList, total, err
}

func (menuService *MenuService) GetMenuListAll() (list interface{}, total int64, err error) {
	var menuList []system.SysMenu

	db := global.AM_DB.Model(&system.SysMenu{})

	err = db.Where("parent_id = ?", 0).Count(&total).Error
	if err != nil {
		return menuList, total, err
	}
	err = db.Where("parent_id = ?", 0).Find(&menuList).Error
	if len(menuList) > 0 {
		for k := range menuList {
			err = menuService.findChildrenMenu(&menuList[k])
		}
	}
	return menuList, total, err
}

func (menuService *MenuService) findChildrenMenu(menu *system.SysMenu) (err error) {
	err = global.AM_DB.Where("parent_id = ?", menu.ID).Find(&menu.Children).Error
	if len(menu.Children) > 0 {
		for k := range menu.Children {
			err = menuService.findChildrenMenu(&menu.Children[k])
		}
	}
	return err
}

func (menuService *MenuService) GetMenuById(r system.SysMenu) (menuRes system.SysMenu, err error) {
	err = global.AM_DB.Where("id = ?", r.ID).First(&menuRes).Error
	return menuRes, err
}

func (menuService *MenuService) UpdateMenu(r systemReq.UpdateMenu) (err error) {
	var menuExist system.SysMenu
	updateMenu := system.SysMenu{MenuName: r.MenuName, Path: r.Path, Hidden: *r.Hidden, MenuKey: r.MenuKey}
	err = global.AM_DB.Where("id = ?", r.ID).First(&menuExist).Select("MenuName", "Path", "Hidden", "MenuKey").Updates(&updateMenu).Error
	return err
}

func (menuService *MenuService) DeleteMenu(r system.SysMenu) (err error) {
	if errors.Is(global.AM_DB.Preload("Authorities").Preload("Children").First(&r).Error, gorm.ErrRecordNotFound) {
		return &response.CusError{Msg: "该菜单不存在"}
	}
	if len(r.Children) != 0 {
		return &response.CusError{Msg: "此菜单有子菜单正在使用，禁止删除"}
	}
	err = global.AM_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&r).Error; err != nil {
			return err
		}
		if len(r.Authorities) != 0 {
			if err := tx.Model(&r).Association("Authorities").Clear(); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func (menuService *MenuService) GetMenuKeysByUserAuthority(ids []string) (keys []string, err error) {
	var allRelations []system.SysMenuAuthority
	if err = global.AM_DB.Where("sys_authority_id in ?", ids).Find(&allRelations).Error; err != nil {
		return keys, err
	}
	var menuIds []global.SnowflakeID
	for _, v := range allRelations {
		menuIds = append(menuIds, v.SysMenuId)
	}
	var menus []system.SysMenu
	if err = global.AM_DB.Where("id in ? and hidden = ?", menuIds, false).Find(&menus).Error; err != nil {
		return keys, err
	}
	for _, v := range menus {
		keys = append(keys, v.MenuKey)
	}
	return keys, err
}
