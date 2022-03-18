package system

import "gandi.icu/demo/global"

type SysFile struct {
	global.CommonModel
	Path     string `json:"path" gorm:"unique;comment:文件路径"`
	Archive  string `json:"archive" gorm:"comment:文件名"`
	FileName string `json:"fileName" gorm:"comment:文件名称"`
	FileSize int64  `json:"fileSize" gorm:"comment:文件大小"`
}
