package system

import (
	"gandi.icu/demo/global"
	"gorm.io/gorm"
)

type JwtBlacklist struct {
	global.CommonModel
	Jwt string `gorm:"type:text;comment:jwt"`
}

func (jwtBlackList *JwtBlacklist) BeforeCreate(tx *gorm.DB) (err error) {
	jwtBlackList.ID = global.SnowflakeID(global.AM_SNOWFLAKE.Generate().Int64())
	return
}
