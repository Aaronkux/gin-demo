package system

import "gandi.icu/demo/global"

type SysMenu struct {
	global.CommonModel
	MenuName    string             `json:"menuName" gorm:"comment:菜单名称"`
	Path        string             `json:"path" gorm:"comment:路径"`
	Icon        string             `json:"icon" gorm:"comment:图标"`
	Hidden      bool               `json:"hidden" gorm:";comment:是否隐藏"`
	Order       uint               `json:"order" gorm:"comment:排序"`
	Authorities []SysAuthority     `json:"authorities" gorm:"many2many:sys_menu_authority;"`
	ParentId    global.SnowflakeID `json:"parentId" gorm:"comment:父级ID"`
	Children    []SysMenu          `json:"children" gorm:"foreignkey:ParentId;"`
}