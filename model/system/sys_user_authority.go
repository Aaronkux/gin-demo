package system

import "gandi.icu/demo/global"

type SysUserAuthority struct {
	SysUserId      global.SnowflakeID `gorm:"column:sys_user_id"`
	SysAuthorityId global.SnowflakeID `gorm:"column:sys_authority_id"`
}

func (s *SysUserAuthority) TableName() string {
	return "sys_user_authority"
}
