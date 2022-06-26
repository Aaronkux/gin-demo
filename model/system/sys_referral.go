package system

import (
	"gandi.icu/demo/global"
	"gorm.io/gorm"
)

type SysReferral struct {
	global.CommonModel
	Name   string `json:"name" gorm:"comment:名称"`
	Avatar string `json:"avatar" gorm:"comment:头像"`
}

func (referral *SysReferral) BeforeCreate(tx *gorm.DB) (err error) {
	referral.ID = global.SnowflakeID(global.AM_SNOWFLAKE.Generate().Int64())
	return
}
