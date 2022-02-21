package system

import (
	"gandi.icu/demo/global"
)

type SysUser struct {
	global.CommonModel
	Email    string `json:"email" gorm:"type:varchar(100);unique;comment:邮箱"`
	Phone    string `json:"phone" gorm:"type:varchar(20);unique;comment:手机号"`
	Password string `json:"-" gorm:"type:varchar(100);comment:密码"`
	Nickname string `json:"nickname" gorm:"type:varchar(100);comment:昵称"`
	Avatar   string `json:"avatar" gorm:"comment:头像"`
	IsActive bool   `json:"is_active" gorm:"type:tinyint(1);default:1;comment:是否激活"`
}
