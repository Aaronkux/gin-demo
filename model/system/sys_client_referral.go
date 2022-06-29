package system

import "gandi.icu/demo/global"

type SysClientReferral struct {
	SysClientId   global.SnowflakeID `gorm:"column:sys_client_id"`
	SysReferralId global.SnowflakeID `gorm:"column:sys_referral_id"`
}

func (s *SysClientReferral) TableName() string {
	return "sys_client_referral"
}
