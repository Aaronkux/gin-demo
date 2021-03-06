package system

import (
	"fmt"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/request"
	"gandi.icu/demo/model/common/response"
	systemReq "gandi.icu/demo/model/system/request"
	systemRes "gandi.icu/demo/model/system/response"
	"gandi.icu/demo/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserApi struct{}

func (u *UserApi) Register(c *gin.Context) {
	var r systemReq.Register
	_ = c.ShouldBind(&r)
	if err := utils.Verify(r, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// limit file size
	if r.Avatar.Size > 1024*1024*2 {
		response.FailWithMessage("文件大小不能超过2M", c)
		return
	}

	if err := userService.Register(c, r); err != nil {
		global.AM_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("注册失败", err, c)
	} else {
		response.OkWithMessage("注册成功", c)
	}
}

func (u *UserApi) UpdateSelf(c *gin.Context) {
	var r systemReq.UpdateSelf
	_ = c.ShouldBindJSON(&r)
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("请先登录", c)
		return
	}
	if err := utils.Verify(r, utils.UpdateSelfVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := userService.UpdateSelf(r, userID); err != nil {
		global.AM_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("更新失败", err, c)
	} else {
		response.OkWithMessage("", c)
	}
}

func (u *UserApi) UpdateSelfAvatar(c *gin.Context) {
	// 从token获取用户id
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("请先登录", c)
		return
	}
	// 获取文件
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		response.FailWithMessage("读取文件失败", c)
		return
	}
	// limit file size
	if file.Size > 1024*1024*2 {
		response.FailWithMessage("文件大小不能超过2M", c)
		return
	}

	fileRes, err := userService.UpdateSelfAvatar(c, userID, file)
	if err != nil {
		response.FailWithMessage("上传失败, 请联系管理员", c)
		global.AM_LOG.Error("上传失败!", zap.Error(err))
		return
	}
	response.OkWithDetailed(systemRes.SysFileResponse{File: fileRes}, "头像更新成功", c)
}

func (u *UserApi) GetUserList(c *gin.Context) {
	var r systemReq.SearchUserParams
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r.PageInfo, utils.GetUserListVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := userService.GetUserList(r); err != nil {
		global.AM_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     r.Page,
			PageSize: r.PageSize,
		}, "", c)
	}
}

func (user *UserApi) GetUserById(c *gin.Context) {
	var r request.GetById
	_ = c.ShouldBindJSON(&r)

	if err := utils.Verify(r, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if userRes, err := userService.GetUserById(r.ID); err != nil {
		global.AM_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysUserResponse{User: userRes}, "", c)
	}
}

func (user *UserApi) UpdateUser(c *gin.Context) {
	var r systemReq.UpdateUser
	_ = c.ShouldBindJSON(&r)

	if err := utils.Verify(r, utils.UpdateUserVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := userService.UpdateUser(r); err != nil {
		global.AM_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (u *UserApi) UpdateUserAvatar(c *gin.Context) {
	var r systemReq.UpdateUserAvatar
	_ = c.ShouldBind(&r)

	if err := utils.Verify(r, utils.UpdateUserAvatarVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if r.ID == 0 {
		response.FailWithMessage("请先登录", c)
		return
	}
	// limit file size
	if r.File.Size > 1024*1024*2 {
		response.FailWithMessage("文件大小不能超过2M", c)
		return
	}

	err := userService.UpdateUserAvatar(c, r.ID, r.File)
	if err != nil {
		response.FailWithMessage("上传失败, 请联系管理员", c)
		global.AM_LOG.Error("上传失败!", zap.Error(err))
		return
	}
	response.OkWithMessage("头像更新成功", c)
}

func (user *UserApi) DeleteUser(c *gin.Context) {
	var r request.GetById
	_ = c.ShouldBindJSON(&r)

	if err := utils.Verify(r, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := userService.DeleteUser(r.ID); err != nil {
		global.AM_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
