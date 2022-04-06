package system

import (
	v1 "gandi.icu/demo/api/v1"
	"github.com/gin-gonic/gin"
)

type BranchRouter struct{}

func (s *BranchRouter) InitBranchRouter(Router *gin.RouterGroup) {
	branchRouter := Router.Group("branch")
	branchApi := v1.ApiGroupApp.SystemApiGroup.BranchApi
	{
		branchRouter.POST("createBranch", branchApi.CreateBranch)
		branchRouter.POST("getBranchList", branchApi.GetBranchList)
		branchRouter.POST("getBranchById", branchApi.GetBranchById)
		branchRouter.POST("updateBranch", branchApi.UpdateBranch)
		branchRouter.POST("deleteBranch", branchApi.DeleteBranch)
	}
}
