package system

import (
	"errors"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/system"
	systemReq "gandi.icu/demo/model/system/request"
	"gorm.io/gorm"
)

type ViewService struct{}

func (viewService *ViewService) CreateView(r system.SysView) (viewRes system.SysView, err error) {
	r.ID = global.SnowflakeID(global.AM_SNOWFLAKE.Generate().Int64())
	return r, global.AM_DB.Create(&r).Error
}

func (viewService *ViewService) UpdateView(r system.SysView) (err error) {
	if errors.Is(global.AM_DB.Where("id = ?", r.ID).First(&system.SysView{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("未找到该客户")
	}
	return global.AM_DB.Select("Value").Updates(&r).Error
}

func (viewService *ViewService) GetViewByUserIdAndType(r systemReq.SearchViewParams) (viewRes system.SysView, err error) {
	tempErr := global.AM_DB.Where("type = ? and user_id = ?", r.Type, r.UserID).First(&viewRes).Error
	if tempErr != nil {
		if errors.Is(tempErr, gorm.ErrRecordNotFound) {
			return viewRes, nil
		} else {
			return viewRes, tempErr
		}
	}
	return viewRes, nil
}

func (viewService *ViewService) GetViewById(id global.SnowflakeID) (viewRes system.SysView, err error) {
	return viewRes, global.AM_DB.Where("id = ?", id).First(&viewRes).Error
}

func (viewService *ViewService) DeleteView(id global.SnowflakeID) (err error) {
	return global.AM_DB.Where("id = ?", id).Delete(&system.SysView{}).Error
}
