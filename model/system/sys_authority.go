package system

import (
	"gandi.icu/demo/global"
)

type SysAuthority struct {
	global.CommonModel
	AuthorityName string             `json:"authorityName" gorm:"comment:权限名称"`
	ParentId      global.SnowflakeID `json:"parentId" gorm:"comment:父级ID"`
	Children      []SysAuthority     `json:"children" gorm:"foreignkey:ParentId;"`
	Users         []SysUser          `json:"users" gorm:"many2many:sys_user_authority;"`
}
