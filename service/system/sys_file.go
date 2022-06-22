package system

import (
	"mime/multipart"
	"net/http"
	"path/filepath"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/system"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type FileService struct{}

func (f *FileService) UploadFile(c *gin.Context, folder string, file *multipart.FileHeader) (fileRes system.SysFile, err error) {
	reader, err := file.Open()
	if err != nil {
		return fileRes, err
	}
	defer reader.Close()
	buf := make([]byte, 512)
	_, err = reader.Read(buf)
	if err != nil {
		return fileRes, err
	}
	contentType := http.DetectContentType(buf)
	if _, err := reader.Seek(0, 0); err != nil {
		return fileRes, err
	}
	convertFileName := uuid.New().String() + filepath.Ext(file.Filename)
	info, err := global.AM_MinIO.PutObject(c, global.AM_CONFIG.MinIO.BucketName, folder+"/"+convertFileName, reader, file.Size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return fileRes, err
	}

	// save to db
	newFile := system.SysFile{ObjectName: info.Key, FileName: file.Filename, FileSize: info.Size, Bucket: global.AM_CONFIG.MinIO.BucketName, ContentType: contentType}
	newFile.ID = global.SnowflakeID(global.AM_SNOWFLAKE.Generate().Int64())
	if err := global.AM_DB.Create(&newFile).Error; err != nil {
		return fileRes, err
	}
	return newFile, nil
}

func (f *FileService) DownloadFileById(c *gin.Context, id global.SnowflakeID) (err error) {
	// get file from db
	var file system.SysFile
	if err := global.AM_DB.Where("id = ?", id).First(&file).Error; err != nil {
		return err
	}
	// get file from minio
	reader, err := global.AM_MinIO.GetObject(c, global.AM_CONFIG.MinIO.BucketName, file.ObjectName, minio.GetObjectOptions{})
	if err != nil {
		return err
	}
	defer reader.Close()
	c.DataFromReader(http.StatusOK, file.FileSize, file.ContentType, reader, map[string]string{
		"Content-Disposition": "attachment; filename=\"" + file.FileName + "\"",
	})
	return nil
}
