package system

import (
	"gandi.icu/demo/global"
)

type SysSale struct {
	global.CommonModel
	Email    string             `json:"email" gorm:"unique;comment:邮箱"`
	Name     string             `json:"name" gorm:"comment:名称"`
	AvatarID global.SnowflakeID `json:"-" gorm:"comment:avatar ID"`
	Avatar   SysFile            `json:"avatar" gorm:"foreignkey:AvatarID"`
	IsActive bool               `json:"isActive" gorm:"comment:是否在职"`
	BranchID global.SnowflakeID `json:"branchId" gorm:"comment:branchID"`
	Branch   SysBranch          `json:"branch" gorm:"foreignkey:BranchID"`
}
