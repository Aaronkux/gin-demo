package system

import (
	"gandi.icu/demo/global"
)

type SysUser struct {
	global.CommonModel
	Email       string         `json:"email" gorm:"unique;comment:邮箱"`
	Phone       string         `json:"phone" gorm:"unique;comment:手机号"`
	Password    string         `json:"-" gorm:"comment:密码"`
	NickName    string         `json:"nickname" gorm:"comment:昵称"`
	Avatar      string         `json:"avatar" gorm:"comment:头像"`
	IsActive    bool           `json:"isActive" gorm:"type:tinyint(1);default:1;comment:是否激活"`
	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
}
