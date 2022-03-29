package system

import (
	"time"

	"gandi.icu/demo/global"
	"gorm.io/plugin/soft_delete"
)

type SysUser struct {
	ID          global.SnowflakeID    `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
	DeletedAt   soft_delete.DeletedAt `json:"-" gorm:"uniqueIndex:idx_user;"`
	Email       string                `json:"email" gorm:"uniqueIndex:idx_user;comment:邮箱"`
	Phone       string                `json:"phone" gorm:"comment:手机号"`
	Password    string                `json:"-" gorm:"comment:密码"`
	NickName    string                `json:"nickname" gorm:"comment:昵称"`
	Avatar      string                `json:"avatar" gorm:"comment:头像"`
	IsActive    bool                  `json:"isActive" gorm:"type:tinyint(1);default:1;comment:是否激活"`
	Authorities []SysAuthority        `json:"authorities" gorm:"many2many:sys_user_authority;"`
}
