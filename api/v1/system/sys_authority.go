package system

import (
	"fmt"

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
	if err := utils.Verify(r, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	parentId, err := r.ParentId.Int64()
	if err != nil {
		response.FailWithMessage("父级ID错误", c)
		return
	}
	authority := system.SysAuthority{AuthorityName: r.AuthorityName, ParentId: global.SnowflakeID(parentId)}
	authority.ID = global.SnowflakeID(global.AM_SNOWFLAKE.Generate().Int64())
	if authBack, err := authorityService.CreateAuthority(authority); err != nil {
		global.AM_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		fmt.Println("test")
		// _ = casbinService.UpdateCasbin(authority.AuthorityId, systemReq.DefaultCasbin())
		response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "创建成功", c)
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
	if list, total, err := authorityService.GetAuthorityInfoList(pageInfo); err != nil {
		global.AM_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
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
	var auth system.SysAuthority
	_ = c.ShouldBindJSON(&auth)
	if err := utils.Verify(auth, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 只更新名字
	updateAuth := system.SysAuthority{AuthorityName: auth.AuthorityName}
	updateAuth.ID = auth.ID
	if authority, err := authorityService.UpdateAuthority(updateAuth); err != nil {
		global.AM_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authority}, "更新成功", c)
	}
}

func (a *AuthorityApi) DeleteAuthority(c *gin.Context) {
	var authority system.SysAuthority
	_ = c.ShouldBindJSON(&authority)
	if err := utils.Verify(authority, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := authorityService.DeleteAuthority(&authority); err != nil { // 删除角色之前需要判断是否有用户正在使用此角色
		global.AM_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
