package system

import (
	"gandi.icu/demo/global"
	"gorm.io/gorm"
)

type SysApi struct {
	global.CommonModel
	Path        string `json:"path" gorm:"comment:api路径"`             // api路径
	Description string `json:"description" gorm:"comment:api中文描述"`    // api中文描述
	ApiGroup    string `json:"apiGroup" gorm:"comment:api组"`          // api组
	Method      string `json:"method" gorm:"default:POST;comment:方法"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
}

func (api *SysApi) BeforeCreate(tx *gorm.DB) (err error) {
	api.ID = global.SnowflakeID(global.AM_SNOWFLAKE.Generate().Int64())
	return
}
