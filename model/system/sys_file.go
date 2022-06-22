package system

import "gandi.icu/demo/global"

type SysFile struct {
	global.CommonModel
	ObjectName  string `json:"objectName" gorm:"unique;comment:文件对象名"`
	FileName    string `json:"fileName" gorm:"comment:文件原始名称"`
	FileSize    int64  `json:"fileSize" gorm:"comment:文件大小"`
	ContentType string `json:"contentType" gorm:"comment:文件类型"`
	Bucket      string `json:"bucket" gorm:"comment:文件所在的bucket"`
}
