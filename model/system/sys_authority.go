package system

import (
	"gandi.icu/demo/global"
	"gorm.io/gorm"
)

type SysAuthority struct {
	global.CommonModel
	AuthorityName string             `json:"authorityName" gorm:"comment:权限名称"`
	ParentId      global.SnowflakeID `json:"parentId" gorm:"comment:父级ID"`
	Children      []SysAuthority     `json:"children" gorm:"foreignkey:ParentId;"`
	Users         []SysUser          `json:"users" gorm:"many2many:sys_user_authority;"`
	Menus         []SysMenu          `json:"menus" gorm:"many2many:sys_menu_authority;"`
}

func (authority *SysAuthority) BeforeCreate(tx *gorm.DB) (err error) {
	authority.ID = global.SnowflakeID(global.AM_SNOWFLAKE.Generate().Int64())
	return
}
