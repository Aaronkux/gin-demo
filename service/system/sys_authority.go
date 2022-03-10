package system

import (
	"errors"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/request"
	"gandi.icu/demo/model/system"
	systemReq "gandi.icu/demo/model/system/request"
	"gorm.io/gorm"
)

type AuthorityService struct{}

var AuthorityServiceApp = new(AuthorityService)

func (authorityService *AuthorityService) CreateAuthority(r systemReq.CreateAuthority) (authorityRes system.SysAuthority, err error) {
	var parentAuthority system.SysAuthority
	if r.ParentId != 0 && errors.Is(global.AM_DB.Where("id = ?", r.ParentId).First(&parentAuthority).Error, gorm.ErrRecordNotFound) {
		return authorityRes, errors.New("父级角色不存在")
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
