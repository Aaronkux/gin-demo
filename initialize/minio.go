package initialize

import (
	"gandi.icu/demo/global"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

func MinIO() {
	minioCfg := global.AM_CONFIG.MinIO
	minioClient, err := minio.New(minioCfg.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioCfg.AccessKeyID, minioCfg.SecretAccessKey, ""),
		Secure: minioCfg.UseSSL,
	})
	if err != nil {
		global.AM_LOG.Error("minio connect failed, err:", zap.Error(err))
	} else {
		global.AM_LOG.Info("minio connect success")
		global.AM_MinIO = minioClient
	}
}
