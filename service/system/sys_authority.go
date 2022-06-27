package system

import (
	"errors"
	"fmt"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/response"
	"gandi.icu/demo/model/system"
	systemReq "gandi.icu/demo/model/system/request"
	"gorm.io/gorm"
)

type AuthorityService struct{}

func (authorityService *AuthorityService) CreateAuthority(r systemReq.CreateAuthority) (authorityRes system.SysAuthority, err error) {
	var parentAuthority system.SysAuthority
	if r.ParentId != 0 && errors.Is(global.AM_DB.Where("id = ?", r.ParentId).First(&parentAuthority).Error, gorm.ErrRecordNotFound) {
		return authorityRes, &response.CusError{Msg: "父级角色不存在"}
	}

	newAuthority := system.SysAuthority{AuthorityName: r.AuthorityName}
	newAuthority.ID = global.SnowflakeID(global.AM_SNOWFLAKE.Generate().Int64())
	err = global.AM_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&newAuthority).Error; err != nil {
			return err
		}
		if r.ParentId != 0 {
			if err := tx.Model(&parentAuthority).Association("Children").Append(&newAuthority); err != nil {
				return err
			}
		}
		return nil
	})
	return newAuthority, err
}

func (authorityService *AuthorityService) GetAuthorityList(r systemReq.SearchAuthorityParams) (list interface{}, total int64, err error) {
	var authority []system.SysAuthority
	limit := r.PageSize
	offset := r.PageSize * (r.Page - 1)
	db := global.AM_DB.Model(&system.SysAuthority{})
	if r.AuthorityName != "" {
		db = db.Where("authority_name like ?", "%"+r.AuthorityName+"%")
	}
	err = db.Where("parent_id = ?", 0).Count(&total).Error
	if err != nil {
		return authority, total, err
	}
	err = db.Limit(limit).Offset(offset).Where("parent_id = ?", 0).Find(&authority).Error
	if len(authority) > 0 {
		for k := range authority {
			err = authorityService.findChildrenAuthority(&authority[k])
		}
	}
	return authority, total, err
}

func (authorityService *AuthorityService) GetAuthorityById(id global.SnowflakeID) (authorityRes system.SysAuthority, err error) {
	var authority system.SysAuthority
	err = global.AM_DB.Where("id = ?", id).First(&authority).Error
	return authority, err
}

func (authorityService *AuthorityService) UpdateAuthority(r system.SysAuthority) (err error) {
	var authorityExist system.SysAuthority
	updateAuthority := system.SysAuthority{AuthorityName: r.AuthorityName}
	err = global.AM_DB.Where("id = ?", r.ID).First(&authorityExist).Select("AuthorityName").Updates(&updateAuthority).Error
	return err
}

func (authorityService *AuthorityService) DeleteAuthority(r system.SysAuthority) (err error) {
	if errors.Is(global.AM_DB.Preload("Users").Preload("Menus").Preload("Children").First(&r).Error, gorm.ErrRecordNotFound) {
		return &response.CusError{Msg: "该角色不存在"}
	}
	if len(r.Users) != 0 {
		return &response.CusError{Msg: "此角色有用户正在使用，禁止删除"}
	}
	if len(r.Children) != 0 {
		return &response.CusError{Msg: "此角色有子角色正在使用，禁止删除"}
	}

	err = global.AM_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&r).Error; err != nil {
			return err
		}

		if len(r.Menus) != 0 {
			if err := tx.Model(&r).Association("Menus").Clear(); err != nil {
				return err
			}
		}
		return nil
	})
	CasbinServiceApp.ClearCasbin(0, r.ID.String())
	return err
}

func (authorityService *AuthorityService) findChildrenAuthority(authority *system.SysAuthority) (err error) {
	err = global.AM_DB.Where("parent_id = ?", authority.ID).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = authorityService.findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}

func (authorityService *AuthorityService) GetAuthorityMenu(id global.SnowflakeID) (list []string, err error) {
	var authority system.SysAuthority
	if errors.Is(global.AM_DB.Where("id = ?", id).Preload("Menus").First(&authority).Error, gorm.ErrRecordNotFound) {
		return list, &response.CusError{Msg: "该角色不存在"}
	}
	for _, v := range authority.Menus {
		fmt.Println(v.MenuName)
		list = append(list, v.ID.String())
	}
	return list, nil
}
func (authorityService *AuthorityService) SetAuthorityMenu(r systemReq.SetAuthorityMenu) (err error) {

	var menus []system.SysMenu
	if err := global.AM_DB.Where("id In ?", r.MenuIds).Find(&menus).Error; err != nil {
		return err
	}
	if len(menus) != len(r.MenuIds) {
		return &response.CusError{Msg: "菜单ID不存在"}
	}

	var authority system.SysAuthority
	if errors.Is(global.AM_DB.Where("id = ?", r.AuthorityId).Preload("Menus").First(&authority).Error, gorm.ErrRecordNotFound) {
		return &response.CusError{Msg: "该角色不存在"}
	}
	global.AM_DB.Preload("Menus").First(&authority, "id = ?", r.AuthorityId)
	err = global.AM_DB.Model(&authority).Association("Menus").Replace(&menus)
	return err
}
