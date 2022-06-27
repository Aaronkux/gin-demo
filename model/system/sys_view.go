package system

import (
	"gandi.icu/demo/global"
)

type SysView struct {
	global.CommonModel
	Type   string             `json:"type" gorm:"uniqueIndex:idx_view;comment:视图类型"`
	Value  string             `json:"value" gorm:"type:text;comment:视图值"`
	UserID global.SnowflakeID `json:"userId" gorm:"uniqueIndex:idx_view;comment:用户ID"`
	User   SysUser            `json:"user" gorm:"foreignkey:UserID"`
}
