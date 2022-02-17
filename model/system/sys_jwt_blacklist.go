package system

import (
	"gandi.icu/demo/global"
)

type JwtBlacklist struct {
	global.CommonModel
	Jwt string `gorm:"type:text;comment:jwt"`
}
