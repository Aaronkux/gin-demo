package system

import (
	"gandi.icu/demo/global"
	// "gorm.io/gorm"
	"time"

	"gorm.io/plugin/soft_delete"
)

type SysMenu struct {
	ID          global.SnowflakeID    `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
	DeletedAt   soft_delete.DeletedAt `json:"-" gorm:"uniqueIndex:idx_menu;"`
	MenuName    string                `json:"menuName" gorm:"comment:菜单名称"`
	Path        string                `json:"path" gorm:"comment:路径"`
	MenuKey     string                `json:"menuKey" gorm:"uniqueIndex:idx_menu;comment:菜单key"`
	Hidden      *bool                 `json:"hidden" gorm:";comment:是否隐藏"`
	Authorities []SysAuthority        `json:"authorities" gorm:"many2many:sys_menu_authority;"`
	ParentId    global.SnowflakeID    `json:"parentId" gorm:"comment:父级ID"`
	Children    []SysMenu             `json:"children" gorm:"foreignkey:ParentId;"`
}
