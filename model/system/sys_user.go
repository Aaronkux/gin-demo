package system

import (
	"time"

	"gandi.icu/demo/global"
	"gorm.io/plugin/soft_delete"
)

type SysUser struct {
	ID          global.SnowflakeID    `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time             `json:"createdAt"`
	UpdatedAt   time.Time             `json:"updatedAt"`
	DeletedAt   soft_delete.DeletedAt `json:"-" gorm:"uniqueIndex:idx_user;"`
	Email       string                `json:"email" gorm:"uniqueIndex:idx_user;comment:邮箱"`
	Phone       string                `json:"phone" gorm:"comment:手机号"`
	Password    string                `json:"-" gorm:"comment:密码"`
	NickName    string                `json:"nickname" gorm:"comment:昵称"`
	AvatarID    global.SnowflakeID    `json:"-" gorm:"comment:avatar ID"`
	Avatar      SysFile               `json:"avatar" gorm:"foreignkey:AvatarID"`
	IsActive    bool                  `json:"isActive" gorm:"type:tinyint(1);default:1;comment:是否激活"`
	Authorities []SysAuthority        `json:"authorities" gorm:"many2many:sys_user_authority;"`
}
