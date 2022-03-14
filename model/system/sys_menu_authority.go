package system

import "gandi.icu/demo/global"

type SysMenuAuthority struct {
	SysMenuId      global.SnowflakeID `gorm:"column:sys_menu_id"`
	SysAuthorityId global.SnowflakeID `gorm:"column:sys_authority_id"`
}

func (s *SysMenuAuthority) TableName() string {
	return "sys_user_authority"
}
