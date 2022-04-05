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

type ReferralApi struct{}

func (referral *ReferralApi) CreateReferral(c *gin.Context) {
	var r systemReq.CreateReferral
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.ReferralCreateVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if referralRes, err := referralService.CreateReferral(r); err != nil {
		global.AM_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("创建失败", err, c)
	} else {
		response.OkWithDetailed(systemRes.SysReferralResponse{Referral: referralRes}, "创建成功", c)
	}
}

func (referral *ReferralApi) GetReferralList(c *gin.Context) {
	var r systemReq.SearchReferralParams
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r.PageInfo, utils.GetReferralListVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := referralService.GetReferralList(r); err != nil {
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

func (referral *ReferralApi) UpdateReferral(c *gin.Context) {
	var r systemReq.UpdateReferral
	_ = c.ShouldBindJSON(&r)

	if err := utils.Verify(r, utils.ReferralUpdateVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := referralService.UpdateReferral(r); err != nil {
		global.AM_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (referral *ReferralApi) DeleteReferral(c *gin.Context) {
	var r request.GetById
	_ = c.ShouldBindJSON(&r)

	if err := utils.Verify(r, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := referralService.DeleteReferral(r.ID); err != nil {
		global.AM_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
