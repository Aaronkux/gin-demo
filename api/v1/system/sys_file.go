package system

import (
	"net/http"
	"path/filepath"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/response"
	systemReq "gandi.icu/demo/model/system/request"
	"github.com/gin-gonic/gin"
)

type FileApi struct{}

func (f *FileApi) Authorize(c *gin.Context) {
	var r systemReq.FileAuthorize
	_ = c.ShouldBind(&r)
	println(r.Code, r.AuthToken, r.FileName, r.Scene, filepath.Clean(r.Path))
	c.String(http.StatusOK, "ok")
}

func (f *FileApi) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage("读取头像文件失败", c)
		return
	}
	fileRes, err := fileService.UploadAvatarFile(file, global.AM_CONFIG.Local.Avatar, "avatar", c)
	if err != nil {
		response.FailWithMessage("上传头像失败, 请联系管理员", c)
		return
	}
	response.OkWithDetailed(gin.H{"filePath": fileRes.Path}, "上传头像成功", c)
}
