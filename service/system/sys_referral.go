package system

import (
	"errors"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/response"
	"gandi.icu/demo/model/system"
	systemReq "gandi.icu/demo/model/system/request"
	"gorm.io/gorm"
)

type ReferralService struct{}

func (referral *ReferralService) CreateReferral(r systemReq.CreateReferral) (referralRes system.SysReferral, err error) {
	var referralExist system.SysReferral
	if !errors.Is(global.AM_DB.Where("name = ?", r.Name).First(&referralExist).Error, gorm.ErrRecordNotFound) {
		return referralRes, &response.CusError{Msg: "已存同名Referral"}
	}

	newReferral := system.SysReferral{Name: r.Name, Avatar: r.Avatar}
	if err := global.AM_DB.Create(&newReferral).Error; err != nil {
		return referralRes, err
	}
	return newReferral, err
}

func (referral *ReferralService) UpdateReferral(r systemReq.UpdateReferral) (err error) {
	var referralExist system.SysReferral
	if errors.Is(global.AM_DB.Where("id = ?", r.ID).First(&referralExist).Error, gorm.ErrRecordNotFound) {
		return &response.CusError{Msg: "该Referral不存在"}
	}
	updateReferral := system.SysReferral{Name: r.Name, Avatar: r.Avatar}
	err = global.AM_DB.Model(&referralExist).Updates(&updateReferral).Error
	return err
}

func (referral *ReferralService) GetReferralList(r systemReq.SearchReferralParams) (list interface{}, total int64, err error) {
	var referralList []system.SysReferral
	limit := r.PageSize
	offset := r.PageSize * (r.Page - 1)
	db := global.AM_DB.Model(&system.SysReferral{})
	if r.Name != "" {
		db = db.Where("name like ?", "%"+r.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return referralList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&referralList).Error
	return referralList, total, err
}

func (referral *ReferralService) DeleteReferral(id global.SnowflakeID) (err error) {
	err = global.AM_DB.Delete(&system.SysReferral{}, id).Error
	return err
}
