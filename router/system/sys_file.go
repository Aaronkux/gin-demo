package system

import (
	v1 "gandi.icu/demo/api/v1"
	"github.com/gin-gonic/gin"
)

type FileRouter struct{}

func (f *FileRouter) InitFileRouter(Router *gin.RouterGroup) {
	fileRouter := Router.Group("file")
	fileApi := v1.ApiGroupApp.SystemApiGroup.FileApi
	{
		fileRouter.POST("authorize", fileApi.Authorize)
		fileRouter.POST("uploadAvatar", fileApi.UploadAvatar)
	}
}
