package system

import "gandi.icu/demo/global"

type SysFile struct {
	global.CommonModel
	Path     string `json:"path" gorm:"unique;comment:文件路径"`
	FileName string `json:"fileName" gorm:"comment:文件名称"`
	FileSize int64  `json:"fileSize" gorm:"comment:文件大小"`
	Bucket   string `json:"bucket" gorm:"comment:文件所在的bucket"`
}
