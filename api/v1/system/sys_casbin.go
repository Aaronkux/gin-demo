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

// TODO 根据authorityIds获取权限列表

type CasbinApi struct{}

// @Tags Casbin
// @Summary 更新角色api权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {object} response.Response{msg=string} "更新角色api权限"
// @Router /casbin/UpdateCasbin [post]
func (cas *CasbinApi) UpdateCasbin(c *gin.Context) {
	var cmr systemReq.CasbinInReceive
	_ = c.ShouldBindJSON(&cmr)
	if err := utils.Verify(cmr, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := casbinService.UpdateCasbin(cmr.AuthorityId.String(), cmr.CasbinInfos); err != nil {
		global.AM_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Casbin
// @Summary 获取权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {object} response.Response{data=systemRes.PolicyPathResponse,msg=string} "获取权限列表,返回包括casbin详情列表"
// @Router /casbin/getPolicyPathByAuthorityId [post]
func (cas *CasbinApi) GetPolicyPathByAuthorityId(c *gin.Context) {
	var r request.GetById
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	paths := casbinService.GetPolicyPathByAuthorityId(r.ID.String())
	response.OkWithDetailed(systemRes.PolicyPathResponse{Paths: paths}, "", c)
}
