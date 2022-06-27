package system

import (
	"gandi.icu/demo/global"
)

type SysReferral struct {
	global.CommonModel
	Name   string `json:"name" gorm:"comment:名称"`
	Avatar string `json:"avatar" gorm:"comment:头像"`
}
