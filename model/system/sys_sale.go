package system

import (
	"gandi.icu/demo/global"
)

type SysSale struct {
	global.CommonModel
	Email    string             `json:"email" gorm:"unique;comment:邮箱"`
	Name     string             `json:"name" gorm:"comment:名称"`
	Avatar   string             `json:"avatar" gorm:"comment:头像"`
	IsActive bool               `json:"isActive" gorm:"comment:是否在职"`
	BranchID global.SnowflakeID `json:"branchId" gorm:"comment:branchID"`
}
