package system

import (
	"gandi.icu/demo/global"
)

type SysSale struct {
	global.CommonModel
	Email    string             `json:"email" gorm:"comment:邮箱"`
	Name     string             `json:"name" gorm:"comment:名称"`
	IsActive bool               `json:"isActive" gorm:"comment:是否激活"`
	BranchID global.SnowflakeID `json:"branchId" gorm:"comment:branchID"`
}
