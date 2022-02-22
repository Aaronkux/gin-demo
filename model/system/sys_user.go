package system

import (
	"gandi.icu/demo/global"
	"gorm.io/gorm"
)

type SysUser struct {
	global.CommonModel
	Email       string         `json:"email" gorm:"type:varchar(100);unique;comment:邮箱"`
	Phone       string         `json:"phone" gorm:"type:varchar(20);unique;comment:手机号"`
	Password    string         `json:"-" gorm:"type:varchar(100);comment:密码"`
	Nickname    string         `json:"nickname" gorm:"type:varchar(100);comment:昵称"`
	Avatar      string         `json:"avatar" gorm:"comment:头像"`
	IsActive    bool           `json:"isActive" gorm:"type:tinyint(1);default:1;comment:是否激活"`
	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
}

func (u *SysUser) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = global.SnowflakeID(global.AM_SNOWFLAKE.Generate().Int64())
	return
}
