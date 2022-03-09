package system

import (
	"errors"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/request"
	"gandi.icu/demo/model/system"
	"gorm.io/gorm"
)

type AuthorityService struct{}

var AuthorityServiceApp = new(AuthorityService)

func (authorityService *AuthorityService) CreateAuthority(auth system.SysAuthority) (authority system.SysAuthority, err error) {
	var authorityBox system.SysAuthority
	if !errors.Is(global.AM_DB.Where("id = ?", auth.ID).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return auth, errors.New("存在相同角色id")
	}
	if auth.ParentId != 0 && errors.Is(global.AM_DB.Where("id = ?", auth.ParentId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return auth, errors.New("父级角色不存在")
	}
	createdAuthority := system.SysAuthority{AuthorityName: auth.AuthorityName}
	err = global.AM_DB.Create(&createdAuthority).Error
	if err == nil && auth.ParentId != 0 {
		err = global.AM_DB.Model(&authorityBox).Association("Children").Append(&createdAuthority)
	}
	return createdAuthority, err
}

func (authorityService *AuthorityService) GetAuthorityInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	var authority []system.SysAuthority
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.AM_DB.Model(&system.SysAuthority{})
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

func (authorityService *AuthorityService) UpdateAuthority(auth system.SysAuthority) (authority system.SysAuthority, err error) {
	err = global.AM_DB.Where("id = ?", auth.ID).First(&system.SysAuthority{}).Updates(&auth).Error
	return auth, err
}

func (authorityService *AuthorityService) DeleteAuthority(auth *system.SysAuthority) (err error) {
	if errors.Is(global.AM_DB.Debug().Preload("Users").Preload("Children").First(&auth).Error, gorm.ErrRecordNotFound) {
		return errors.New("该角色不存在")
	}
	if len(auth.Users) != 0 {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if len(auth.Children) != 0 {
		return errors.New("此角色有子角色正在使用禁止删除")
	}
	// if !errors.Is(global.AM_DB.Where("id = ?", auth.ID).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
	// 	return errors.New("此角色有用户正在使用禁止删除")
	// }
	// if !errors.Is(global.AM_DB.Where("parent_id = ?", auth.ID).First(&system.SysAuthority{}).Error, gorm.ErrRecordNotFound) {
	// 	return errors.New("此角色存在子角色不允许删除")
	// }
	err = global.AM_DB.Unscoped().Delete(auth).Error
	if err != nil {
		return
	}

	err = global.AM_DB.Delete(&[]system.SysUserAuthority{}, "sys_authority_id = ?", auth.ID).Error
	// CasbinServiceApp.ClearCasbin(0, string(auth.ID))
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
