package system

import (
	"fmt"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/request"
	"gandi.icu/demo/model/common/response"
	systemRes "gandi.icu/demo/model/system/response"
	"gandi.icu/demo/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileApi struct{}

func (f *FileApi) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		response.FailWithMessage("读取头像文件失败", c)
		return
	}

	// limit file size
	if file.Size > 1024*1024*2 {
		response.FailWithMessage("文件大小不能超过2M", c)
		return
	}

	fileRes, err := fileService.UploadFile(c, "avatar", file)
	if err != nil {
		response.FailWithMessage("上传失败, 请联系管理员", c)
		global.AM_LOG.Error("上传失败!", zap.Error(err))
		return
	}
	response.OkWithDetailed(systemRes.SysFileResponse{File: fileRes}, "", c)
}

func (f *FileApi) DownloadFileById(c *gin.Context) {
	var r request.GetById
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := fileService.DownloadFileById(c, r.ID); err != nil {
		response.FailWithMessage("下载失败", c)
		global.AM_LOG.Error("下载失败!", zap.Error(err))
		return
	}
}
