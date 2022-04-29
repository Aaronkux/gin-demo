package system

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"gandi.icu/demo/global"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type FileService struct{}

func (f *FileService) UploadFile(c *gin.Context, folder string, file *multipart.FileHeader) (err error) {
	reader, err := file.Open()
	if err != nil {
		return err
	}
	defer reader.Close()
	buf := make([]byte, 512)
	_, err = reader.Read(buf)
	if err != nil {
		return err
	}
	contentType := http.DetectContentType(buf)
	if _, err := reader.Seek(0, 0); err != nil {
		return err
	}
	convertFileName := uuid.New().String() + filepath.Ext(file.Filename)
	info, err := global.AM_MinIO.PutObject(c, global.AM_CONFIG.MinIO.BucketName, folder+"/"+convertFileName, reader, file.Size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}

	fmt.Println(info.Bucket, info.Key, info.Size)
	return nil
}
