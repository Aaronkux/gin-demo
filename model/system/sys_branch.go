package system

import (
	"gandi.icu/demo/global"
	"gorm.io/gorm"
)

type SysBranch struct {
	global.CommonModel
	Name  string    `json:"name" gorm:"unique;comment:名称"`
	Sales []SysSale `json:"sales" gorm:"foreignKey:BranchID;"`
}

func (branch *SysBranch) BeforeCreate(tx *gorm.DB) (err error) {
	branch.ID = global.SnowflakeID(global.AM_SNOWFLAKE.Generate().Int64())
	return
}
