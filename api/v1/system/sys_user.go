package system

import (
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

func (u *UserApi) UploadAvatar(c *gin.Context) {
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

func (u *UserApi) Register(c *gin.Context) {
	var r systemReq.Register
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if userRes, err := userService.Register(r); err != nil {
		global.AM_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("注册失败", err, c)
	} else {
		response.OkWithDetailed(systemRes.SysUserResponse{User: userRes}, "注册成功", c)
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

	if userRes, err := userService.UpdateSelf(r, userID); err != nil {
		global.AM_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("更新失败", err, c)
	} else {
		response.OkWithDetailed(systemRes.SysUserResponse{User: userRes}, "更新成功", c)
	}
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

func (m *UserApi) GetUserById(c *gin.Context) {
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

func (m *UserApi) UpdateUser(c *gin.Context) {
	var r systemReq.UpdateUser
	_ = c.ShouldBindJSON(&r)

	if err := utils.Verify(r, utils.UserUpdateVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if userRes, err := userService.UpdateUser(r); err != nil {
		global.AM_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysUserResponse{User: userRes}, "更新成功", c)
	}
}

func (m *UserApi) DeleteUser(c *gin.Context) {
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
