package system

import (
	"gandi.icu/demo/global"
	"gorm.io/gorm"
)

type SysAuthority struct {
	global.CommonModel
	AuthorityId   string         `json:"authorityId" gorm:"unique;comment:权限id"`
	AuthorityName string         `json:"authorityName" gorm:"comment:权限名称"`
	ParentId      *int64         `json:"parentId" gorm:"comment:父级ID"`
	Children      []SysAuthority `json:"children" gorm:"foreignkey:ParentId;"`
	Users         []SysUser      `json:"users" gorm:"many2many:sys_user_authority;"`
}

func (u *SysAuthority) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = global.SnowflakeID(global.AM_SNOWFLAKE.Generate().Int64())
	return
}
