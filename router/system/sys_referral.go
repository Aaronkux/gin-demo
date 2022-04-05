package system

import (
	v1 "gandi.icu/demo/api/v1"
	"github.com/gin-gonic/gin"
)

type ReferralRouter struct{}

func (s *ReferralRouter) InitReferralRouter(Router *gin.RouterGroup) {
	referralRouter := Router.Group("referral")
	referralApi := v1.ApiGroupApp.SystemApiGroup.ReferralApi
	{
		referralRouter.POST("createReferral", referralApi.CreateReferral)
		referralRouter.POST("getReferralList", referralApi.GetReferralList)
		referralRouter.POST("updateReferral", referralApi.UpdateReferral)
		referralRouter.POST("deleteReferral", referralApi.DeleteReferral)
	}
}
