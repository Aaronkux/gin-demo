package system

import (
	"gandi.icu/demo/global"
)

type SysBranch struct {
	global.CommonModel
	Name  string    `json:"name" gorm:"unique;comment:名称"`
	Sales []SysSale `json:"sales" gorm:"foreignKey:BranchID;"`
}
