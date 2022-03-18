package system

import (
	"mime/multipart"
	"strings"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/response"
	"gandi.icu/demo/model/system"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FileService struct{}

func (f *FileService) UploadAvatarFile(file *multipart.FileHeader, folder string, archive string, c *gin.Context) (fileRes system.SysFile, err error) {
	dotIndex := strings.LastIndex(file.Filename, ".")
	if dotIndex == -1 || dotIndex == len(file.Filename)-1 {
		return fileRes, &response.CusError{Msg: "不支持的文件名"}
	}
	extension := file.Filename[dotIndex:]
	id := global.SnowflakeID(global.AM_SNOWFLAKE.Generate().Int64())
	newFile := system.SysFile{FileName: file.Filename, FileSize: file.Size, Archive: archive, Path: folder + id.String() + extension}
	newFile.ID = id
	err = global.AM_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&newFile).Error; err != nil {
			return err
		}
		if err = c.SaveUploadedFile(file, newFile.Path); err != nil {
			return err
		}
		return nil
	})
	return newFile, err
}
