package system

import (
	"gandi.icu/demo/global"
)

type SysReferral struct {
	global.CommonModel
	Name    string      `json:"name" gorm:"comment:名称"`
	Clients []SysClient `json:"clients" gorm:"many2many:sys_client_referral;"`
}
