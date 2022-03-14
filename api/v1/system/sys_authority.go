package system

import (
	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/request"
	"gandi.icu/demo/model/common/response"
	"gandi.icu/demo/model/system"
	systemReq "gandi.icu/demo/model/system/request"
	systemRes "gandi.icu/demo/model/system/response"
	"gandi.icu/demo/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorityApi struct{}

func (a *AuthorityApi) CreateAuthority(c *gin.Context) {
	var r systemReq.CreateAuthority
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.AuthorityCreateVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if authorityRes, err := authorityService.CreateAuthority(r); err != nil {
		global.AM_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("创建失败", err, c)
	} else {
		// _ = casbinService.UpdateCasbin(authority.AuthorityId, systemReq.DefaultCasbin())
		response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authorityRes}, "创建成功", c)
	}
}

func (a *AuthorityApi) GetAuthorityList(c *gin.Context) {
	// 解析分页参数
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	// 验证分页参数
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 获取角色列表
	if list, total, err := authorityService.GetAuthorityList(pageInfo); err != nil {
		global.AM_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("获取失败", err, c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (a *AuthorityApi) UpdateAuthority(c *gin.Context) {
	var r system.SysAuthority
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.AuthorityUpdateVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if authority, err := authorityService.UpdateAuthority(r); err != nil {
		global.AM_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("更新失败", err, c)
	} else {
		response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authority}, "更新成功", c)
	}
}

func (a *AuthorityApi) DeleteAuthority(c *gin.Context) {
	var r system.SysAuthority
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := authorityService.DeleteAuthority(r); err != nil { // 删除角色之前需要判断是否有用户正在使用此角色
		global.AM_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("删除失败", err, c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (a *AuthorityApi) SetAuthorityMenu(c *gin.Context) {
	var r systemReq.SetAuthorityMenu
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.AuthorityMenuVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := authorityService.SetAuthorityMenu(r); err != nil {
		global.AM_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("设置失败", err, c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}
