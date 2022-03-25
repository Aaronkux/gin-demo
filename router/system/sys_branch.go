package system

import (
	v1 "gandi.icu/demo/api/v1"
	"github.com/gin-gonic/gin"
)

type BranchRouter struct{}

func (s *BranchRouter) InitBranchRouter(Router *gin.RouterGroup) {
	casbinRouter := Router.Group("branch")
	casbinApi := v1.ApiGroupApp.SystemApiGroup.BranchApi
	{
		casbinRouter.POST("createBranch", casbinApi.CreateBranch)
		casbinRouter.POST("getBranchList", casbinApi.GetBranchList)
		casbinRouter.POST("getBranchById", casbinApi.GetBranchById)
		casbinRouter.POST("updateBranch", casbinApi.UpdateBranch)
		casbinRouter.POST("deleteBranch", casbinApi.DeleteBranch)
	}
}
