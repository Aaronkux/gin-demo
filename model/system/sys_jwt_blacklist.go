package system

import (
	"gandi.icu/demo/global"
	"gorm.io/gorm"
)

type JwtBlacklist struct {
	global.CommonModel
	Jwt string `gorm:"type:text;comment:jwt"`
}

func (u *JwtBlacklist) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = global.SnowflakeID(global.AM_SNOWFLAKE.Generate().Int64())
	return
}
